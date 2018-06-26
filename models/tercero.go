package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Tercero struct {
	ID             uint       `json:"id" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	Identificacion string     `json:"identificacion" gorm:"type:varchar(32); unique; not null"`
	Nombre         string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Direccion      string     `json:"direccion" gorm:"type:varchar(255)"`
	Telefono       string     `json:"telefono" gorm:"type:varchar(64)"`
	Ventas         []Venta    `json:"ventas" gorm:"foreignkey:ProductoID"`
	Compras        []Compra   `json:"compras" gorm:"foreignkey:ProductoID"`
}

var TerceroType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tercero",
		Fields: graphql.Fields{
			"id":             &graphql.Field{Type: graphql.Int},
			"created_at":     &graphql.Field{Type: graphql.DateTime},
			"updated_at":     &graphql.Field{Type: graphql.DateTime},
			"deleted_at":     &graphql.Field{Type: graphql.DateTime},
			"identificacion": &graphql.Field{Type: graphql.String},
			"nombre":         &graphql.Field{Type: graphql.String},
			"direccion":      &graphql.Field{Type: graphql.String},
			"telefono":       &graphql.Field{Type: graphql.String},
		},
	},
)
