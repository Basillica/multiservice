package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/x/multiservice/parsejson/types"
	"github.com/x/multiservice/parsejson/types/appenv"
)

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}

func (t *Test) setup() {
}

func (t *Test) teardown() {
}

func (t *Test) SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.Use(
		t.MockConfigMiddleware(),
	)
	return router
}

func (t *Test) MockConfigMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("appenv", &appenv.AppConfig{
			AllowedOrigin: "http://localhost:3000",
		})
	}
}

func TestEncodeDecode(t *testing.T) {
	newTest := &Test{}
	newTest.setup()
	defer newTest.teardown()

	r := newTest.SetUpRouter()
	r.GET("/parsejson", ParseJSON)
	req, _ := http.NewRequest("GET", "/parsejson", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusOK, w.Code)

	var jsonResult Result
	json.NewDecoder(w.Body).Decode(&jsonResult)
	assert.Equal(t, jsonResult.Persons[0].Address, `munich`)

	// }
}

type Result struct {
	Persons []types.Person `json:"persons"`
}
