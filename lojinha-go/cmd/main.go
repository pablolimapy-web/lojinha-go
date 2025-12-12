package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablolimapy-web/lojinha-go/lojinha-go/controller"
	"github.com/pablolimapy-web/lojinha-go/lojinha-go/db"
	"github.com/pablolimapy-web/lojinha-go/lojinha-go/repository"
	"github.com/pablolimapy-web/lojinha-go/lojinha-go/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")

}
