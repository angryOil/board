package main

import (
	"board/cmd/app/handler"
	"board/internal/controller"
	"board/internal/repository"
	"board/internal/repository/infla"
	"board/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	h := getHandler()

	r.PathPrefix("/boards/{cafeId:[0-9]+}").Handler(h)

	err := http.ListenAndServe(":8089", r)
	if err != nil {
		panic(err)
	}
}

func getHandler() http.Handler {
	return handler.NewHandler(controller.NewController(service.NewService(repository.NewRepository(infla.NewDB()))))
}
