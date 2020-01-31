package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhijitWakchaure/lanmusic/gosrc/music"
	"github.com/gorilla/handlers"
)

func main() {
	shutdown := make(chan bool)
	musicRouter := music.NewRouter()

	// these two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"})

	// launch server with CORS validations
	go func() {
		log.Fatal(http.ListenAndServe(":9000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(musicRouter)))
		shutdown <- true
	}()
	fmt.Printf("LanMusic api server started on port 9000\n")
	fmt.Println("You can visit the web interface at:  http://localhost")

	<-shutdown
}
