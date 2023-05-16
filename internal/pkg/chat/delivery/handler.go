package delivery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/staurran/messengerKR.git/internal/pkg/chat"
	"github.com/staurran/messengerKR.git/utils/logger"
	"github.com/staurran/messengerKR.git/utils/writer"
)

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var userJson chat.ChatInp
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		err = fmt.Errorf("cant get userId")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	chatId, err := h.useCase.CreateChat(userJson, uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{"chatId": chatId})
}

func (h *Handler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	chatIdStr, ok := params["chat"]
	if !ok {
		err := fmt.Errorf("no chat")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		err = fmt.Errorf("cant get userId")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	err = h.useCase.DeleteChat(uint(chatId), uint(userId))

	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("error")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) ChangeChat(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	chatIdStr, ok := params["chat"]
	if !ok {
		err := fmt.Errorf("no chat")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		err = fmt.Errorf("cant get userId")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	defer func() {
		err = r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var userJson chat.ChatInp
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	err = h.useCase.ChangeChat(uint(chatId), uint(userId), userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}

func (h *Handler) GetChats(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		err := fmt.Errorf("cant get userId")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	chats, err := h.useCase.GetChats(uint(userId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{"chats": chats})
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		err := fmt.Errorf("cant get userId")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	chatIdStr, ok := params["chat"]
	if !ok {
		err := fmt.Errorf("no chat")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	msgs, err := h.useCase.GetMessages(uint(userId), uint(chatId))
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{"messages": msgs})
}

func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	userIdDB := r.Context().Value("userId")
	userId, ok := userIdDB.(uint32)
	if !ok {
		err := fmt.Errorf("cant get userId")
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	defer func() {
		err := r.Body.Close()
		if err != nil {
			logger.Log(http.StatusInternalServerError, err.Error(), r.Method, r.URL.Path, true)
			writer.ErrorRespond(w, r, err, http.StatusInternalServerError)
			return
		}
	}()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	var userJson chat.MessageInp
	err = json.Unmarshal(reqBody, &userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	err = h.useCase.CreateMessage(uint(userId), userJson)
	if err != nil {
		logger.Log(http.StatusBadRequest, err.Error(), r.Method, r.URL.Path, true)
		err = fmt.Errorf("cant parse json")
		writer.ErrorRespond(w, r, err, http.StatusBadRequest)
		return
	}

	logger.Log(http.StatusOK, "Success", r.Method, r.URL.Path, false)
	writer.Respond(w, r, map[string]interface{}{})
}
