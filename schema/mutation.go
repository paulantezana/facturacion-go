package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/mutations"
)

func RootMutation() graphql.Fields {
	return graphql.Fields{
		"CreateUsuario": mutations.CreateUsuarioMutation(),
	}
}
