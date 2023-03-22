package main

import (
	"fmt"
	"go-url-shortener/handlers"
	"go-url-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleShortUrlRedirect(c)
	})

	// Note store initialization happens here
	store.InitializeStore()

	err := router.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
