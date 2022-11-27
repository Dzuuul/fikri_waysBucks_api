package routes

import (
	"ways-bucks-api/handlers"
	"ways-bucks-api/pkg/middleware"
	"ways-bucks-api/pkg/mysql"
	"ways-bucks-api/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", middleware.Auth(h.FindToppings)).Methods("GET")
	r.HandleFunc("/topping/{id}", h.GetTopping).Methods("GET")
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", middleware.UploadFile(h.UpdateTopping)).Methods("PATCH")
	r.HandleFunc("/topping/{id}", h.DeleteTopping).Methods("DELETE")
}
