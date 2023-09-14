package product_handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-microservices/shopping-cart-service/internal/app/api/dto"
	"github.com/go-microservices/shopping-cart-service/internal/domain/product"
	template_errors "github.com/go-microservices/shopping-cart-service/internal/shared/error"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	Create(product product.Product) error
	GetAll(shoppingCartId uuid.UUID) ([]product.Product, error)
	Delete(productId uuid.UUID) error
	UpdateQuantity(productID uuid.UUID, quantity int) error
}

type productHandler struct {
	logger         *logrus.Logger
	productService ProductService
}

func NewProductHandler(productService ProductService, logger *logrus.Logger) productHandler {
	return productHandler{
		logger:         logger,
		productService: productService,
	}
}

// Create godoc
//
//	@Summary        Adds products to shopping cart
//	@Description    Add product to shopping cart
//	@Tags           Products
//	@Accept         json
//	@Param          request body    dto.ProductReq  true    "shopping cart info"
//	@Success        201
//	@Failure        400 {object}    template_errors.TemplateError
//	@Failure        500 {object}    template_errors.TemplateError
//	@Router         /product [post]
func (p productHandler) Create(c *gin.Context) {
	p.logger.Info("On Create product handler")

	var input dto.ProductReq
	if err := c.BindJSON(&input); err != nil {
		p.logger.Errorf("On Create product handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewBadRequestError("bad request", err))
		return
	}

	if err := input.Validate(); err != nil {
		p.logger.Errorf("On Create product handler - invalid request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewValidationError("invalid params", err))
		return
	}

	newProduct := product.NewProduct(input.Name, input.Quantity, input.Description, input.ShoppingCartID)
	err := p.productService.Create(newProduct)
	var errWraping template_errors.TemplateError
	if errors.As(err, &errWraping) {
		p.logger.Errorf("On Create product handler - error: %v", errWraping.Message)
		c.AbortWithStatusJSON(errWraping.Status, template_errors.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		p.logger.Errorf("On Create product handler - error creating new record: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, template_errors.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// Get godoc
//
//	@Summary        Get shopping cart products
//	@Description    List products in shopping cart
//	@Tags           Products
//	@Produce        json
//	@Param          shoppingCartId  path        string  true    "shopping cart ID"
//	@Success        202             {array}     dto.ProductResp
//	@Failure        400             {object}    template_errors.TemplateError
//	@Failure        404             {object}    template_errors.TemplateError
//	@Failure        500             {object}    template_errors.TemplateError
//	@Router         /product/{shoppingCartId} [get]
func (p productHandler) Get(c *gin.Context) {
	p.logger.Info("On List products handler")

	id, err := uuid.FromString(c.Param("shoppingCartId"))
	if err != nil {
		p.logger.Errorf("On List products handler: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewErrorResponse(err.Error()))
		return
	}

	products, err := p.productService.GetAll(id)
	var errWraping template_errors.TemplateError
	if errors.As(err, &errWraping) {
		p.logger.Errorf("On List products handler - error: %v", errWraping.Message)
		c.AbortWithStatusJSON(errWraping.Status, template_errors.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		p.logger.Errorf("On List products handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	productsResp := []dto.ProductResp{}
	for _, p := range products {
		productsResp = append(productsResp, dto.NewProductResponse(p.ID, p.Name, p.Quantity, p.Description, p.ShoppingCartID))
	}

	c.JSON(http.StatusOK, productsResp)
}

// Delete godoc
//
//	@Summary        Delete shopping cart products by ID
//	@Description    Delete products in shopping cart
//	@Tags           Products
//	@Param          productId   path    string  true    "product ID"
//	@Success        202
//	@Failure        400 {object}    template_errors.TemplateError
//	@Failure        404 {object}    template_errors.TemplateError
//	@Failure        500 {object}    template_errors.TemplateError
//	@Router         /product/{productId} [delete]
func (p productHandler) Delete(c *gin.Context) {
	p.logger.Info("On Delete product handler")

	id, err := uuid.FromString(c.Param("productId"))
	if err != nil {
		p.logger.Errorf("On Delete product handler: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewErrorResponse(err.Error()))
		return
	}

	err = p.productService.Delete(id)
	var errWraping template_errors.TemplateError
	if errors.As(err, &errWraping) {
		p.logger.Errorf("On Delete product handler - error: %v", err)
		c.AbortWithStatusJSON(errWraping.Status, template_errors.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		p.logger.Errorf("On Delete product quantity handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, template_errors.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Update godoc
//
//	@Summary        Update shopping cart products
//	@Description    Update product quatity in shopping cart
//	@Tags           Products
//	@Accept         json
//	@Param          productId   path    string              true    "Product ID"
//	@Param          request     body    dto.ProductQuantity true    "Product quantity request"
//	@Success        202
//	@Failure        400 {object}    template_errors.TemplateError
//	@Failure        404 {object}    template_errors.TemplateError
//	@Failure        500 {object}    template_errors.TemplateError
//	@Router         /product/{productId} [patch]
func (p productHandler) Update(c *gin.Context) {
	p.logger.Info("On Update product quantity handler")

	id, err := uuid.FromString(c.Param("productId"))
	if err != nil {
		p.logger.Errorf("On Update product quantity handler error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewErrorResponse(err.Error()))
		return
	}

	var input dto.ProductQuantity
	if err := c.BindJSON(&input); err != nil {
		p.logger.Errorf("On Update product quantity handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewBadRequestError("bad request", err))
		return
	}

	if err := input.Validate(); err != nil {
		p.logger.Errorf("On Update product quantity handler - invalid request: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, template_errors.NewValidationError("invalid params", err))
		return
	}
	err = p.productService.UpdateQuantity(id, input.Quantity)
	var errWraping template_errors.TemplateError
	if errors.As(err, &errWraping) {
		p.logger.Errorf("On Update product quantity - error: %v", errWraping.Message)
		c.AbortWithStatusJSON(errWraping.Status, template_errors.NewErrorResponse(errWraping.Message))
		return
	}
	if err != nil {
		p.logger.Errorf("On Update product quantity handler - error: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, template_errors.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusAccepted, nil)
}
