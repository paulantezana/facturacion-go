package queries

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/facturacion-go/models"
    "github.com/paulantezana/facturacion-go/config"
)

func ProdutoQuery() *graphql.Field {
    return &graphql.Field{
        Type: graphql.NewList(models.ProductoType),
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            db := config.GetConnection()
            defer db.Close()

            product := make([]models.Producto, 0)
            if err := db.Find(&product).Error; err != nil {
                return nil, err
            }

            return product, nil
        },
    }
}