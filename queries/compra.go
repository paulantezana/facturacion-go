package queries

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/facturacion-go/models"
    "github.com/paulantezana/facturacion-go/config"
)

func CompraQuery() *graphql.Field {
    return &graphql.Field{
        Type: graphql.NewList(models.CompraType),
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            db := config.GetConnection()
            defer db.Close()

            compra := make([]models.Compra, 0)

            if err := db.Find(&compra).Error; err != nil {
                return nil, err
            }
            return compra, nil
        },
    }
}