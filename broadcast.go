package main

import (
	"log"

	"github.com/dtannen/sseserver"
	"github.com/garyburd/redigo/redis"
)

type broadcaster struct {
	redisPool *redis.Pool
	server    *sseserver.Server
}

func (b *broadcaster) start() {
	r := b.redisPool.Get()
	defer r.Close()
	psc := redis.PubSubConn{Conn: r}
	psc.PSubscribe("*")
	defer psc.Close()
	for {
		switch n := psc.Receive().(type) {
		case redis.PMessage:
			log.Printf("broadcasting: %s\n", string(n.Data))
			msg := sseserver.SSEMessage{Data: n.Data, Namespace: n.Channel}
			b.server.Broadcast <- msg
		case error:
			log.Printf("error: %v\n", n)
			return
		}
	}
}

func NewBroadcastHandler(pool *redis.Pool, s *sseserver.Server) {
	b := &broadcaster{redisPool: pool, server: s}
	go b.start()
	return
}
