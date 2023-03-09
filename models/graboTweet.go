package models

import "time"

// GraboTweet es el formato o escructura que tendrá nuestro Tweet en la BBDD
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
