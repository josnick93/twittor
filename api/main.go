package main

import(
	"log"
	"api/db"
	"api/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("No connect to database")
		return
	}
	
	handlers.InitHandlers()
}