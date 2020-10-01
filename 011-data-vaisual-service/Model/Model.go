package Model

import (
	"database/sql"
	"time"
)

/**
统一返回对象
*/
type R struct {
	Code    int         `json:"code" default:0`
	Data    interface{} `json:"data"`
	Success bool        `json:"success" default:1`
	Msg     string      `json:"msg"`
}
type BaseModel struct {
	ID        uint `gorm:"primarykey" json:"id" gorm:"AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type VisualMap struct {
	BaseModel
	//Id uint `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Name string `json:"name"`
	Data string `json:"data" gorm:"type:text"`
	//CreatedAt    time.Time
	//UpdatedAt    time.Time
}

type Visual struct {
	BaseModel
	Title         string `json:"title"`
	BackgroundUrl string `json:"backgroundUrl"`
	Category      string `json:"category"`
	Password      string `json:"password"`
}

type VisualCategory struct {
	BaseModel
	CategoryKey   string `json:"categoryKey"`
	CategoryValue string `json:"categoryValue"`
}

type VisualConfig struct {
	BaseModel
	VisualId  uint64 `json:"visualId"`
	Detail    string `json:"detail"`
	Component string `json:"component" gorm:"size:2000"`
}
