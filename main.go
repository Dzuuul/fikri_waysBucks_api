package main

import (
	"fmt"
	"net/http"
	"ways-bucks-api/database"
	"ways-bucks-api/pkg/mysql"
	"ways-bucks-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
