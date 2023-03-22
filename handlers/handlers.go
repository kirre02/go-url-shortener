package handlers

import (
	"go-url-shortener/shortener"
	"go-url-shortener/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlCreationReq struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(ctx *gin.Context) {
	var creationRequest UrlCreationReq
	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrl(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	ctx.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
