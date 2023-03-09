package bd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el objeto de conexión a la bd
var MongoCN = ConectarBD()

// ConectarBD es la función que me permite conectar la BD
func ConectarBD() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando variables de entorno:", err)
		os.Exit(1)
	}
	dbPass := os.Getenv("DBPASSWORD")
	uri := fmt.Sprintf("mongodb+srv://dgb_01:%s@cluster0.kpusu.mongodb.net/?retryWrites=true&w=majority", dbPass)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

// ChequeoConnection es el ping a la bd
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
