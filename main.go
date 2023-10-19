package main

import (
	"fmt"
	"net/http"

	"github.com/gupta29470/golang-sql-crud-with-orm/databases"
	"github.com/gupta29470/golang-sql-crud-with-orm/routes"
)

func init() {
	databases.InitDB()
}

func main() {
	mux := http.NewServeMux()

	routes.UserRoutes(mux)
	fmt.Println("Running on port 9000")
	http.ListenAndServe(":9000", mux)
}
