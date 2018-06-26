package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/queries"
)

func RootQuery() graphql.Fields {
	return graphql.Fields{
		"Usuarios": queries.UsuarioQuery(),
	}
}
