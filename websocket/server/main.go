package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type clients map[string]*websocket.Conn

var (
	clientsMutex sync.Mutex
	hc           int
	counter      int
	hubs         = make([]clients, 0)
	upgrader     = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

func writeConn(conn *websocket.Conn, msg string) {
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.Fatal(err)
	}
}

func writeLoop(ch <-chan string) {
	for str := range ch {
		comm := strings.Split(str, " ")
		comm[0] = strings.ToLower(comm[0])
		comm[1] = strings.TrimPrefix(comm[1], "--")
		comm[1] = strings.Trim(comm[1], "\n")
		switch {
		case comm[0] == "send":
			num, err := strconv.Atoi(comm[1])
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if num > len(hubs)-1 {
				fmt.Fprintln(os.Stdout, "hub not exist")
				continue
			}

			for k := range hubs[num] {
				writeConn(hubs[num][k], "Hello, hub "+comm[1])
			}
		case comm[0] == "sendc":
			var conn *websocket.Conn
			exist := false
			for k := range hubs {
				if v, InMap := hubs[k][comm[1]]; InMap {
					conn = v
					exist = true
					break
				}
			}
			if exist {
				writeConn(conn, "Hello, cid "+comm[1])
			} else {
				fmt.Fprintln(os.Stdout, "cid not exist")
			}
		}
	}
}

func fillHubs(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	if cid == "" {
		http.Error(w, "cid not passed", 400)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Printf("add connection, cid=%s", cid)

	clientsMutex.Lock()
	if len(hubs[counter]) == hc {
		hubs = append(hubs, make(clients))
		counter++
		hubs[counter][cid] = conn
	} else {
		hubs[counter][cid] = conn
	}
	for k := range hubs {
		fmt.Fprint(os.Stdout, "hub ", k, " cids(")
		for l := range hubs[k] {
			fmt.Fprint(os.Stdout, l, "  ")
		}
		fmt.Fprint(os.Stdout, ")\n")
	}
	clientsMutex.Unlock()
}

func init() {
	flag.IntVar(&hc, "hc", 2, "number of clients in a hub")
}

func main() {
	flag.Parse()
	hubs = append(hubs, make(clients))
	ch := make(chan string)
	go writeLoop(ch)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			ch <- text
		}
	}()

	http.HandleFunc("/ws", fillHubs)
	log.Fatal(http.ListenAndServe("localhost:33333", nil))
}
