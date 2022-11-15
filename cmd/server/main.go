// Package main is used to build the app
package main

import (
	"fmt"
	"github.com/OmgAbear/aoe/internal/presentation"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	presentation.RegisterRoutes(router)
	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Println(err)
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			dt := time.Now()
			fmt.Println("Current date and time is: ", dt.String())
		}
	}()
	log.Println("started")
}
