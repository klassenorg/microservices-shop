package v1

import (
	"bytes"
	"cartsvc/internal/service"
	mock_service "cartsvc/internal/service/mocks"
	log "cartsvc/pkg/logger"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_getCart(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCart, userID string, cart map[string]string)

	tests := []struct {
		name            string
		hasUserIDCookie bool
		userID          string
		cart            map[string]string
		mockBehavior    mockBehavior
		statusCode      int
		responseBody    string
	}{
		{
			name:            "all good",
			hasUserIDCookie: true,
			userID:          "123qwe-321",
			cart:            map[string]string{"1": "10"},
			mockBehavior: func(s *mock_service.MockCart, userID string, cart map[string]string) {
				s.EXPECT().GetCart(context.Background(), userID).Return(cart, nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "{\"cart\":{\"1\":\"10\"}}",
		},
		{
			name:            "empty cart",
			hasUserIDCookie: true,
			userID:          "123qwe-321",
			cart:            map[string]string{},
			mockBehavior: func(s *mock_service.MockCart, userID string, cart map[string]string) {
				s.EXPECT().GetCart(context.Background(), userID).Return(cart, nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "{\"cart\":{}}",
		},
		{
			name:            "USER_ID cookie is empty",
			hasUserIDCookie: true,
			userID:          "",
			cart:            nil,
			mockBehavior:    func(s *mock_service.MockCart, userID string, cart map[string]string) {},
			statusCode:      http.StatusBadRequest,
			responseBody:    "{\"error\":\"USER_ID cookie is empty\"}",
		},
		{
			name:            "haven't USER_ID cookie",
			hasUserIDCookie: false,
			userID:          "",
			cart:            nil,
			mockBehavior:    func(s *mock_service.MockCart, userID string, cart map[string]string) {},
			statusCode:      http.StatusBadRequest,
			responseBody:    "{\"error\":\"USER_ID cookie not present\"}",
		},
		{
			name:            "service error",
			hasUserIDCookie: true,
			userID:          "123qwe-321",
			cart:            map[string]string{},
			mockBehavior: func(s *mock_service.MockCart, userID string, cart map[string]string) {
				s.EXPECT().GetCart(context.Background(), userID).Return(cart, errors.New("some error"))
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: "{\"error\":\"internal server error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_service.NewMockCart(c)

			tt.mockBehavior(mockCart, tt.userID, tt.cart)

			//non-logging logger
			l := zap.New(nil).Sugar()
			logger := &log.Logger{SugaredLogger: l}

			services := &service.Services{Cart: mockCart}
			handler := Handler{
				services: services,
				logger:   logger,
			}

			gin.SetMode(gin.ReleaseMode)
			router := gin.New()
			router.Use(handler.hasUserIDCookie())
			router.GET("/cart/v1", handler.getCart)

			r := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/cart/v1", nil)
			if tt.hasUserIDCookie {
				req.AddCookie(&http.Cookie{Name: "USER_ID", Value: tt.userID})
			}
			router.ServeHTTP(r, req)

			assert.Equal(t, tt.statusCode, r.Code)
			assert.Equal(t, tt.responseBody, r.Body.String())
		})

	}
}

func TestHandler_addToCart(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCart, userID string, productID string, count int64)

	tests := []struct {
		name                 string
		hasUserIDCookie      bool
		userID               string
		productID            string
		productCount         int64
		expectedProductCount int64
		mockBehavior         mockBehavior
		statusCode           int
		responseBody         string
	}{
		{
			name:                 "all good",
			hasUserIDCookie:      true,
			userID:               "123-test",
			productID:            "123456",
			productCount:         10,
			expectedProductCount: 10,
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().AddToCart(context.Background(), userID, productID, count).Return(nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "",
		},
		{
			name:                 "empty or zero count",
			hasUserIDCookie:      true,
			userID:               "123-test",
			productID:            "10000",
			expectedProductCount: 1,
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().AddToCart(context.Background(), userID, productID, count).Return(nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "",
		},
		{
			name:                 "negative count",
			hasUserIDCookie:      true,
			userID:               "123-test",
			productID:            "10000",
			productCount:         -10,
			expectedProductCount: 1,
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().AddToCart(context.Background(), userID, productID, count).Return(nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "",
		},
		{
			name:                 "empty USER_ID cookie",
			hasUserIDCookie:      true,
			userID:               "",
			productID:            "123456",
			productCount:         10,
			expectedProductCount: 10,
			mockBehavior:         func(s *mock_service.MockCart, userID string, productID string, count int64) {},
			statusCode:           http.StatusBadRequest,
			responseBody:         "{\"error\":\"USER_ID cookie is empty\"}",
		},
		{
			name:                 "USER_ID cookie not present",
			hasUserIDCookie:      false,
			userID:               "",
			productID:            "123456",
			productCount:         10,
			expectedProductCount: 10,
			mockBehavior:         func(s *mock_service.MockCart, userID string, productID string, count int64) {},
			statusCode:           http.StatusBadRequest,
			responseBody:         "{\"error\":\"USER_ID cookie not present\"}",
		},
		{
			name:                 "service error",
			hasUserIDCookie:      true,
			userID:               "123-test",
			productID:            "10000",
			productCount:         1,
			expectedProductCount: 1,
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().AddToCart(context.Background(), userID, productID, count).Return(errors.New("test error"))
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: "{\"error\":\"internal server error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_service.NewMockCart(c)

			tt.mockBehavior(mockCart, tt.userID, tt.productID, tt.expectedProductCount)
			//non-logging logger
			l := zap.New(nil).Sugar()
			logger := &log.Logger{SugaredLogger: l}

			services := &service.Services{Cart: mockCart}
			handler := Handler{
				services: services,
				logger:   logger,
			}

			gin.SetMode(gin.ReleaseMode)
			router := gin.New()
			router.Use(handler.hasUserIDCookie())
			router.POST("/cart/v1/add", handler.addToCart)

			r := httptest.NewRecorder()
			body, _ := json.Marshal(map[string]interface{}{"product_id": tt.productID, "count": tt.productCount})
			req := httptest.NewRequest("POST", "/cart/v1/add", bytes.NewBuffer(body))
			if tt.hasUserIDCookie {
				req.AddCookie(&http.Cookie{Name: "USER_ID", Value: tt.userID})
			}
			router.ServeHTTP(r, req)

			assert.Equal(t, tt.statusCode, r.Code)
			assert.Equal(t, tt.responseBody, r.Body.String())
		})

	}
}

func TestHandler_removeFromCart(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCart, userID string, productID string, count int64)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		hasUserIDCookie      bool
		userID               string
		productID            string
		productCount         int64
		expectedProductCount int64
		statusCode           int
		responseBody         string
	}{
		{
			name: "all good",
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().RemoveFromCart(context.Background(), userID, productID, count).Return(nil)
			},
			hasUserIDCookie:      true,
			userID:               "1234-test",
			productID:            "12345",
			productCount:         1,
			expectedProductCount: 1,
			statusCode:           http.StatusOK,
			responseBody:         "",
		},
		{
			name: "negative count",
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().RemoveFromCart(context.Background(), userID, productID, count).Return(nil)
			},
			hasUserIDCookie:      true,
			userID:               "1234-test",
			productID:            "12345",
			productCount:         -1,
			expectedProductCount: 1,
			statusCode:           http.StatusOK,
			responseBody:         "",
		},
		{
			name: "empty or zero count",
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().RemoveFromCart(context.Background(), userID, productID, count).Return(nil)
			},
			hasUserIDCookie:      true,
			userID:               "1234-test",
			productID:            "12345",
			expectedProductCount: 1,
			statusCode:           http.StatusOK,
			responseBody:         "",
		},
		{
			name:                 "empty USER_ID cookie",
			mockBehavior:         func(s *mock_service.MockCart, userID string, productID string, count int64) {},
			hasUserIDCookie:      true,
			userID:               "",
			productID:            "12345",
			productCount:         1,
			expectedProductCount: 1,
			statusCode:           http.StatusBadRequest,
			responseBody:         "{\"error\":\"USER_ID cookie is empty\"}",
		},
		{
			name:                 "USER_ID cookie not present",
			mockBehavior:         func(s *mock_service.MockCart, userID string, productID string, count int64) {},
			hasUserIDCookie:      false,
			userID:               "",
			productID:            "12345",
			productCount:         1,
			expectedProductCount: 1,
			statusCode:           http.StatusBadRequest,
			responseBody:         "{\"error\":\"USER_ID cookie not present\"}",
		},
		{
			name: "service error",
			mockBehavior: func(s *mock_service.MockCart, userID string, productID string, count int64) {
				s.EXPECT().RemoveFromCart(context.Background(), userID, productID, count).Return(errors.New("test error"))
			},
			hasUserIDCookie:      true,
			userID:               "1234-test",
			productID:            "12345",
			productCount:         1,
			expectedProductCount: 1,
			statusCode:           http.StatusInternalServerError,
			responseBody:         "{\"error\":\"internal server error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_service.NewMockCart(c)

			tt.mockBehavior(mockCart, tt.userID, tt.productID, tt.expectedProductCount)
			//non-logging logger
			l := zap.New(nil).Sugar()
			logger := &log.Logger{SugaredLogger: l}

			services := &service.Services{Cart: mockCart}
			handler := Handler{
				services: services,
				logger:   logger,
			}

			gin.SetMode(gin.ReleaseMode)
			router := gin.New()
			router.Use(handler.hasUserIDCookie())
			router.POST("/cart/v1/remove", handler.removeFromCart)

			r := httptest.NewRecorder()
			body, _ := json.Marshal(map[string]interface{}{"product_id": tt.productID, "count": tt.productCount})
			req := httptest.NewRequest("POST", "/cart/v1/remove", bytes.NewBuffer(body))
			if tt.hasUserIDCookie {
				req.AddCookie(&http.Cookie{Name: "USER_ID", Value: tt.userID})
			}
			router.ServeHTTP(r, req)

			assert.Equal(t, tt.statusCode, r.Code)
			assert.Equal(t, tt.responseBody, r.Body.String())
		})
	}
}

