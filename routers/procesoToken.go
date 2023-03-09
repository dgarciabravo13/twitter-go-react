package routers

import (
	"errors"
	"os"
	"strings"
	"twitter-go-react/bd"
	"twitter-go-react/models"

	"github.com/dgrijalva/jwt-go"
)

// Email valor de Email usado en todos los endPoints
var Email string

// IDUsuario es el ID devuelto del modelo, que se usará en todos los Endpoints
var IDUsuario string

// ProcesoToken proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	key := os.Getenv("JWTKEYSECRET")
	miClave := []byte(key)
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
