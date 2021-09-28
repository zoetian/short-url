package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zoetian/short-url/shortener"
	"github.com/zoetian/short-url/store"
)

type ShortenUrlRequest struct {
	URL    string `json:"url" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

func ShortenURL(c *gin.Context) {
	var shortenRequest ShortenUrlRequest
	err := c.ShouldBindJSON(&shortenRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortUrl(shortenRequest.URL, shortenRequest.UserID)
	store.SaveUrlMapping(shortUrl, shortenRequest.URL, shortenRequest.UserID)

	c.JSON(http.StatusOK, gin.H{
		"message":   "successfully created short url",
		"short_url": "http://localhost:8080/" + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initUrl := store.FetchOriginalUrl(shortUrl)
	c.Redirect(302, initUrl)
}
