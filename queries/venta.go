package queries

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/facturacion-go/models"
    "github.com/paulantezana/facturacion-go/config"
)

func VentaQuery() *graphql.Field {
    return &graphql.Field{
        Type: graphql.NewList(models.VentaType),
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            db := config.GetConnection()
            defer db.Close()

            venta := make([]models.Venta, 0)

            if err := db.Find(&venta).Error; err != nil {
                return nil, err
            }

            return venta, nil
        },
    }
}
