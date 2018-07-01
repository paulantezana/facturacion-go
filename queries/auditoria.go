package queries

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/facturacion-go/models"
    "github.com/paulantezana/facturacion-go/config"
)

func AuditoriaQuery() *graphql.Field {
    return &graphql.Field{
        Type: graphql.NewList(models.AuditoriaType),
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            db := config.GetConnection()
            defer db.Close()

            auditoria := make([]models.Auditoria, 0)

            if err := db.Find(&auditoria).Error; err != nil {
                return nil, err
            }

            return auditoria, nil
        },
    }
}
