package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/app/repository"
	"github.com/staurran/messengerKR.git/internal/app/utils/token"
)

type ChatStruct struct {
	ID                  uint `json:"id"`
	Name                string
	Avatar              string
	QuantityUnshownMess int64 `json:"quantityUnshownMess"`
	repository.LastMessage
}
type GetChatsResult struct {
	chats []ChatStruct
}

func (a *Application) GetChats(gCtx *gin.Context) {
	userID, err := token.ExtractTokenID(gCtx)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	chats, err := a.repo.GetChats(userID, 1)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant extract userID"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	var result GetChatsResult
	for _, chat := range chats {
		lastMes, err := a.repo.GetLastMes(chat.ID)
		if err != nil {
			answer := AnswerJSON{Status: "error", Description: "cant extract userID"}
			gCtx.IndentedJSON(http.StatusInternalServerError, answer)
			return
		}
		UnshownMes, err := a.repo.GetQuantityShownMes(chat.ID, userID)
		if err != nil {
			answer := AnswerJSON{Status: "error", Description: "cant extract userID"}
			gCtx.IndentedJSON(http.StatusInternalServerError, answer)
			return
		}
		resultChat := ChatStruct{chat.ID, chat.Name, chat.Avatar, UnshownMes, lastMes}

		result.chats = append(result.chats, resultChat)
	}
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get all rows"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, chats)

}

type UserChat struct {
	UserId   uint                  `json:"userId"`
	Username string                `json:"username"`
	Role     constProject.ChatRole `json:"role"`
}
type InpCreateChat struct {
	Users       []UserChat            `json:"users"`
	Name        string                `json:"name"`
	Avatar      string                `json:"avatar"`
	Description string                `json:"description"`
	Type        constProject.TypeChat `json:"type"`
}

