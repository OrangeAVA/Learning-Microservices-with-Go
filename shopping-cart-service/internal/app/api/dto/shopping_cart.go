package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
)

type ShoppingCartResponse struct {
	ID     *uuid.UUID `json:"id,omitempty"`
	UserID *uuid.UUID `json:"user_id"`
}

func NewShoppingCartResponse(id *uuid.UUID, userID *uuid.UUID) ShoppingCartResponse {
	return ShoppingCartResponse{
		ID:     id,
		UserID: userID,
	}
}

type ShoppingCartRequest struct {
	UserID *uuid.UUID `json:"user_id"`
}

func (sc ShoppingCartRequest) Validate() error {
	return validation.ValidateStruct(&sc,
		validation.Field(&sc.UserID, validation.Required))
}
