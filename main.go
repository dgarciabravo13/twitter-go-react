package main

import (
	"log"
	"twitter-go-react/bd"
	"twitter-go-react/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
