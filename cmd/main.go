package main

import (
	"github.com/gin-gonic/gin"
	"seckill/api"
	"seckill/dao"
	"seckill/model"
	"seckill/service"
)

func main() {
	//redis
	dao.ConnRedis()
	dao.InitDB()

	defer dao.DB.Close()

	//全局锁
	model.InitChannel()

	go service.Listener()

	r := gin.Default()

	r.POST("/admin/activities/:type", api.CreateActivities)
	//抢购
	r.POST("/user/spike/:actiNum", api.HandlerFunc)
	//查询订单信息
	r.GET("/user/spike/:orderId", api.SearchOrder)

	r.Run(":8088")

}
