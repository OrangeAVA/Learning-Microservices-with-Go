package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	uuid "github.com/satori/go.uuid"
)

type ProductReq struct {
	Name           string     `json:"name"`
	Quantity       int        `json:"quantity"`
	Description    string     `json:"description"`
	ShoppingCartID *uuid.UUID `json:"shoppingCartId"`
}

func (p ProductReq) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Quantity, validation.Required, validation.Min(1)),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.ShoppingCartID, validation.Required),
	)
}

type ProductResp struct {
	ID             *uuid.UUID `json:"id,omitempty"`
	Name           string     `json:"name"`
	Quantity       int        `json:"quantity"`
	Description    string     `json:"description"`
	ShoppingCartID *uuid.UUID `json:"shoppingCartId"`
}

func NewProductResponse(
	id *uuid.UUID,
	name string,
	quantity int,
	description string,
	shoppingCartID *uuid.UUID,
) ProductResp {
	return ProductResp{
		ID:             id,
		Name:           name,
		Quantity:       quantity,
		Description:    description,
		ShoppingCartID: shoppingCartID,
	}
}

type ProductQuantity struct {
	Quantity int `json:"quantity"`
}

func (pq ProductQuantity) Validate() error {
	return validation.ValidateStruct(&pq,
		validation.Field(&pq.Quantity, validation.Required, validation.Min(1)),
	)
}
