package model

import "gorm.io/gorm"

type OrderInfo struct {
	gorm.Model
	UserId   int
	ActiNum  string
	OrderNum string
	Cancle   int
}
