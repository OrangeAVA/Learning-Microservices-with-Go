package main

import (
	"log"

	api "github.com/go-microservices/shopping-cart-service/internal/app/api"

	_ "github.com/go-microservices/shopping-cart-service/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

//	@title			Shopping Cart API
//	@version		1.0
//	@description	APIs to manage sopping cart.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	log.Println("starting shopping-cart services")
	api.LoadApiV1()
}
