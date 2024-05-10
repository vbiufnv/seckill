package service

import (
	"context"
	"log"
	"seckill/dao"
	"strconv"
	"time"
)

// 创建订单并返回订单号
func CreateOrder(actiNum string, userid int) (string, error) {
	t := time.Now().Unix()
	timeStr := strconv.FormatInt(t, 10)
	str := strconv.Itoa(userid)
	//用户id后三位
	userStr := str[len(str)-3:]
	orderId := timeStr + actiNum + userStr

	//存入redis
	_, err := dao.Client.Set(context.Background(), orderId, str, time.Second*20).Result()

	if err != nil {
		log.Println("订单号存入失败")
		return "", err
	}

	return orderId, nil
}
