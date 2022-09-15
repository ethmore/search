package routes

import (
	"search/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.POST("/search", controllers.Search())
	g.POST("/addProduct", controllers.AddProduct())
}
