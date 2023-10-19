package routes

import (
	"net/http"

	"github.com/gupta29470/golang-sql-crud-with-orm/controllers"
)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users/create", controllers.CreateUser)
	mux.HandleFunc("/users", controllers.GetUsers)
	mux.HandleFunc("/user/", controllers.GetUser)
	mux.HandleFunc("/user/update", controllers.UpdateUser)
	mux.HandleFunc("/user/delete/", controllers.DeleteUser)
}
