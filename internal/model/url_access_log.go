package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type UrlAccessLog struct {
	gorm.Model
	ShortUrl  string         `gorm:"column:short_url" json:"short_url"`
	Ip        sql.NullString `gorm:"ip" json:"ip"`
	UserAgent sql.NullString `gorm:"user_agent" json:"user_agent"`
	Referrer  sql.NullString `gorm:"referrer" json:"referrer"`
}

func (u *UrlAccessLog) TableName() string {
	return "url_access_log"
}
