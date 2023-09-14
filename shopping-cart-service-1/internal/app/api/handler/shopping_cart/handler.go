package shopping_cart_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-microservices/shopping-cart-service/internal/app/api/dto"
	shoppingcart "github.com/go-microservices/shopping-cart-service/internal/domain/shopping_cart"
	template_errors "github.com/go-microservices/shopping-cart-service/internal/shared/error"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	_ "github.com/swaggo/swag/example/celler/httputil"
)

type ShoppingCartService interface {
	Create(spc shoppingcart.ShoppingCart) error
	GetByUserID(userId *uuid.UUID) (*shoppingcart.ShoppingCart, error)
}

type shoppingCartHandler struct {
	logger              *logrus.Logger
	shoppingCartService ShoppingCartService
}

func NewShoppingCartHandler(
	shoppingCartService ShoppingCartService, logger *logrus.Logger) shoppingCartHandler {
	return shoppingCartHandler{
		logger:              logger,
		shoppingCartService: shoppingCartService,
	}
}

// Create godoc
//
//	@Summary		Creates Shopping Cart
//	@Description	Create a new Shopping Cart
//	@Tags			Shopping Cart
//	@Accept			json
//	@Param			request	body	dto.ShoppingCartRequest	true	"shopping cart info"
//	@Success		201
//	@Failure		400	{object}	template_errors.TemplateError
//	@Failure		500	{object}	template_errors.TemplateError
//	@Router			/shopping-cart [post]
func (sc shoppingCartHandler) Create(c *gin.Context) {
	sc.logger.Info("On create sopping cart")

	var input dto.ShoppingCartRequest
	if err := c.BindJSON(&input); err != nil {
		sc.logger.Errorf("On create shopping cart handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewBadRequestError("bad request", err))
		return
	}

	if err := input.Validate(); err != nil {
		sc.logger.Errorf("On create shopping cart handler - invalid request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewValidationError("invalid values", err))
		return
	}

	err := sc.shoppingCartService.Create(shoppingcart.NewShoppingCart(nil, input.UserID))
	var errWraping template_errors.TemplateError
	if errors.As(err, &errWraping) {
		sc.logger.Errorf("On create shopping cart handler - error: %v", errWraping.Message)
		c.AbortWithStatusJSON(errWraping.Status, template_errors.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		sc.logger.Errorf("On create shopping cart handler - error creating new record: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, template_errors.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// Get godoc
//
//	@Summary		Get User Shopping Cart
//	@Description	Returns a Shopping Cart
//	@Tags			Shopping Cart
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	dto.ShoppingCartResponse
//	@Failure		400	{object}	template_errors.TemplateError
//	@Failure		404	{object}	template_errors.TemplateError
//	@Failure		500	{object}	template_errors.TemplateError
//	@Router			/shopping-cart/{id} [get]
func (sc shoppingCartHandler) Get(c *gin.Context) {
	sc.logger.Info("On get shorpping cart")

	id, err := uuid.FromString(c.Param("userID"))
	if err != nil {
		sc.logger.Errorf("On get shorpping cart handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewErrorResponse(err.Error()))
		return
	}

	resp, err := sc.shoppingCartService.GetByUserID(&id)
	var errWraping template_errors.TemplateError
	if errors.As(err, &errWraping) {
		sc.logger.Errorf("On get shorpping cart handler - error: %v", errWraping.Message)
		c.AbortWithStatusJSON(errWraping.Status, template_errors.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		sc.logger.Errorf("On get shorpping cart handler: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, template_errors.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.NewShoppingCartResponse(resp.ID, resp.UserID))
}
