package products

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerProducts() *gin.Engine {

	rp := NewStorageMock()
	sv := NewServiceLocal(rp)

	//server
	server := gin.Default()

	//router
	router := server.Group("/api/v1")
	{
		h := NewHandler(sv)
		group := router.Group("/products")
		group.GET("", h.GetProducts)

	}

	return server
}

func NewRequest(method, path, body string) (req *http.Request, res *httptest.ResponseRecorder) {
	// request
	req = httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	// response
	res = httptest.NewRecorder()

	return
}

// ______________________________________________________
// tests
var (
	ErrMock = errors.New("internal error")
)

func TestGetProducts(t *testing.T) {
	// arrange
	server := createServerProducts()

	// act
	t.Run("Happy Path", func(t *testing.T) {
		// arrange
		req, res := NewRequest(http.MethodGet, "/api/v1/products?seller_id='mock'", "")

		expectedProd := Product{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       30.25,
		}

		var expectedProdSlice = []Product{expectedProd}
		// act
		server.ServeHTTP(res, req)
		var r []Product
		err := json.Unmarshal(res.Body.Bytes(), &r)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, expectedProdSlice, r)
	})
}
