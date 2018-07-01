package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/config"
	"github.com/paulantezana/facturacion-go/models"
)

func UsuarioQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.UsuarioType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()

			us := make([]models.Usuario, 0)

			if err := db.Find(&us).Error; err != nil {
				return nil, err
			}

			return us, nil
		},
	}
}
