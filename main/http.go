package main

import (
	"log"
	"net/http"
)
import "socken/server"

func main() {
	http.Handle("/", server.StaticServer)
	http.HandleFunc("/newplayer", server.NewPlayer)
	http.HandleFunc("/player/", server.Player)

	http.HandleFunc("/board", server.Board)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
