package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Usuario struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at"`
	Usuario    string      `json:"usuario" gorm:"type:varchar(128); unique; not null"`
	Email      string      `json:"email" gorm:"type:varchar(128); unique; not null"`
	Nombre     string      `json:"nombre" gorm:"type:varchar(255); not null"`
	Avatar     string      `json:"avatar" gorm:"type:varchar(255); not null"`
	Clave      string      `json:"clave" gorm:"type:varchar(128); not null"`
	Productos  []Producto  `json:"productos" gorm:"foreignkey:UsuarioID"`
	Compras    []Compra    `json:"compras" gorm:"foreignkey:UsuarioID"`
	Ventas     []Venta     `json:"venta" gorm:"foreignkey:UsuarioID"`
	Auditorias []Auditoria `json:"auditorias" gorm:"foreignkey:UsuarioID"`
}

var UsuarioType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Usuario",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"usuario":    &graphql.Field{Type: graphql.String},
            "email":    &graphql.Field{Type: graphql.String},
			"nombre":     &graphql.Field{Type: graphql.String},
			"avatar":     &graphql.Field{Type: graphql.String},
			"clave":      &graphql.Field{Type: graphql.String},
		},
	},
)
