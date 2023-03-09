package main

import (
	"fmt"
	"log"
	"os"
	"twitter-go-react/bd"
	"twitter-go-react/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando variables de entorno:", err)
		os.Exit(1)
	}
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
