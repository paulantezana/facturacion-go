package models

import (
    "time"
    "github.com/graphql-go/graphql"
)

type Auditoria struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Fecha     time.Time  `json:"fecha"`
	Accion    string     `json:"accion" gorm:"type:varchar(64); not null"`
	Tabla     string     `json:"tabla" gorm:"type:varchar(64); not null"`
	Anterior  string     `json:"anterior"`
	Nuevo     string     `json:"nuevo"`
	UsuarioID uint       `json:"usuario_id"`
}

var AuditoriaType = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Auditoria",
        Fields: graphql.Fields{
            "id":           &graphql.Field{Type: graphql.Int},
            "created_at":   &graphql.Field{Type: graphql.DateTime},
            "updated_at":   &graphql.Field{Type: graphql.DateTime},
            "deleted_at":   &graphql.Field{Type: graphql.DateTime},
            "fecha":        &graphql.Field{Type: graphql.DateTime},
            "accion": &graphql.Field{Type: graphql.String},
            "tabla":  &graphql.Field{Type: graphql.String},
            "anterior":     &graphql.Field{Type: graphql.String},
            "nuevo":        &graphql.Field{Type: graphql.String},
        },
    },
)
