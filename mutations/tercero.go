package mutations

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/config"
	"github.com/paulantezana/facturacion-go/models"
)

func CreateTerceroMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.TerceroType,
		Args: graphql.FieldConfigArgument{
            "identificacion": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "nombre":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "direccion":      &graphql.ArgumentConfig{Type: graphql.String},
            "telefono":       &graphql.ArgumentConfig{Type: graphql.String},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			tercero := models.Tercero{
				Identificacion: p.Args["identificacion"].(string),
                Nombre: p.Args["nombre"].(string),
			}

            // Optional arguments
            if p.Args["direccion"] != nil {
                tercero.Direccion = p.Args["direccion"].(string)
            }

            if p.Args["telefono"] != nil {
                tercero.Telefono = p.Args["telefono"].(string)
            }

            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&tercero).Error; err != nil {
                return nil, err
            }

			return tercero, nil
		},
	}
}

func UpdateTerceroMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.TerceroType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
            "identificacion": &graphql.ArgumentConfig{Type: graphql.String},
            "nombre":         &graphql.ArgumentConfig{Type: graphql.String},
            "direccion":      &graphql.ArgumentConfig{Type: graphql.String},
            "telefono":       &graphql.ArgumentConfig{Type: graphql.String},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            tercero := models.Tercero{
				ID: uint(p.Args["id"].(int)),
			}

            db := config.GetConnection()
            defer db.Close()

            if db.First(&tercero).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",tercero.ID))
            }

            // Optional arguments
            if p.Args["identificacion"] != nil {
                tercero.Identificacion = p.Args["identificacion"].(string)
            }

            if p.Args["nombre"] != nil {
                tercero.Nombre = p.Args["nombre"].(string)
            }

            if p.Args["direccion"] != nil {
                tercero.Direccion = p.Args["direccion"].(string)
            }

            if p.Args["telefono"] != nil {
                tercero.Telefono = p.Args["telefono"].(string)
            }

            // Execute operations
            if err := db.Model(&tercero).Update(tercero).Error; err != nil {
                return nil, err
            }

            return tercero, nil
		},
	}
}

func DeleteTerceroMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.TerceroType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			tercero := models.Tercero{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&tercero).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", tercero.ID))
			}

			// Execute operations
			if err := db.Delete(&tercero).Error; err != nil {
				return nil, err
			}

			return tercero, nil
		},
	}
}
