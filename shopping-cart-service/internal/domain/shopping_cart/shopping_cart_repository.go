package shoppingcart

import uuid "github.com/satori/go.uuid"

type ShoppingCartRepository interface {
	Create(spc ShoppingCart) error
	GetByUserID(userId *uuid.UUID) (*ShoppingCart, error)
}
