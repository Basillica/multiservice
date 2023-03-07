package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/x/multiservice/algorithm/types"
	"github.com/x/multiservice/algorithm/types/appenv"
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

func (t *Test) testcase(r *gin.Engine, payload types.ArrayRequest) (*httptest.ResponseRecorder, Result) {
	response, _ := json.Marshal(payload)
	algReq, _ := http.NewRequest("POST", "/algorithm", bytes.NewBuffer(response))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, algReq)
	jsonResult := Result{}
	json.NewDecoder(w.Body).Decode(&jsonResult)
	return w, jsonResult
}

func (t *Test) testcases() []types.ArrayRequest {
	return []types.ArrayRequest{
		{
			// test case when the given integer is found in array
			Array: []int{5, 6, 5, 4, 289, 345, 8988, -343, 9090, 5543, 45433, 5667, 22322, 5654, 6554, 65544, 543455, 23221},
			Int:   9090,
		},
		{
			// example test case when no match if found
			Array: []int{},
			Int:   300,
		},
		{
			// test scenario when there are dulplicate values in array
			Array: []int{500, 500, 500, 500, 500, 500, 6, 5, 4, 289, 345, 8988, -343, 9090, 5543, 45433, 5667, 22322, 5654, 6554, 65544, 543455, 23221},
			Int:   500,
		},
	}
}

func (t *Test) testResults(index int) int {
	results := [3]int{12, -1, 6}
	return results[index]
}

func TestAlgorithm(t *testing.T) {
	newTest := &Test{}
	newTest.setup()
	defer newTest.teardown()

	r := newTest.SetUpRouter()
	r.POST("/algorithm", Algorithm)

	for index, testPayload := range newTest.testcases() {
		w, jsonResult := newTest.testcase(r, testPayload)
		assert.Equal(t, jsonResult.Result.Integer, newTest.testResults(index))
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

type Result struct {
	Result T `json:"result"`
}

type T struct {
	Integer int `json:"integer"`
}
