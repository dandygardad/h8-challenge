package main

import (
	"challenge10/config"
	"challenge10/model/web"
	"challenge10/repository"
	"challenge10/routes"
	"challenge10/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	err := config.InitGorm()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(config.NewGorm.DB)
	serv := service.NewService(repo)

	newRouter := gin.New()
	routes.BookRouter(newRouter, serv)
	newRouter.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, web.BookResponse{Message: "Page not found"})
	})

	port := os.Getenv("PORT")
	err = newRouter.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
