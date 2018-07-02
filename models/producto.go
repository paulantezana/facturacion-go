package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Producto struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nombre    string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Cantidad  int16      `json:"cantidad" gorm:"not null"`
	Precio    int16      `json:"precio" gorm:"not null"`
	UsuarioID uint       `json:"usuario_id"`
	Ventas    []Venta    `json:"ventas" gorm:"foreignkey:ClienteID"`
	Compras   []Compra   `json:"compras" gorm:"foreignkey:ProveedorID"`
}

var ProductoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Producto",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"nombre":     &graphql.Field{Type: graphql.String},
			"cantidad":   &graphql.Field{Type: graphql.Int},
			"precio":     &graphql.Field{Type: graphql.Int},
		},
	},
)
