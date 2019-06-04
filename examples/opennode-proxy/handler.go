package main

import (
	"fmt"
	"net/http"

	"github.com/edouardparis/opennode-go/opennode"
)

func NewHandler(client *opennode.Client) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/hello", pongHandler)
	return router
}

func pongHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "world")
}
