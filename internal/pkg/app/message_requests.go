package app

import (
	"github.com/gin-gonic/gin"
	"lab3/internal/app/ds"
	"lab3/internal/app/repository"
	"lab3/internal/app/utils/token"
	"log"
	"net/http"
	"strconv"
	"time"
)

type GoodQuantity struct {
	Id_good  uint `json:"id_good"`
	Quantity int  `json:"quantity"`
}

type ReqStruct struct {
	Baskets []GoodQuantity `json:"baskets"`
	Total   int            `json:"total"`
}

func (a *Application) AddOrder(gCtx *gin.Context) {
	var params ReqStruct

	err := gCtx.BindJSON(&params)
	log.Println(params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	date := time.Now().Format("01-02-2006")
	user_id, err := token.ExtractTokenID(gCtx)

	order := ds.Orders{Status: 1, Date: date, Id_user: user_id, Total: params.Total}
	err = a.repo.CreateOrder(&order)

	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create order"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}

	for _, good := range params.Baskets {
		var orderGood ds.GoodOrder

		orderGood.Id_good = good.Id_good
		orderGood.Id_order = order.Id_order
		orderGood.Quantity = good.Quantity
		err = a.repo.CreateGoodOrder(&orderGood)
		if err != nil {

			answer := AnswerJSON{Status: "error", Description: "cant add good in order"}
			gCtx.IndentedJSON(http.StatusInternalServerError, answer)
			return
		}
	}

	answer := AnswerJSON{Status: "successful", Description: "good was added to basket"}
	gCtx.IndentedJSON(http.StatusOK, answer)
}

type OrderRes struct {
	Id_order    uint                      `json:"id_order"`
	Date        string                    `json:"date"`
	Status      string                    `json:"status"`
	Description string                    `json:"description"`
	Total       int                       `json:"total"`
	Goods       []repository.GoodQuantity `json:"goods"`
	Login       string                    `json:"login"`
}

func (a *Application) GetAllOrders(gCtx *gin.Context) {
	id_user, err := token.ExtractTokenID(gCtx)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	order, err := a.repo.GetOrder(id_user)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get rows in basket"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	var results []OrderRes
	for _, ord := range order {
		var row OrderRes
		row.Date = ord.Date
		row.Id_order = ord.Id_order
		row.Status = ord.Name
		row.Description = ord.Description
		row.Total = ord.Total
		row.Goods, err = a.repo.GetGoodOrder(ord.Id_order)
		results = append(results, row)
	}
	gCtx.IndentedJSON(http.StatusOK, results)
}

func (a *Application) GetOrders(gCtx *gin.Context) {
	order, err := a.repo.GetAllOrders()
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get all rows"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	var results []OrderRes
	for _, ord := range order {
		var row OrderRes
		row.Date = ord.Date
		row.Id_order = ord.Id_order
		row.Status = ord.Name
		row.Description = ord.Description
		row.Total = ord.Total
		row.Login = ord.Login
		row.Goods, err = a.repo.GetGoodOrder(ord.Id_order)
		results = append(results, row)
	}
	gCtx.IndentedJSON(http.StatusOK, results)

}

func (a *Application) DeleteOrder(gCtx *gin.Context) {
	id_order := gCtx.Param("id")
	id_order_int, err := strconv.Atoi(id_order)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	err = a.repo.ChangeStatus(uint(id_order_int), 4)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant change status"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	answer := AnswerJSON{Status: "success", Description: "changed"}
	gCtx.IndentedJSON(http.StatusOK, answer)
}

func (a *Application) GetStatus(gCtx *gin.Context) {
	all_rows, err := a.repo.GetAllStatuses()
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get all rows"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, all_rows)
}

func (a *Application) ChangeStatus(gCtx *gin.Context) {
	id_status := gCtx.Param("id_status")
	id_status_int, err := strconv.Atoi(id_status)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	id_order := gCtx.Param("id_order")
	id_order_int, err := strconv.Atoi(id_order)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	err = a.repo.ChangeStatus(uint(id_order_int), uint(id_status_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant change status"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	answer := AnswerJSON{Status: "success", Description: "changed"}
	gCtx.IndentedJSON(http.StatusOK, answer)
}
