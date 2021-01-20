package services

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"go-challenge/internals/models"
	"go-challenge/internals/repository"
	"go-challenge/pkg/httpclient"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

// Importation defines the abstraction of
// open food facts data into mongodb
type Importation interface {
	GetFilenames() ([]string, error)
	ToBeImported([]string) ([]models.Import, error)
	ImportFiles([]models.Import) error
}

// ...
const (
	OpenFoodFactsList = "https://static.openfoodfacts.org/data/delta/index.txt"
	OpenFoodFactsData = "https://static.openfoodfacts.org/data/delta/%s"
	MaxRange          = 100
)

type importation struct {
	http            httpclient.HTTPClient
	importationRepo repository.Import
	productsRepo    repository.Product
	hcService       Healthcheck
}

// NewImportation ...
func NewImportation(http httpclient.HTTPClient, h Healthcheck, i repository.Import) Importation {
	return &importation{
		http:            http,
		importationRepo: i,
		hcService:       h,
	}
}

func (i *importation) GetFilenames() ([]string, error) {
	var err error
	var request *http.Request
	var response *http.Response

	if request, err = createRequest(OpenFoodFactsList); err != nil {
		return nil, err
	}

	if response, err = i.http.Do(request); err != nil {
		return nil, err
	}

	var bodyBytes []byte
	defer response.Body.Close()
	if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	filenames := strings.Split(strings.Trim(string(bodyBytes), "\n"), "\n")
	sort.Strings(filenames)

	return filenames, nil
}

func (i *importation) ToBeImported(filenames []string) ([]models.Import, error) {
	var err error
	var imports []models.Import

	if imports, err = i.importationRepo.GetAllImports(); err != nil {
		return nil, err
	}

	toBeImported := make([]models.Import, 0)

	for idx := len(filenames) - 1; idx >= 0; idx-- {
		i := getFileInImports(filenames[idx], imports)

		// if the file is not included at
		// imports list then create a new Import for it
		if i == nil {
			toBeImported = append(toBeImported, models.Import{
				Filename: filenames[idx],
			})

			break
		}
	}

	return toBeImported, nil
}

func (i *importation) ImportFiles(filenames []models.Import) error {
	if len(filenames) < 1 {
		return nil
	}

	errChan := make(chan error, len(filenames))
	defer close(errChan)

	wg := sync.WaitGroup{}

	for _, filename := range filenames {
		wg.Add(1)
		go i.retrieveAndImport(filename, &wg, errChan)
	}

	wg.Wait()
	if len(errChan) > 0 {
		return <-errChan
	}

	return nil
}

func (i *importation) retrieveAndImport(imp models.Import, wg *sync.WaitGroup, errChan chan<- error) {
	defer wg.Done()

	var err error
	var request *http.Request

	url := fmt.Sprintf(OpenFoodFactsData, imp.Filename)
	if request, err = createRequest(url); err != nil {
		errChan <- createError(imp, err)
		return
	}

	log.Printf("Request to %s started: %s\n", imp.Filename, url)
	var response *http.Response
	if response, err = i.http.Do(request); err != nil {
		errChan <- createError(imp, err)
		return
	}

	var jsonGz *gzip.Reader
	defer response.Body.Close()
	if jsonGz, err = gzip.NewReader(response.Body); err != nil {
		errChan <- createError(imp, err)
		return
	}

	var bodyBytes []byte
	if bodyBytes, err = ioutil.ReadAll(jsonGz); err != nil {
		errChan <- createError(imp, err)
		return
	}

	dataRows := bytes.Split(bytes.Trim(bodyBytes, "\n"), []byte("\n"))
	log.Printf("Amount of data to %s: %v\n", imp.Filename, len(dataRows))

	if len(dataRows) > MaxRange {
		dataRows = dataRows[:MaxRange]
	}

	log.Printf("Amount of data to %s after cut: %v\n", imp.Filename, len(dataRows))

	productsToInsert := make([]interface{}, 0)

	for _, content := range dataRows {
		p := models.Product{}

		content = bytes.Replace(content, []byte("_id"), []byte("_exported_id"), -1)

		if err = json.Unmarshal(content, &p); err != nil {
			errChan <- createError(imp, err)
			return
		}

		p.Status = models.Published
		p.ImportedT = time.Now().Unix()
		productsToInsert = append(productsToInsert, p)
	}

	imp.ImportedT = time.Now().Unix()
	imp.Quantity = len(productsToInsert)

	log.Printf("Import of data to %s started: %v products\n", imp.Filename, len(productsToInsert))

	if err = i.importationRepo.ExecuteImport(&imp, productsToInsert); err != nil {
		errChan <- createError(imp, err)
		return
	}
}

func createError(imp models.Import, err error) error {
	return fmt.Errorf("error while importing %s: %v", imp.Filename, err)
}

func createRequest(addr string) (*http.Request, error) {
	return http.NewRequest(
		http.MethodGet,
		addr,
		nil,
	)
}

func getFileInImports(filename string, imports []models.Import) *models.Import {
	if len(imports) < 1 {
		return nil
	}

	for _, i := range imports {
		if i.Filename == filename {
			return &i
		}
	}

	return nil
}
