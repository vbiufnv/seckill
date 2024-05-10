package dao

import (
	"context"
	"github.com/go-redis/redis/v8" // 假设你使用的是go-redis/redis客户端库
	"log"
	"strconv"
)

var Client *redis.Client

func ConnRedis() {
	cli := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})

	Client = cli
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("连接失败:", err)
		return
	}

}

func SearchStock(actiNum string) int {
	stockStr, err := Client.HGet(context.Background(), "activities", actiNum).Result()
	if err != nil {
		log.Fatal("获取库存失败", err)
	}
	stock, _ := strconv.Atoi(stockStr)
	return stock
}

func UpdateStock(actiNum string, num int) error {
	stockStr, err := Client.HGet(context.Background(), "activities", actiNum).Result()
	if err != nil {
		log.Println("获取失败", err)
		return err
	}
	stock, _ := strconv.Atoi(stockStr)

	stock -= num
	stockStr = strconv.Itoa(stock)
	_, err = Client.HSet(context.Background(), "activities", actiNum, stockStr).Result()
	if err != nil {
		log.Println("更新失败", err)
		return err
	}

	return nil
}
