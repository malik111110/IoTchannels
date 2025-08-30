package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"chanIoT/pkg/event"
	"chanIoT/pkg/simulator"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan event.Event)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	sim := simulator.New()

	go func() {
		for evt := range sim.Events {
			broadcast <- evt
		}
	}()

	fs := http.FileServer(http.Dir("./dashboard/build"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	http.HandleFunc("/api/start", func(w http.ResponseWriter, r *http.Request) {
		sim.Start()
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api/stop", func(w http.ResponseWriter, r *http.Request) {
		sim.Stop()
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		status := make(map[string]bool)
		status["running"] = sim.Running
		json.NewEncoder(w).Encode(status)
	})

	go func() {
		fmt.Println("HTTP server started on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("ListenAndServe: ", err)
		}
	}()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	fmt.Println("\nShutting down...")
}
