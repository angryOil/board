package handler

import (
	"board/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	c controller.Controller
}

func NewHandler(c controller.Controller) http.Handler {
	r := mux.NewRouter()
	h := Handler{c: c}
	r.HandleFunc("/board/{cafeId:[0-9]+}/{boardType:[0-9]+}", h.getList)
	return r
}

func (h Handler) getList(writer http.ResponseWriter, request *http.Request) {

}
