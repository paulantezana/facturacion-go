package mutations

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/config"
	"github.com/paulantezana/facturacion-go/models"
    "time"
)

func CreateProductoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProductoType,
		Args: graphql.FieldConfigArgument{
			"nombre":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"cantidad": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"precio":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := models.Producto{
				Nombre:   p.Args["nombre"].(string),
				Cantidad: int16(p.Args["cantidad"].(int)),
				Precio:   int16(p.Args["precio"].(int)),
			}

			db := config.GetConnection()
			defer db.Close()

            audi := models.Auditoria{
                Fecha: time.Now(),
                Accion: "Crear",
                Tabla: "Productos",
                Anterior: "",
                Nuevo: "",
            }

			// Execute operations
			if err := db.Create(&product).Error; err != nil {
				return nil, err
			}

            if err := db.Create(&audi).Error; err != nil {
                return nil, err
            }

			return product, nil
		},
	}
}

func UpdateProductoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProductoType,
		Args: graphql.FieldConfigArgument{
			"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":   &graphql.ArgumentConfig{Type: graphql.String},
			"cantidad": &graphql.ArgumentConfig{Type: graphql.Int},
			"precio":   &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := models.Producto{
				ID:       uint(p.Args["id"].(int)),
			}

            db := config.GetConnection()
            defer db.Close()

            if db.First(&product).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",product.ID))
            }

            // Optional arguments
            if p.Args["nombre"] != nil {
                product.Nombre = p.Args["nombre"].(string)
            }

            if p.Args["cantidad"] != nil {
                product.Cantidad =  int16(p.Args["cantidad"].(int))
            }

            if p.Args["precio"] != nil {
                product.Precio = int16(p.Args["precio"].(int))
            }

            audi := models.Auditoria{
                Fecha: time.Now(),
                Accion: "Modificar",
                Tabla: "Productos",
                Anterior: "",
                Nuevo: "",
            }

			// Execute operations
			if err := db.Model(&product).Update(product).Error; err != nil {
                return nil, err
            }

            if err := db.Create(&audi).Error; err != nil {
                return nil, err
            }

			return product, nil
		},
	}
}

func DeleteProductoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProductoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			product := models.Producto{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&product).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", product.ID))
			}

            audi := models.Auditoria{
                Fecha: time.Now(),
                Accion: "Eliminar",
                Tabla: "Productos",
                Anterior: "",
                Nuevo: "",
            }

			// Execute operations
			if err := db.Delete(&product).Error; err != nil {
				return nil, err
			}

            if err := db.Create(&audi).Error; err != nil {
                return nil, err
            }

			return product, nil
		},
	}
}
