package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zoetian/short-url/handler"
	"github.com/zoetian/short-url/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "up and running"})
	})

	r.POST("/shortening", func(c *gin.Context) {
		handler.ShortenURL(c)
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitStore()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Server Error: %v", err))
	}
}
