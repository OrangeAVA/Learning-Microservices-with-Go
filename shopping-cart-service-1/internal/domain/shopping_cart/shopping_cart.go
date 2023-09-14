package shoppingcart

import (
	uuid "github.com/satori/go.uuid"
)

type ShoppingCart struct {
	ID     *uuid.UUID `json:"id,omitempty"`
	UserID *uuid.UUID `json:"user_id"`
}

func NewShoppingCart(id *uuid.UUID, userID *uuid.UUID) ShoppingCart {
	return ShoppingCart{
		ID:     id,
		UserID: userID,
	}
}
