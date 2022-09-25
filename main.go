package main

import (
	"search-service/dotEnv"
	"search-service/routes"
	"search-service/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{dotEnv.GoDotEnvVariable("BFF_URL")}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PublicRoutes(public)

	services.Search("example kalem")

	if err := router.Run(":3006"); err != nil {
		panic(err)
	}
}
