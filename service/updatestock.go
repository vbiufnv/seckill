package service

import (
	"log"
	"seckill/dao"
)

func UpdateStock(actiNum string, num int) bool {
	err := dao.UpdateStock(actiNum, num)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
