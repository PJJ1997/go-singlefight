package main

import (
	"log"
	"time"

	"golang.org/x/sync/singleflight"
)

// 每次使用singleflight进行函数调用，就会查看Group里有没有已注册的相同调用。
// 如果有则等待调用完成，并返回其结果。如果没有，则注册一个，并且执行函数。
// 执行完成后，将注册的调用删掉

var singleCache singleflight.Group

func main() {
	cacheKey := "cacheKey"
	for i := 1; i < 10; i++ { //模拟多个协程同时请求
		go func(requestID int) {
			value, _ := getData(requestID, cacheKey)
			log.Printf("request %v get value: %v", requestID, value)
		}(i)
	}
	time.Sleep(20 * time.Second)
}

func getData(requestID int, cacheKey string) (string, error) {
	log.Printf("request %v start to get data from mysql...", requestID)
	value, _, _ := singleCache.Do(cacheKey, func() (ret interface{}, err error) {
		log.Printf("request %v is getting data from mysql...", requestID)
		time.Sleep(3 * time.Second)
		log.Printf("request %v get data from mysql success!", requestID)
		return "VALUE", nil
	})
	return value.(string), nil
}
