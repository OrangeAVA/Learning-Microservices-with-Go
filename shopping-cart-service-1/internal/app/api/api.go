package api

import (
	"github.com/gin-gonic/gin"
	product_handler "github.com/go-microservices/shopping-cart-service/internal/app/api/handler/product"
	scHandler "github.com/go-microservices/shopping-cart-service/internal/app/api/handler/shopping_cart"
	"github.com/go-microservices/shopping-cart-service/internal/domain/product"
	shoppingcart "github.com/go-microservices/shopping-cart-service/internal/domain/shopping_cart"
	"github.com/go-microservices/shopping-cart-service/internal/infrastructure/database"
	"github.com/go-microservices/shopping-cart-service/internal/infrastructure/persistence"
	"github.com/go-microservices/shopping-cart-service/internal/shared/logger"
)

type ShoppingCartHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

type ProductHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type apiV1 struct {
	shoppingCartHandler ShoppingCartHandler
	productHandler      ProductHandler
}

func NewApiV1(shoppingCartHandler ShoppingCartHandler, productHandler ProductHandler) apiV1 {
	return apiV1{
		shoppingCartHandler: shoppingCartHandler,
		productHandler:      productHandler,
	}
}

func LoadApiV1() {
	logger := logger.NewLogger()
	logger.Info("loading shopping-cart API")
	dbConfig, err := database.NewConfig()
	if err != nil {
		logger.Errorf("error loading configuration: %v", err)
	}
	dbConnection, err := database.NewDatabseConnection(dbConfig)
	if err != nil {
		logger.Errorf("error connecting to database: %v", err)
	}
	defer dbConnection.Close()
	concreteShoppinCartRepoImpl := persistence.NewshoppingCartMySQLRepo(dbConnection, logger)
	shoppingCartSvc := shoppingcart.NewShoppingCartService(concreteShoppinCartRepoImpl, logger)
	spHandler := scHandler.NewShoppingCartHandler(shoppingCartSvc, logger)
	concreteProductRepoImp := persistence.NewProductMySQLRepo(dbConnection, logger)
	producScv := product.NewProductService(concreteProductRepoImp, logger)
	pHandler := product_handler.NewProductHandler(producScv, logger)
	apiV1 := NewApiV1(spHandler, pHandler)
	apiV1.loadRoutes()
}
