package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill/model"
	"seckill/service"
	"strconv"
	"time"
)

// 创建活动 保存在数据库
func CreateActivities(c *gin.Context) {
	//商户id
	adminId := 1122
	t := c.Param("type")[1:]

	actiNum := time.Now().Unix()

	//假定
	model.Acti = model.Activities{
		AdminId:   adminId,
		ActiNum:   strconv.FormatInt(actiNum, 10) + t + strconv.Itoa(adminId),
		Stock:     10000,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Hour * 2),
	}

	//放入redis
	service.InfoToRedis(&model.Acti)

	//返回活动编号
	c.JSON(http.StatusOK, gin.H{
		"ActiveNum": model.Acti.ActiNum,
	})

}
