package service

import (
	"log"
	"net/http"
)

// StartWebServer -
func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port" + port)
		log.Println("Error: ", err.Error())
	}
}
