package routers

import (
	"io"
	"net/http"
	"os"
	"twitter-go-react/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Banner no encontrado", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar el banner", http.StatusBadRequest)
	}

}
