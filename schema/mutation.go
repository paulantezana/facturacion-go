package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/mutations"
)

func RootMutation() graphql.Fields {
	return graphql.Fields{
		"CreateUsuario": mutations.CreateUsuarioMutation(),
		"UpdateUsuario": mutations.UpdateUsuarioMutation(),
		"DeleteUsuario": mutations.DeleteUsuarioMutation(),

		"CreateProducto": mutations.CreateProductoMutation(),
		"UpdateProducto": mutations.UpdateProductoMutation(),
		"DeleteProducto": mutations.DeleteProductoMutation(),

		"CreateTercero": mutations.CreateTerceroMutation(),
		"UpdateTercero": mutations.UpdateTerceroMutation(),
		"DeleteTercero": mutations.DeleteTerceroMutation(),
	}
}
