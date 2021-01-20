package handlers_test

import (
	"bytes"
	"encoding/json"
	"go-challenge/internals/handlers"
	"go-challenge/internals/models"
	"go-challenge/mocks/internals/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func createRecorderAndCtx(method, endpoint string, data io.Reader) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()

	req := httptest.NewRequest(method, endpoint, data)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)

	return rr, c
}

var _ = Describe("Given an user wants retrive a list of products", func() {
	var (
		productService *services.Product
		h              *handlers.Products
	)

	BeforeEach(func() {

		productService = &services.Product{}
		h = handlers.NewProductsHandler(productService)
	})

	When("user does not send the pagination params", func() {
		It("should use the default pagination params", func() {
			rr, c := createRecorderAndCtx(http.MethodGet, "/products", nil)

			productService.On("GetProducts", nil, 0, 10).
				Return([]models.Product{}, nil)

			err := h.GetProductsList(c)

			expected := []byte("{\"data\":[],\"success\":true}\n")

			Expect(err).To(BeNil())
			Expect(rr.Body.Bytes()).To(Equal(expected))
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

	When("user send the pagination params", func() {
		It("should use the client params", func() {
			rr, c := createRecorderAndCtx(http.MethodGet, "/products?page=2&size=70", nil)

			productService.On("GetProducts", nil, 2, 70).
				Return([]models.Product{}, nil)

			err := h.GetProductsList(c)

			expected := []byte("{\"data\":[],\"success\":true}\n")

			Expect(err).To(BeNil())
			Expect(rr.Body.Bytes()).To(Equal(expected))
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

	When("application found one or more products", func() {
		It("should return a list to the client", func() {
			rr, c := createRecorderAndCtx(http.MethodGet, "/products?page=2&size=70", nil)

			productService.On("GetProducts", nil, 2, 70).
				Return([]models.Product{
					{Code: "01"},
					{Code: "02"},
					{Code: "03"},
				}, nil)

			err := h.GetProductsList(c)

			resp := new(handlers.Response)
			json.Unmarshal(bytes.TrimSpace(rr.Body.Bytes()), resp)

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
			Expect(resp.Body).To(HaveLen(3))
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})
})

var _ = Describe("Given an user wants retrive a product by ID", func() {
	var (
		productService *services.Product
		h              *handlers.Products
	)

	BeforeEach(func() {

		productService = &services.Product{}
		h = handlers.NewProductsHandler(productService)
	})

	When("the document does not exists on database", func() {
		It("Should return No Content", func() {
			rr, c := createRecorderAndCtx(http.MethodGet, "/products/productid", nil)
			c.SetParamNames("pcode")
			c.SetParamValues("productid")

			productService.On("GetProductByID", "productid").
				Return(nil, mongo.ErrNoDocuments)

			err := h.GetProductByID(c)

			Expect(err).To(BeNil())
			Expect(rr.Code).To(Equal(http.StatusNoContent))
		})
	})

	When("the document exists on database", func() {
		It("Should return the product", func() {
			rr, c := createRecorderAndCtx(http.MethodGet, "/products/productid", nil)
			c.SetParamNames("pcode")
			c.SetParamValues("productid")

			product := new(models.Product)
			product.ID = primitive.NilObjectID
			product.Code = "somecode"
			product.Status = models.Published

			productService.On("GetProductByID", "productid").
				Return(product, nil)

			err := h.GetProductByID(c)

			resp := new(handlers.Response)
			json.Unmarshal(bytes.TrimSpace(rr.Body.Bytes()), resp)

			expectedBody := map[string]interface{}{
				"_id":    primitive.NilObjectID.Hex(),
				"code":   "somecode",
				"status": "published",
			}

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
			Expect(resp.Body).To(Equal(expectedBody))
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

})

var _ = Describe("Given an user wants to update a product", func() {
	var (
		productService *services.Product
		h              *handlers.Products
	)

	BeforeEach(func() {

		productService = &services.Product{}
		h = handlers.NewProductsHandler(productService)
	})

	When("user sends an invalid body", func() {
		It("should return Bad Request", func() {
			tests := []struct {
				body    string
				message string
			}{
				{
					body:    `{"_id": "110000000000000000000000"}`,
					message: "Invalid body data: _id",
				},
				{
					body:    `{"code": "invalid prop"}`,
					message: "Invalid body data: code",
				},
				{
					body:    `{"imported_t": 99999}`,
					message: "Invalid body data: imported_t",
				},
				{
					body:    `{"status": "invalid prop"}`,
					message: "Invalid body data: status",
				},
			}

			for _, t := range tests {
				rr, c := createRecorderAndCtx(http.MethodPut, "/products/someproductid", strings.NewReader(t.body))

				err := h.UpdateProductByID(c)

				resp := new(handlers.Response)
				json.Unmarshal(bytes.TrimSpace(rr.Body.Bytes()), resp)

				Expect(err).To(BeNil())
				Expect(resp.Message).To(Equal(t.message))
				Expect(resp.Body).To(BeNil())
				Expect(resp.Success).To(BeFalse())
				Expect(rr.Code).To(Equal(http.StatusBadRequest))
			}
		})
	})

	When("user send valid information but document does not exists", func() {
		It("should reutrn No Content", func() {
			body := `{
				"url": "url nova do produto",
				"product_name": "Borizanos kouloirakia",
				"quantity": "10",
				"brands": "new brands",
				"categories": "",
				"labels": "",
				"cities": "",
				"purchase_places": "",
				"stores": "",
				"ingredients_text": "",
				"traces": "",
				"serving_size": "",
				"serving_quantity": 0,
				"nutriscore_score": 0,
				"nutriscore_grade": "",
				"main_category": "",
				"image_url": ""
			}`

			rr, c := createRecorderAndCtx(http.MethodPut, "/products/someproductid", strings.NewReader(body))
			c.SetParamNames("pcode")
			c.SetParamValues("someproductid")

			productService.On("UpdateProductByID", "someproductid", mock.AnythingOfType("*models.Product")).
				Return(nil, mongo.ErrNoDocuments)

			err := h.UpdateProductByID(c)

			Expect(err).To(BeNil())
			Expect(rr.Code).To(Equal(http.StatusNoContent))
		})
	})

	When("user send a valid information and document was updated", func() {
		It("should return OK", func() {
			body := `{
				"url": "url nova do produto",
				"product_name": "Borizanos kouloirakia",
				"quantity": "10",
				"brands": "new brands",
				"categories": "",
				"labels": "",
				"cities": "",
				"purchase_places": "",
				"stores": "",
				"ingredients_text": "",
				"traces": "",
				"serving_size": "",
				"serving_quantity": 0,
				"nutriscore_score": 0,
				"nutriscore_grade": "",
				"main_category": "",
				"image_url": ""
			}`

			rr, c := createRecorderAndCtx(http.MethodPut, "/products/someproductid", strings.NewReader(body))
			c.SetParamNames("pcode")
			c.SetParamValues("someproductid")

			productService.On("UpdateProductByID", "someproductid", mock.AnythingOfType("*models.Product")).
				Return(&models.Product{}, nil)

			err := h.UpdateProductByID(c)

			resp := new(handlers.Response)
			json.Unmarshal(bytes.TrimSpace(rr.Body.Bytes()), resp)

			Expect(err).To(BeNil())
			Expect(resp.Success).To(BeTrue())
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})
})

func TestProducts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Products handler suite")
}
