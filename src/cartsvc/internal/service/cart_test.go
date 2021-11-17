package service

import (
	mock_repository "cartsvc/internal/repository/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testError = errors.New("test error")
var ctx = context.Background()

func TestCartService_GetCart(t *testing.T) {
	type mockBehavior func(cart *mock_repository.MockCart, userID string)

	tests := []struct {
		name           string
		mockBehavior   mockBehavior
		userID         string
		err            error
		expectedResult map[string]string
	}{
		{
			name: "no error",
			mockBehavior: func(cart *mock_repository.MockCart, userID string) {
				cart.EXPECT().GetByID(ctx, userID).Return(map[string]string{"1": "1"}, nil)
			},
			userID:         "1234-test",
			err:            nil,
			expectedResult: map[string]string{"1": "1"},
		},
		{
			name: "error",
			mockBehavior: func(cart *mock_repository.MockCart, userID string) {
				cart.EXPECT().GetByID(ctx, userID).Return(map[string]string{}, testError)
			},
			userID:         "1234-test",
			err:            testError,
			expectedResult: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_repository.NewMockCart(c)
			cartService := NewCartService(mockCart)

			tt.mockBehavior(mockCart, tt.userID)

			res, err := cartService.GetCart(ctx, tt.userID)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.expectedResult, res)
		})
	}

}

func TestCartService_AddToCart(t *testing.T) {
	type mockBehavior func(cart *mock_repository.MockCart, userID string, productID string, count int64)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		userID       string
		productID    string
		count        int64
		err          error
	}{
		{
			name: "all good",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().AddByID(ctx, userID, productID, count).Return(nil)
			},
			userID:    "1234-test",
			productID: "10000",
			count:     1,
			err:       nil,
		},
		{
			name: "zero count",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().AddByID(ctx, userID, productID, int64(1)).Return(nil)
			},
			userID:    "1234-test",
			productID: "10000",
			err:       nil,
		},
		{
			name: "negative count",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().AddByID(ctx, userID, productID, int64(1)).Return(nil)
			},
			userID:    "1234-test",
			count:     -1,
			productID: "10000",
			err:       nil,
		},
		{
			name: "error",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().AddByID(ctx, userID, productID, count).Return(testError)
			},
			userID:    "1234-test",
			productID: "10000",
			count:     1,
			err:       testError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_repository.NewMockCart(c)
			cartService := NewCartService(mockCart)

			tt.mockBehavior(mockCart, tt.userID, tt.productID, tt.count)

			assert.Equal(t, tt.err, cartService.AddToCart(ctx, tt.userID, tt.productID, tt.count))
		})
	}
}

func TestCartService_RemoveFromCart(t *testing.T) {
	type mockBehavior func(cart *mock_repository.MockCart, userID string, productID string, count int64)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		userID       string
		productID    string
		count        int64
		err          error
	}{
		{
			name: "all good",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().RemoveByID(ctx, userID, productID, count).Return(nil)
			},
			userID:    "1234-test",
			productID: "10000",
			count:     1,
			err:       nil,
		},
		{
			name: "zero count",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().RemoveByID(ctx, userID, productID, int64(1)).Return(nil)
			},
			userID:    "1234-test",
			productID: "10000",
			err:       nil,
		},
		{
			name: "negative count",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().RemoveByID(ctx, userID, productID, int64(1)).Return(nil)
			},
			userID:    "1234-test",
			count:     -1,
			productID: "10000",
			err:       nil,
		},
		{
			name: "error",
			mockBehavior: func(cart *mock_repository.MockCart, userID string, productID string, count int64) {
				cart.EXPECT().RemoveByID(ctx, userID, productID, count).Return(testError)
			},
			userID:    "1234-test",
			productID: "10000",
			count:     1,
			err:       testError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_repository.NewMockCart(c)
			cartService := NewCartService(mockCart)

			tt.mockBehavior(mockCart, tt.userID, tt.productID, tt.count)

			assert.Equal(t, tt.err, cartService.RemoveFromCart(ctx, tt.userID, tt.productID, tt.count))
		})
	}
}

func TestCartService_RemoveAllFromCart(t *testing.T) {
	type mockBehavior func(cart *mock_repository.MockCart, userID string)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		userID       string
		err          error
	}{
		{
			name: "all good",
			mockBehavior: func(cart *mock_repository.MockCart, userID string) {
				cart.EXPECT().RemoveAllByID(ctx, userID).Return(nil)
			},
			userID: "1234-test",
			err:    nil,
		},
		{
			name: "error",
			mockBehavior: func(cart *mock_repository.MockCart, userID string) {
				cart.EXPECT().RemoveAllByID(ctx, userID).Return(testError)
			},
			userID: "1234-test",
			err:    testError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCart := mock_repository.NewMockCart(c)
			cartService := NewCartService(mockCart)

			tt.mockBehavior(mockCart, tt.userID)

			assert.Equal(t, tt.err, cartService.RemoveAllFromCart(ctx, tt.userID))
		})
	}
}
