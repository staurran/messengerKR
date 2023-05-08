package delivery

import (
	"github.com/gorilla/mux"

	"github.com/staurran/messengerKR.git/internal/pkg/photo"
)

type Handler struct {
	useCase photo.UseCase
}

func NewHandler(useCase photo.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc photo.UseCase) {
	h := NewHandler(uc)
	router.HandleFunc("/iuchat/photos/upload", h.AddPhoto).Methods("POST")
	router.HandleFunc("/iuchat/photo/{photo}", h.GetPhoto).Methods("GET")
	router.HandleFunc("/iuchat/photo/{photo}", h.DeletePhoto).Methods("DELETE")
	router.HandleFunc("/iuchat/photo/{photo}", h.ChangePhoto).Methods("PUT", "OPTIONS")
}
