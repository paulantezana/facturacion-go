package queries

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/facturacion-go/models"
    "github.com/paulantezana/facturacion-go/config"
)

func TerceroQuery() *graphql.Field {
    return &graphql.Field{
        Type: graphql.NewList(models.TerceroType),
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            db := config.GetConnection()
            defer db.Close()

            tercero := make([]models.Tercero, 0)

            if err := db.Find(&tercero).Error; err != nil {
                return nil, err
            }

            return tercero, nil
        },
    }
}
