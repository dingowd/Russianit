package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var cid string

func init() {
	flag.StringVar(&cid, "cid", "1", "CID")
}

func readLoop(c *websocket.Conn) {
	for {
		_, buff, err := c.ReadMessage()
		if err != nil {
			c.Close()
			break
		}
		log.Printf("received message: %s", string(buff))
	}
}

func main() {
	flag.Parse()
	url := "ws://localhost:33333/ws?cid=" + cid
	wsDialer := websocket.Dialer{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	conn, _, err := wsDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	go readLoop(conn)
	<-ctx.Done()
	fmt.Println("Done")
	conn.Close()
}
