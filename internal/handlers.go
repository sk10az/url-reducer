package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/sk10az/url-reducer/internal/response"
	"gorm.io/gorm"
	"net/http"
)

type handler struct {
	DB *gorm.DB
}

func (h *handler) ReadUrl(c *gin.Context) {
	shortUrl := c.Query("hash")
	orgUrl, err := GetByShort(h.DB, ShortUrl(shortUrl))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error("Not found"))
	}
	c.JSON(http.StatusOK, response.Success(map[string]string{"url": string(orgUrl.OriginUrl)}))
}

func (h *handler) PutUrl(c *gin.Context) {
	orgUrl := c.PostForm("url")
	if orgUrl == "" {
		c.JSON(http.StatusNotFound, response.Error("Bad request"))
		return
	}

	newUrl, err := CreateShort(h.DB, OriginUrl(orgUrl))
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error("Something went wrong"))
		return
	}
	c.JSON(http.StatusOK, response.Success(map[string]string{"hash": string(newUrl.ShortUrl)}))
}

func GetHandler(db *gorm.DB) handler {
	return handler{DB: db}
}