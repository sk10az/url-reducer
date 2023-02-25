package internal

import (
	"encoding/base32"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type OriginUrl string
type ShortUrl string

type URL struct {
	ID        uint `gorm:"primaryKey"`
	OriginUrl OriginUrl
	ShortUrl  ShortUrl `gorm:"index"`
}

func createHash(id uint) ShortUrl {
	hash := base32.StdEncoding.EncodeToString([]byte(fmt.Sprint(id)))
	return ShortUrl(strings.ReplaceAll(hash, "=", ""))
}

func CreateShort(db *gorm.DB, originUrl OriginUrl) (URL, error) {
	url := URL{OriginUrl: originUrl, ShortUrl: ""}
	db.Create(&url)
	url.ShortUrl = createHash(url.ID)
	db.Save(&url)
	return url, nil
}

func GetByShort(db *gorm.DB, shortUrl ShortUrl) (URL, error) {
	var url URL
	result := db.Take(&url, "short_url = ?", shortUrl)
	if result.Error != nil {
		return url, result.Error
	}
	return url, nil
}
