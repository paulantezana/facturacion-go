package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mnmtanish/go-graphiql"
	"github.com/paulantezana/facturacion-go/config"
	"github.com/paulantezana/facturacion-go/models"
	"github.com/paulantezana/facturacion-go/schema"
	"log"
	"net/http"
	"os"
)

func main() {
	// Migration database
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	//if migrate == "yes" {
	fmt.Println("Init migration")
	Migrate()
	fmt.Println("Finalize migration")
	//}

	// Create new server mux
	router := mux.NewRouter()

	// Custom port
	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetConfig().Server.Port
	}

	//router.HandleFunc("/login",security.Login).Methods("POST")
	router.Handle("/graphql", schema.GraphQL())            // GraphQL Server
	router.HandleFunc("/graphiql", graphiql.ServeGraphiQL) // GraphiQL Server

	fmt.Println("=> http server started on http://localhost:" + port) // listening server message

	// Config coors server listener
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedOrigins([]string{"*"}))(router),
	))
}

func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	db.AutoMigrate(
		&models.Auditoria{},
		&models.Compra{},
		&models.Producto{},
		&models.Tercero{},
		&models.Usuario{},
		&models.Venta{},
	)
}
