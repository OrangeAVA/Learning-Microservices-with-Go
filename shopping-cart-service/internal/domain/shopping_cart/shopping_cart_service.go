package shoppingcart

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type shoppingCartService struct {
	logger           *logrus.Logger
	shoppingCartRepo ShoppingCartRepository
}

func NewShoppingCartService(scRepo ShoppingCartRepository, logger *logrus.Logger) shoppingCartService {
	return shoppingCartService{
		logger:           logger,
		shoppingCartRepo: scRepo,
	}
}

func (sc shoppingCartService) Create(spc ShoppingCart) error {
	sc.logger.Info("On Create Shopping")

	return sc.shoppingCartRepo.Create(spc)
}

func (sc shoppingCartService) GetByUserID(userId *uuid.UUID) (*ShoppingCart, error) {
	sc.logger.Info("On Get Shopping Cart for user")

	shoppingCart, err := sc.shoppingCartRepo.GetByUserID(userId)
	if err != nil {
		sc.logger.Errorf("On Get Shopping Cart for user - error quering db: %d", err)
		return nil, err
	}

	return shoppingCart, nil
}
