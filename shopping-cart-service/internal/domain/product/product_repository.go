package product

import uuid "github.com/satori/go.uuid"

type ProductRepository interface {
	Create(product Product) error
	Get(shoppingCartID uuid.UUID) ([]Product, error)
	Delete(productID uuid.UUID) error
	UpdateQuantity(productID uuid.UUID, quantity int) error
}
