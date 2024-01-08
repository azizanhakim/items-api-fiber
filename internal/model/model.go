package model

import (
	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	ID      uint   `json:"typeID" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name"`
	IsHeavy bool   `json:"isHeavy"`
}

type Item struct {
	gorm.Model
	ID     uint    `json:"itemID" gorm:"primaryKey;autoIncrement"`
	Name   string  `json:"name"`
	TypeID uint    `json:"typeID"`
	Type   Type    `json:"type" gorm:"foreignKey:TypeID;references:ID"`
	Price  int     `json:"price"`
	Color  *string `json:"color"`
	Qty    int     `json:"qty"`
}

type Sales struct {
	gorm.Model
	ID    uint   `json:"salesID" gorm:"primaryKey;autoIncrement"`
	Price int    `json:"price"`
	Notes string `json:"notes"`
}

type SalesItem struct {
	gorm.Model
	ID      uint  `json:"salesItemID" gorm:"primaryKey;autoIncrement"`
	ItemID  uint  `json:"itemID"`
	Item    Item  `json:"item" gorm:"foreignKey:ItemID;references:ID"`
	SalesID uint  `json:"salesID"`
	Sales   Sales `json:"sales" gorm:"foreignKey:SalesID;references:ID"`
	Qty     int   `json:"qty"`
	Price   int   `json:"price"`
}