func TestHandler_cleanCart(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCart, userID string)

	tests := []struct {
		name            string
		hasUserIDCookie bool
		userID          string
		mockBehavior    mockBehavior
		statusCode      int
		responseBody    string
	}{
		{
			name:            "all good",
			hasUserIDCookie: true,
			userID:          "1234-test",
			mockBehavior: func(s *mock_service.MockCart, userID string) {
				s.EXPECT().RemoveAllFromCart(context.Background(), userID).Return(nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "",
		},
		{
			name:            "empty cart",
			hasUserIDCookie: true,
			userID:          "1234",
			mockBehavior: func(s *mock_service.MockCart, userID string) {
				s.EXPECT().RemoveAllFromCart(context.Background(), userID).Return(nil)
			},
			statusCode:   http.StatusOK,
			responseBody: "",
		},
		{
			name:            "empty USER_ID cookie",
			hasUserIDCookie: true,
			userID:          "",
			mockBehavior:    func(s *mock_service.MockCart, userID string) {},
			statusCode:      http.StatusBadRequest,
			responseBody:    "{\"error\":\"USER_ID cookie is empty\"}",
		},
		{
			name:            "USER_ID cookie not present",
			hasUserIDCookie: false,
			userID:          "",
			mockBehavior:    func(s *mock_service.MockCart, userID string) {},
			statusCode:      http.StatusBadRequest,
			responseBody:    "{\"error\":\"USER_ID cookie not present\"}",
		},
		{
			name:            "service error",
			hasUserIDCookie: true,
			userID:          "1234-test",
			mockBehavior: func(s *mock_service.MockCart, userID string) {
				s.EXPECT().RemoveAllFromCart(context.Background(), userID).Return(errors.New("test error"))
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: "{\"error\":\"internal server error\"}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_service.NewMockCart(c)

			tt.mockBehavior(mockCart, tt.userID)

			l := zap.New(nil).Sugar()
			logger := &log.Logger{SugaredLogger: l}

			services := &service.Services{Cart: mockCart}
			handler := Handler{
				services: services,
				logger:   logger,
			}

			gin.SetMode(gin.ReleaseMode)
			router := gin.New()
			router.Use(handler.hasUserIDCookie())
			router.POST("/cart/v1/remove/all", handler.cleanCart)

			r := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/cart/v1/remove/all", nil)
			if tt.hasUserIDCookie {
				req.AddCookie(&http.Cookie{Name: "USER_ID", Value: tt.userID})
			}
			router.ServeHTTP(r, req)

			assert.Equal(t, tt.statusCode, r.Code)
			assert.Equal(t, tt.responseBody, r.Body.String())
		})
	}
}
