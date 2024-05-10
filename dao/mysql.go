package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"seckill/model"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:147258aa@tcp(127.0.0.1:3306)/dyt?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("数据库打开失败，err:%v", err)
		return
	}
	DB = db

	DB.AutoMigrate(&model.OrderInfo{})

	if DB.DB().Ping() != nil {
		log.Println(err)
		return
	}
}

func SaveOrder(info model.OrderInfo) error {
	err := DB.Create(&info).Error
	return err
}

func UpdateOrder(orderNum string) error {
	err := DB.Model(&model.OrderInfo{}).Where("order_num=?", orderNum).Updates(model.OrderInfo{Cancle: 1}).Error
	return err
}
