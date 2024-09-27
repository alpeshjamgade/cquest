package app

import (
	"context"
	"cquest/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	config.LoadConf()

	ctx := context.Background()

	r := mux.NewRouter()
	log.Printf("Starting server on port %s", "8080")
	go func() {
		http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r)
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
}
