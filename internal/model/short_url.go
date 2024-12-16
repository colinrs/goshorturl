package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type ShortUrl struct {
	gorm.Model
	ShortUrl    string         `gorm:"column:short_url" json:"short_url"`
	OriginUrl   string         `gorm:"column:origin_url" json:"origin_url"`
	Description sql.NullString `gorm:"column:description" json:"description"`
	UrlType     int            `gorm:"column:url_type" json:"url_type"` // type=1:system  type=2:custom
	ExpireAt    sql.NullTime   `gorm:"column:expire_at" json:"expire_at"`
}

func (ShortUrl) TableName() string {
	return "short_url"
}
