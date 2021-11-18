package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
)

func (s *IntegrationTestSuite) TestGetCart() {
	router := gin.New()
	s.httpHandler.Init(router.Group("/cart"))

	r := s.Require()

	req, _ := http.NewRequest("GET", "/cart/v1/", nil)
	req.AddCookie(&http.Cookie{Name: "USER_ID", Value: testGetCartUserID})
	req.Header.Set("Content-type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	expectedJson, _ := json.Marshal(map[string]map[string]string{"cart": {testProductID: strconv.FormatInt(testBigCount, 10)}})
	expectedBody := bytes.NewBuffer(expectedJson).String()

	r.Equal(expectedBody, rec.Body.String())
	r.Equal(http.StatusOK, rec.Result().StatusCode)
}

type cartInput struct {
	ProductID string `json:"product_id" binding:"required"`
	Count     int64  `json:"count"`
}

func (s *IntegrationTestSuite) TestAddToCart() {
	router := gin.New()
	s.httpHandler.Init(router.Group("/cart"))

	r := s.Require()

	body, _ := json.Marshal(cartInput{
		ProductID: testProductID,
		Count:     testSmallCount,
	})

	req, _ := http.NewRequest("POST", "/cart/v1/add", bytes.NewBuffer(body))
	req.AddCookie(&http.Cookie{Name: "USER_ID", Value: testAddToCartUserID})
	req.Header.Set("Content-type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	resultMap, err := s.db.HGetAll(context.Background(), testAddToCartUserID).Result()
	r.NoError(err)
	expectedMap := map[string]string{testProductID: strconv.FormatInt(testSmallCount, 10)}

	r.Truef(reflect.DeepEqual(resultMap, expectedMap), "result maps are not equal", resultMap, expectedMap)
	r.Equal("", rec.Body.String())
	r.Equal(http.StatusOK, rec.Result().StatusCode)
}

func (s *IntegrationTestSuite) TestRemoveFromCart() {
	router := gin.New()
	s.httpHandler.Init(router.Group("/cart"))

	r := s.Require()

	body, _ := json.Marshal(cartInput{
		ProductID: testProductID,
		Count:     testSmallCount,
	})

	req, _ := http.NewRequest("POST", "/cart/v1/remove", bytes.NewBuffer(body))
	req.AddCookie(&http.Cookie{Name: "USER_ID", Value: testRemoveFromCartUserID})
	req.Header.Set("Content-type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	resultMap, err := s.db.HGetAll(context.Background(), testRemoveFromCartUserID).Result()
	r.NoError(err)
	expectedMap := map[string]string{testProductID: strconv.FormatInt(testBigCount-testSmallCount, 10)}

	r.Truef(reflect.DeepEqual(resultMap, expectedMap), "result maps are not equal", resultMap, expectedMap)
	r.Equal("", rec.Body.String())
	r.Equal(http.StatusOK, rec.Result().StatusCode)
}

func (s *IntegrationTestSuite) TestCleanCart() {
	router := gin.New()
	s.httpHandler.Init(router.Group("/cart"))

	r := s.Require()

	req, _ := http.NewRequest("POST", "/cart/v1/remove/all", nil)
	req.AddCookie(&http.Cookie{Name: "USER_ID", Value: testCleanCartUserID})

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	resultMap, err := s.db.HGetAll(context.Background(), testCleanCartUserID).Result()
	r.NoError(err)
	expectedMap := map[string]string{}

	r.Truef(reflect.DeepEqual(resultMap, expectedMap), "result maps are not equal", resultMap, expectedMap)
	r.Equal("", rec.Body.String())
	r.Equal(http.StatusOK, rec.Result().StatusCode)
}
