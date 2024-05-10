package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"seckill/dao"
	"seckill/model"
	"seckill/service"
	"time"
)

func HandlerFunc(c *gin.Context) {
	userid := 12233
	actiNum := c.Param("actiNum")[1:]

	<-model.Done

	if time.Now().Unix() > model.Acti.EndTime.Unix() {
		c.JSON(http.StatusOK, "活动结束")
		return
	}

	var stock int
	if stock = dao.SearchStock(actiNum); stock <= 0 {
		c.JSON(http.StatusOK, "售空")
		return
	}
	fmt.Println(stock)

	orderId, err := service.CreateOrder(actiNum, userid)

	if err != nil {
		log.Println(err)
		return
	}

	ok := service.UpdateStock(actiNum, 1)
	if !ok {
		c.JSON(http.StatusOK, "更新库存失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "成功",
		"orderid": orderId,
	})

	model.Done <- 1

	go service.SaveOrderInMysql(userid, actiNum, orderId)

}
