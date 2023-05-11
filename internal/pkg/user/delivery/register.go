package delivery

import (
	"github.com/gorilla/mux"

	"github.com/staurran/messengerKR.git/internal/pkg/user"
)

type Handler struct {
	useCase user.UseCase
}

func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc user.UseCase) {
	h := NewHandler(uc)
	router.HandleFunc("/iuchat/contacts/{phone}", h.CreateContact).Methods("POST")
	router.HandleFunc("/iuchat/contacts/{contact}", h.DeleteContact).Methods("DELETE")
	router.HandleFunc("/iuchat/contacts", h.GetContacts).Methods("GET")
	router.HandleFunc("/iuchat/user/{user_id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/iuchat/user", h.GetCurrentUser).Methods("GET")

}
