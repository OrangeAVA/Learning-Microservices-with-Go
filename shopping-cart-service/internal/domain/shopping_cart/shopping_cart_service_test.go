package shoppingcart_test

import (
	"errors"
	"testing"

	shoppingcart "github.com/go-microservices/shopping-cart-service/internal/domain/shopping_cart"
	mock_repository "github.com/go-microservices/shopping-cart-service/internal/domain/shopping_cart/mock"
	"github.com/go-microservices/shopping-cart-service/internal/shared/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testLogger := logger.NewLogger()

	mockRepo := mock_repository.NewMockShoppingCartRepository(ctrl)
	scService := shoppingcart.NewShoppingCartService(mockRepo, testLogger)

	randUUID := uuid.NewV4()
	testErrMsg := "error creating shopping cart"
	testErr := errors.New(testErrMsg)

	tt := []struct {
		name     string
		obj      shoppingcart.ShoppingCart
		err      error
		wantsErr bool
	}{
		{
			name: "create shopping cart - success result",
			obj: shoppingcart.ShoppingCart{
				UserID: &randUUID,
				ID:     nil,
			},
			err:      nil,
			wantsErr: false,
		},
		{
			name: "create shopping cart - bad request result",
			obj: shoppingcart.ShoppingCart{
				UserID: nil,
				ID:     nil,
			},
			err:      testErr,
			wantsErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().
				Create(gomock.Any()).
				Times(1).
				Return(tc.err)

			err := scService.Create(tc.obj)

			if tc.wantsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testLogger := logger.NewLogger()

	mockRepo := mock_repository.NewMockShoppingCartRepository(ctrl)
	scService := shoppingcart.NewShoppingCartService(mockRepo, testLogger)

	randUUID := uuid.NewV4()
	testErrMsg := "error getting shopping cart"
	testNotFoundErrMsg := "shopping cart not found"

	testErr := errors.New(testErrMsg)
	testNotFoundErr := errors.New(testNotFoundErrMsg)

	tt := []struct {
		name     string
		obj      *shoppingcart.ShoppingCart
		err      error
		wantsErr bool
	}{
		{
			name: "list shopping cart - success result",
			obj: &shoppingcart.ShoppingCart{
				UserID: &randUUID,
				ID:     &randUUID,
			},
			err:      nil,
			wantsErr: false,
		},
		{
			name:     "list shopping cart - error not found",
			obj:      nil,
			err:      testNotFoundErr,
			wantsErr: true,
		},
		{
			name:     "list shopping cart - db error",
			obj:      nil,
			err:      testErr,
			wantsErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().
				GetByUserID(gomock.Any()).
				Times(1).
				Return(tc.obj, tc.err)

			sp, err := scService.GetByUserID(&randUUID)

			if tc.wantsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, sp.ID, tc.obj.ID)
				assert.Equal(t, sp.ID, tc.obj.UserID)
			}
		})
	}
}
