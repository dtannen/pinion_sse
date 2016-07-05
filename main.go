package main

import (
	"time"

	"github.com/dtannen/sseserver"
)

func main() {
	s := sseserver.NewServer()
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
	s.Serve(":8001")
}
