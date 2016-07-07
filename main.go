package main

import (
	"log"
	"os"
	"time"

	"github.com/dtannen/sseserver"
	"github.com/garyburd/redigo/redis"
)

func main() {
	s := sseserver.NewServer()
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	if redisHost == "" {
		redisHost = "localhost"
	}
	redisHost = redisHost + ":" + redisPort
	if redisPassword == "" {
		redisPassword = ""
	}
	pool := newPool(redisHost, redisPassword)

	NewBroadcastHandler(pool, s)
	/*
		go func() {
			ticker := time.Tick(time.Duration(1 * time.Second))
			for {
				// wait for the ticker to fire
				t := <-ticker
				// create the message payload, can be any []byte value
				data := []byte(t.Format("3:04:05 pm (MST)"))
				// send a message without an event on the "/time" namespace
				s.Broadcast <- sseserver.SSEMessage{"", data, "/time"}
			}
		}()
	*/

	s.Serve(":8080")
}

func newPool(host, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     2,
		MaxActive:   5,
		IdleTimeout: 5 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					log.Println(err)
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}
