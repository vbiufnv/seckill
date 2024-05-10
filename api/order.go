package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill/dao"
)

func SearchOrder(c *gin.Context) {
	idStr := c.Param("orderId")[1:]

	_, err := dao.Client.Get(context.Background(), idStr).Result()
	if err != nil {
		c.JSON(http.StatusOK, "订单已过期")
	} else {
		c.JSON(http.StatusOK, "查询成功")

	}

}
