package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"

	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//	Conexao DB
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//	Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//	Camada Usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//	Camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

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
