package services

import (
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
}

// ...
const (
	OpenFoodFactsList = "https://static.openfoodfacts.org/data/delta/index.txt"
	OpenFoodFactsData = "https://static.openfoodfacts.org/data/delta/%s"
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

	filenames := strings.Split(string(bodyBytes), "\n")

	return filenames, nil
}

func (i *importation) ToBeImported(filenames []string) ([]models.Import, error) {
	var err error
	var imports []models.Import

	if imports, err = i.importationRepo.GetAllImports(); err != nil {
		return nil, err
	}

	toBeImported := make([]models.Import, 0)

	for _, i := range imports {
		if i.Imported {
			continue
		}

		toBeImported = append(toBeImported, i)
	}

	return toBeImported, nil
}

func createRequest(addr string) (*http.Request, error) {
	return http.NewRequest(
		http.MethodGet,
		addr,
		nil,
	)
}
