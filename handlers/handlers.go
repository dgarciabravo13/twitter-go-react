package handlers

import (
	"log"
	"net/http"
	"os"
	"twitter-go-react/middlew"
	"twitter-go-react/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto, el Handler y pongo a escuchar el Server
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
