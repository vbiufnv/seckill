package service

import (
	"fmt"
	"log"
	"seckill/dao"
	"seckill/model"
)

func SaveOrderInMysql(userid int, actiNum, orderNum string) {

	var info model.OrderInfo
	info = model.OrderInfo{OrderNum: orderNum,
		UserId:  userid,
		ActiNum: actiNum,
	}

	err := dao.SaveOrder(info)
	if err != nil {
		log.Printf("订单保存失败,err:%v", err)
		return
	}
	fmt.Println("保存成功:", orderNum)
}
