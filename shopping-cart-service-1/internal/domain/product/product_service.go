package product

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type productService struct {
	logger      *logrus.Logger
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository, logger *logrus.Logger) productService {
	return productService{
		logger:      logger,
		productRepo: productRepo,
	}
}

func (ps productService) Create(product Product) error {
	ps.logger.Info("On Create product service")

	if err := ps.productRepo.Create(product); err != nil {
		return err
	}

	return nil
}

func (ps productService) GetAll(shoppingCartId uuid.UUID) ([]Product, error) {
	ps.logger.Info("On GetAll product service")

	return ps.productRepo.Get(shoppingCartId)
}

func (ps productService) Delete(productId uuid.UUID) error {
	ps.logger.Info("On Delete product service")

	err := ps.productRepo.Delete(productId)
	if err != nil {
		return err
	}

	return nil
}

func (ps productService) UpdateQuantity(productID uuid.UUID, quantity int) error {
	ps.logger.Info("On Update product service")

	err := ps.productRepo.UpdateQuantity(productID, quantity)
	if err != nil {
		return err
	}

	return nil
}
