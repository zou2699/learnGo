package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	// connect redis
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	log.Info(pong)

	// redis
	// Set
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		log.Error(err)
	}
	// Get
	val, err := client.Get("key").Result()
	if err != nil {
		log.Error(err)
	}
	log.Info("key", val)

	//time.Sleep(1*time.Second)
	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		log.Error("key2 does not exist")
	} else if err != nil {
		log.Panic(err)
	} else {
		log.Info("key2", val2)
	}

	//test
	go func(client *redis.Client) {
		for i := 0; i < 100; i++ {
			client.RPush("test", i)
			log.Info("i:", i)
		}
	}(client)
	time.Sleep(1 * time.Second)
}
