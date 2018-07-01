package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/queries"
)

func RootQuery() graphql.Fields {
	return graphql.Fields{
        "Auditorias": queries.AuditoriaQuery(),
        "Compras": queries.CompraQuery(),
        "Productos": queries.ProdutoQuery(),
        "Terceros": queries.TerceroQuery(),
		"Usuarios": queries.UsuarioQuery(),
        "Ventas": queries.VentaQuery(),
	}
}
