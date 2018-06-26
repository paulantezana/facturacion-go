package models

import "time"

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
