package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handerls to set the port, starts to listen the server */
func InitHandlers() {
	router := mux.NewRouter()


	//routes to handle

	//priority to port
	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "8000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT,handler))
}