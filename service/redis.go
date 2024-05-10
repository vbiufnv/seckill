package service

import (
	"context"
	"fmt"
	"log"
	"seckill/dao"
	"seckill/model"
	"strconv"
)

func InfoToRedis(acti *model.Activities) {
	stock := strconv.Itoa(acti.Stock)
	err := dao.Client.HSet(context.Background(), "activities", acti.ActiNum, stock).Err()
	if err != nil {
		log.Println(err)
		return
	}
}

func Listener() {
	pubsub := dao.Client.PSubscribe(context.Background(), "__keyevent@0__:expired")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		err := dao.UpdateStock(model.Acti.ActiNum, -1)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("库存更新成功")
		err = dao.UpdateOrder(msg.Payload)
		if err != nil {
			log.Println(err)
			return
		}

	}

}
