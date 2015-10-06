package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("upgrading connection")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("processing messages")
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Println(messageType, p)
		if err = conn.WriteMessage(messageType, p); err != nil {
			log.Fatal(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/events", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
