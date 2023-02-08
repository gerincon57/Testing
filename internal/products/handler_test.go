package products

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func createServerProducts(st Repository) *gin.Engine {

	rp := NewStorageMock(st)
	sv := NewServiceLocal(rp)

	//server
	server := gin.Default()

	//router
	router := server.Group("/api/v1")
	{
		//h := NewControllerMovie(sv)
		group := router.Group("/movies")
		group.POST("", h.Create())

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

type responseMovie struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}
