package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func pong(c *gin.Context) {
	log.Println("pong handler")
	c.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func (api apiV1) loadRoutes() {
	log.Println("loading Routes")

	r := gin.Default()
	r.GET("/ping", pong)

	v1 := r.Group("/api/v1")
	{
		shoppingCart := v1.Group("shopping-cart")
		{
			shoppingCart.GET("ping", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
			})

			shoppingCart.POST("", api.shoppingCartHandler.Create)
			shoppingCart.GET(":userID", api.shoppingCartHandler.Get)

		}

		product := v1.Group("/product")
		{
			product.GET("ping", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
			})

			product.POST("", api.productHandler.Create)
			product.PATCH(":productId", api.productHandler.Update)
			product.GET(":shoppingCartId", api.productHandler.Get)
			product.DELETE(":productId", api.productHandler.Delete)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		log.Println("can not start service")
	}

}
