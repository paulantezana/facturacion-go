package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/facturacion-go/models"
    "github.com/paulantezana/facturacion-go/config"
    "crypto/sha256"
    "fmt"
    "crypto/md5"
)

func CreateUsuarioMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UsuarioType,
		Args: graphql.FieldConfigArgument{
			"usuario": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"nombre":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"clave":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			usuario := models.Usuario{
				Usuario: p.Args["usuario"].(string),
				Email:   p.Args["email"].(string),
				Nombre:  p.Args["nombre"].(string),
				Clave:   p.Args["clave"].(string),
			}

			// Get Avatar from gravatar
            picmd5 := md5.Sum([]byte(usuario.Email))
            picstr := fmt.Sprintf("%x", picmd5)
            usuario.Avatar = "https://gravatar.com/avatar/" + picstr + "?s=100"

			// get connection
			db := config.GetConnection()
			defer db.Close()

			cc := sha256.Sum256([]byte(usuario.Clave))
			pwd := fmt.Sprintf("%x", cc)
			usuario.Clave = pwd

			if err := db.Create(&usuario).Error; err != nil {
			   return nil, err
			}

			return usuario, nil
		},
	}
}