func (a *Application) CreateChat(gCtx *gin.Context) {
	var input InpCreateChat
	err := gCtx.ShouldBindJSON(&input)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, u := range input.Users {
		userId, err := a.repo.GetIdByUsername(u.Username)
		if err != nil {
			gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Users[i].UserId = userId
	}

	var chat ds.Chat
	chat.Name = input.Name
	chat.Description = input.Description
	chat.Avatar = input.Avatar
	chat.Type = input.Type
	err = a.repo.CreateChat(&chat)
	if err != nil {
		gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, u := range input.Users {
		var chatUser ds.ChatUser
		chatUser.ChatID = chat.ID
		chatUser.ChatRole = u.Role
		chatUser.UserID = u.UserId
		err = a.repo.CreateChatUser(&chatUser)
		if err != nil {
			gCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

func (a *Application) DeleteChat(gCtx *gin.Context) {
	chatId := gCtx.Param("id_chat")
	err := a.repo.DeleteChat(chatId)
}

//// GetAll godoc
//// @Summary      Show all rows in db
//// @Description  Return all product and info about rows
//// @Tags         Tests
//// @Produce      json
//// @Success      200  {object} []ds.Goods
//// @Router       /goods [get]
//func (a *Application) GetAll(gCtx *gin.Context) {
//	all_rows, err := a.repo.GetAllProducts()
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant get all rows"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	gCtx.IndentedJSON(http.StatusOK, all_rows)
//
//}
//
//// GetProduct godoc
//// @Summary      Show product info by id
//// @Description  Return all info of one product by id
////@Parameters	id
//// @Tags         Tests
//// @Produce      json
//// @Success      200  {object}  ds.Goods
//// @Router       /goods/{id} [get]
//func (a *Application) GetProduct(gCtx *gin.Context) {
//	id_product := gCtx.Param("id")
//	id_product_int, err := strconv.Atoi(id_product)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	product, err := a.repo.GetProductByID(uint(id_product_int))
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	gCtx.IndentedJSON(http.StatusOK, &product)
//}
//
//// ChangePrice godoc
//// @Summary      Change price of product by id
//// @Description  Change price of product by id. Price can't be 0
//// @Tags         Tests
//// @Produce      json
//// @Success      200  {object}  ds.Goods
//// @Router       /goods/{id} [put]
//func (a *Application) ChangeProduct(gCtx *gin.Context) {
//	var params ds.Goods
//	err := gCtx.BindJSON(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant parse json, "}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//		return
//	}
//
//	err = a.repo.ChangeProduct(params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant change price"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	product, err := a.repo.GetProductByID(params.Id_good)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	gCtx.IndentedJSON(http.StatusOK, &product)
//}
//
//// PostProduct godoc
//// @Summary      Add new row
//// @Description  add new row with parameters in json
//// @Tags         Tests
//// @Produce      json
//// @Success      200  {object}  ds.Goods
//// @Router       /goods [post]
//func (a *Application) PostProduct(gCtx *gin.Context) {
//	var params ds.Goods
//	err := gCtx.BindJSON(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//		return
//	}
//	if params.Price <= 0 {
//		answer := AnswerJSON{Status: "error", Description: "price cant be <= 0"}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//	}
//
//	err = a.repo.CreateProduct(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant create product row"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	gCtx.IndentedJSON(http.StatusOK, params)
//}
//
//// DeleteProduct godoc
//// @Summary      Delete row by id
//// @Description  Delete row by id. If there is not this id return error
//// @Tags         Tests
//// @Produce      json
//// @Success      200  {object}  AnswerJSON
//// @Router       /goods/{id} [delete]
//func (a *Application) DeleteProduct(gCtx *gin.Context) {
//	id_product := gCtx.Param("id")
//	id_product_int, err := strconv.Atoi(id_product)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//		return
//	}
//	err = a.repo.DeleteProduct(uint(id_product_int))
//
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant delete row"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	answer := AnswerJSON{Status: "successful", Description: "row was deleted"}
//	gCtx.IndentedJSON(http.StatusOK, answer)
//}
//
//func (a *Application) AddBasketRow(gCtx *gin.Context) {
//	var params ds.Basket
//	err := gCtx.BindJSON(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//		return
//	}
//	id_user, err := token.ExtractTokenID(gCtx)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	params.Id_user = id_user
//	err = a.repo.CreateBasketRow(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "good cant be added to basket"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	answer := AnswerJSON{Status: "successful", Description: "good was added to basket"}
//	gCtx.IndentedJSON(http.StatusOK, answer)
//}
//
//func (a *Application) GetBasket(gCtx *gin.Context) {
//	id_user, err := token.ExtractTokenID(gCtx)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	basket, err := a.repo.GetBasket(id_user)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant get rows in basket"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	gCtx.IndentedJSON(http.StatusOK, basket)
//}
//
//func (a *Application) DeleteBasketRow(gCtx *gin.Context) {
//	id_user, err := token.ExtractTokenID(gCtx)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	var params ds.Basket
//	id_product := gCtx.Param("id")
//	id_product_int, err := strconv.Atoi(id_product)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//		return
//	}
//	params.Id_good = uint(id_product_int)
//	params.Id_user = id_user
//	err = a.repo.DeleteBasketRow(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "good cant be added to basket"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	answer := AnswerJSON{Status: "successful", Description: "good was added to basket"}
//	gCtx.IndentedJSON(http.StatusOK, answer)
//}
//
//func (a *Application) ChangeQuantity(gCtx *gin.Context) {
//	var params ds.Basket
//	err := gCtx.BindJSON(&params)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
//		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
//		return
//	}
//	id_user, err := token.ExtractTokenID(gCtx)
//	if err != nil {
//		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
//		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//		return
//	}
//	params.Id_user = id_user
//	if params.Quantity >= 0 {
//		err = a.repo.DeleteBasketRow(&params)
//		if err != nil {
//			answer := AnswerJSON{Status: "error", Description: "good cant be added to basket"}
//			gCtx.IndentedJSON(http.StatusInternalServerError, answer)
//			return
//		}
//		answer := AnswerJSON{Status: "successful", Description: "good was added to basket"}
//		gCtx.IndentedJSON(http.StatusOK, answer)
//	}
//	err = a.repo.ChangeQuantity(&params, params.Quantity)
//}
