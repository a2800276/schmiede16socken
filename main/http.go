package main

import (
	"log"
	"net/http"
)
import "socken/server"

func main() {
	http.Handle("/", server.StaticServer)
	http.Handle("/sharedboard", server.SharedBoard)
	http.Handle("/playerboard", server.PlayerBoard)

	//	http.HandleFunc("/board", server.Board)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
