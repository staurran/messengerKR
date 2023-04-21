package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/app/utils/token"
)

func (a *Application) GetProfile(gCtx *gin.Context) {
	username := gCtx.Param("username")
	userId, err := a.repo.GetIdByUsername(username)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.repo.GetUserByID(userId)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultProfile := CurrentUserResult{user.Username, user.Avatar, user.Phone, user.Bio}
	gCtx.JSON(http.StatusBadRequest, resultProfile)

	return

}

func (a *Application) ChangeProfile(gCtx *gin.Context) {
	userId, err := token.ExtractTokenID(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input ds.User
	err = gCtx.ShouldBindJSON(&input)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = userId
	err = a.repo.ChangeUser(input)
}

func (a *Application) GetContacts(gCtx *gin.Context) {
	userId, err := token.ExtractTokenID(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contacts, err := a.repo.GetContacts(userId)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gCtx.JSON(http.StatusOK, gin.H{"contacts": contacts})
	return
}

func (a *Application) CreateContact(gCtx *gin.Context) {
	userId, err := token.ExtractTokenID(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input ContactUser
	err = gCtx.ShouldBindJSON(&input)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactId, err := a.repo.GetUserIdByPhone(input.Phone)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var contact ds.Contact
	contact.UserID = userId
	contact.ContactID = contactId
	err = a.repo.CreateContact(contact)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gCtx.JSON(http.StatusOK, nil)
	return
}

func (a *Application) DeleteContact(gCtx *gin.Context) {
	userId, err := token.ExtractTokenID(gCtx)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contactUsername := gCtx.Param("username")

	contactId, err := a.repo.GetIdByUsername(contactUsername)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.repo.DeleteContact(userId, contactId)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gCtx.JSON(http.StatusOK, nil)
	return
}
