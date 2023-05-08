package delivery

import (
	"github.com/gorilla/mux"

	"github.com/staurran/messengerKR.git/internal/pkg/chat"
)

type Handler struct {
	useCase chat.UseCase
}

func NewHandler(useCase chat.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func RegisterHTTPEndpoints(router *mux.Router, uc chat.UseCase) {
	h := NewHandler(uc)
	router.HandleFunc("/iuchat/chat", h.CreateChat).Methods("POST")
	router.HandleFunc("/iuchat/chat/{chat}", h.DeleteChat).Methods("DELETE")
	router.HandleFunc("/iuchat/chat/{chat}", h.ChangeChat).Methods("PUT", "OPTION")
	router.HandleFunc("/iuchat/chats", h.GetChats).Methods("GET")
	router.HandleFunc("/iuchat/message/{chat}", h.GetMessages).Methods("GET")
	router.HandleFunc("/iuchat/message/{chat}", h.SendMessage).Methods("POST")
}
