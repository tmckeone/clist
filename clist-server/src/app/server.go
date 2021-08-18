package main

import (
	"clist-server/src/api"
	"log"
	"net/http"
)

func main() {
	mux := api.MasterHandler()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
