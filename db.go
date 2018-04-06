package wxspider

import (
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Article 微信公众号文章
type Article struct {
	ID        uint   `gorm:"primary_key"`
	Title     string // 标题
	Author    string // 作者
	AppName   string // 公众号名称
	AppID     string // 公众号ID
	Cover     string // 文章封面
	Intro     string // 描述
	PubAt     string // 发布时间
	URL       string `gorm:"type:varchar(255);unique_index"` // 微信文章链接地址
	RoundHead string // 圆头像
	OriHead   string // 原头像
	SourceURL string // 公众号原文地址
	PublishAt int64  `sql:"index" default:"0"` //采集器发布时间
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

var db *gorm.DB

//DB 返回 *gorm.DB
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.LogMode(false)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "wxspider.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}