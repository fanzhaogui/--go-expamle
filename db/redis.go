package db

import (
	"api.service.com/config"
	"fmt"
	"github.com/go-redis/redis"
	//"time"
)

var redisClient *redis.Client

func InitRedis() {
	conf := config.GetConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("redis.addr"),
		Password: conf.GetString("redis.password"),
		DB:       conf.GetInt("redis.db"), // use default DB
	})
	_, err := client.Ping().Result()

	if err != nil {
		panic("redis error:" + fmt.Sprintf("%v", err))
	}
	redisClient = client
}

func GetRedis() *redis.Client {
	return redisClient
}

func Test_Receive() {
	InitRedis()
	redisdb := GetRedis()
	pubsub := redisdb.Subscribe("test", "here")
	defer pubsub.Close()

	for {
		// ReceiveTimeout is a low level API. Use ReceiveMessage instead.
		//msgi, err := pubsub.ReceiveTimeout(time.Second * 30)

		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			break
		}
		fmt.Println("received", msg.Payload, "from", msg.Channel)

	}

	// sent message to 1 redisdb
	// received hello from test
}
