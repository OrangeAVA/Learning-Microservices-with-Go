package product

import (
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID             *uuid.UUID
	Name           string
	Quantity       int
	Description    string
	ShoppingCartID *uuid.UUID
}

func NewProduct(name string, quantity int, description string,
	shoppingCartID *uuid.UUID) Product {
	return Product{
		Name:           name,
		Quantity:       quantity,
		Description:    description,
		ShoppingCartID: shoppingCartID,
	}
}
