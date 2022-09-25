package controllers

import (
	"fmt"
	"net/http"
	"search/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SearchQuery struct {
	SearchQuery string
}

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var searchQuery SearchQuery
		if err := ctx.ShouldBindBodyWith(&searchQuery, binding.JSON); err != nil {
			fmt.Println("body:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		products, getErr := services.Search(searchQuery.SearchQuery)
		if getErr != nil {
			fmt.Println("GetAllProducts:", getErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"products": products})
	}
}

func AddProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product services.Product
		if err := ctx.ShouldBindBodyWith(&product, binding.JSON); err != nil {
			fmt.Println("body:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		result, err := services.AddProduct(product)
		if err != nil {
			fmt.Println(err)
		}

		ctx.JSON(http.StatusOK, gin.H{"message": result.Result})
	}
}

func AddAllProducts() {
	products, err := services.GetAllProducts()
	if err != nil {
		fmt.Println(err)
	}

	for _, j := range products {
		result, err := services.AddProduct(j)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result.Shards)
	}
}
