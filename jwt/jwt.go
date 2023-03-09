package jwt

import (
	"os"
	"time"
	"twitter-go-react/models"

	"github.com/dgrijalva/jwt-go"
)

// GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {
	key := os.Getenv("JWTKEYSECRET")
	miClave := []byte(key)

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"ubicacion":        t.Ubicacion,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
