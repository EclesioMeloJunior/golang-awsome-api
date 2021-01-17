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
	"net/http"
	"strings"
)

// Importation defines the abstraction of
// open food facts data into mongodb
type Importation interface {
	GetFilenames() ([]string, error)
	ToBeImported([]string) ([]models.Import, error)
	ImportFiles([]models.Import) ([]models.Product, error)
}

// ...
const (
	OpenFoodFactsList = "https://static.openfoodfacts.org/data/delta/index.txt"
	OpenFoodFactsData = "https://static.openfoodfacts.org/data/delta/%s"
	MaxRange          = 100
)

type importation struct {
	http            httpclient.HTTPClient
	importationRepo repository.Importation
	productsRepo    repository.Product
}

// NewImportation ...
func NewImportation(http httpclient.HTTPClient, i repository.Importation) Importation {
	return &importation{
		http:            http,
		importationRepo: i,
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

	return filenames, nil
}

func (i *importation) ToBeImported(filenames []string) ([]models.Import, error) {
	var err error
	var imports []models.Import

	if imports, err = i.importationRepo.GetAllImports(); err != nil {
		return nil, err
	}

	toBeImported := make([]models.Import, 0)

	for _, file := range filenames {
		i := getFileInImports(file, imports)

		// if the file is not included at
		// imports list then create a new Import for it
		if i == nil {
			toBeImported = append(toBeImported, models.Import{
				Filename: file,
				Imported: false,
				StopedAt: -1,
			})

			continue
		}

		if !i.Imported || i.StopedAt != -1 {
			toBeImported = append(toBeImported, *i)
		}
	}

	for _, i := range imports {
		if i.Imported {
			continue
		}

		toBeImported = append(toBeImported, i)
	}

	return toBeImported, nil
}

func (i *importation) ImportFiles(filenames []models.Import) ([]models.Product, error) {
	var err error
	var request *http.Request

	if request, err = createRequest(fmt.Sprintf(OpenFoodFactsData, filenames[0].Filename)); err != nil {
		return nil, err
	}

	var response *http.Response
	if response, err = i.http.Do(request); err != nil {
		return nil, err
	}

	var jsonGz *gzip.Reader
	defer response.Body.Close()
	if jsonGz, err = gzip.NewReader(response.Body); err != nil {
		return nil, err
	}

	var bodyBytes []byte
	if bodyBytes, err = ioutil.ReadAll(jsonGz); err != nil {
		return nil, err
	}

	productsToImport := make([]models.Product, 0)
	splitedContent := bytes.Split(bytes.Trim(bodyBytes, "\n"), []byte("\n"))

	for _, content := range splitedContent[:MaxRange] {
		p := models.Product{}

		if err = json.Unmarshal(content, &p); err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		productsToImport = append(productsToImport, p)
	}

	return productsToImport, nil
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
