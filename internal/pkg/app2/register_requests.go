package app2

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"log"
	"net/http"

	"github.com/staurran/messengerKR.git/internal/app/constProject"
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/app/utils/token"
)

//функция регистрации
func (a *Application) Register(gCtx *gin.Context) {
	var input RegisterInput
	err := gCtx.ShouldBindJSON(&input)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := CreatePass(input.Password)
	log.Println(hashedPassword)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := ds.User{}
	u.Username = input.Username
	u.Avatar = input.Avatar
	u.Phone = input.Phone
	u.Password = hashedPassword
	u.Role = constProject.User

	err = a.repo.CheckUsername(u.Username)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "login was used before"})
		return
	}
	err = a.repo.CheckPhone(u.Phone)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "login was used before"})
		return
	}

	err = a.repo.CreateUser(&u)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token_user, err := token.GenerateToken(u.ID, u.Role)
	gCtx.JSON(http.StatusOK, gin.H{"token": token_user, "constProject": u.Role})
}

//функция входа/авторизации
func (a *Application) Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := ds.User{}
	u.Username = input.Username
	u.Password = input.Password

	if err := a.repo.Login(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token_user, err := token.GenerateToken(u.ID, u.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem with token."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token_user, "constProject": u.Role})
}

//создание зашифрованного пароля
func CreatePass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

//возвращает информацию о пользователе. Только для админов
func (a *Application) CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.repo.GetUserByID(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultUser := CurrentUserResult{user.Username, user.Avatar, user.Phone, user.Bio}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resultUser})
}
