package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
	Router mux.Router
}

func New(h mux.Router) *handler {
	return &handler{
		Router: h,
	}
}

func (h *handler) CreateHandlers(){

	h.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello, world!"))
	})
}
