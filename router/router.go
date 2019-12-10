package router

import (
	"user/domain"
	"user/repositories"
	"user/services"

	"net/http"

	"github.com/gorilla/mux"
)

func New() (routers *mux.Router) {
	routers = mux.NewRouter()
	db := repositories.New()
	serv := services.New(db)
	user := domain.New(serv)

	routers.HandleFunc("/register", user.Register).Methods(http.MethodPost)

	return
}
