package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Venta struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Fecha      time.Time  `json:"fecha"`
	ClienteID  uint       `json:"cliente_id"`
	ProductoID uint       `json:"producto_id"`
	Cantidad   int16      `json:"cantidad"`
	Valor      int16      `json:"valor"`
	UsuarioID  uint       `json:"usuario_id"`
}

var VentaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Venta",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"created_at":  &graphql.Field{Type: graphql.DateTime},
			"updated_at":  &graphql.Field{Type: graphql.DateTime},
			"deleted_at":  &graphql.Field{Type: graphql.DateTime},
			"fecha":       &graphql.Field{Type: graphql.DateTime},
			"cliente_id":  &graphql.Field{Type: graphql.Int},
			"producto_id": &graphql.Field{Type: graphql.Int},
			"cantidad":    &graphql.Field{Type: graphql.Int},
			"valor":       &graphql.Field{Type: graphql.Int},
			"usuario_id":  &graphql.Field{Type: graphql.Int},
		},
	},
)
