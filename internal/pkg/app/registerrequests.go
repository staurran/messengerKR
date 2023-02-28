package app

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lab3/internal/app/ds"
	"lab3/internal/app/role"
	"lab3/internal/app/utils/token"
	"log"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//функция регистрации
func (a *Application) Register(gCtx *gin.Context) {
	var input RegisterInput
	err := gCtx.ShouldBindJSON(&input)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := ds.Users{}
	u.Login = input.Username
	hashedPassword, err := CreatePass(input.Password)
	log.Println(hashedPassword)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.Password = hashedPassword
	u.Role = role.User
	err = a.repo.CheckLogin(u.Login)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": "login was used before"})
		return
	}
	err = a.repo.CreateUser(&u)
	if err != nil {
		gCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token_user, err := token.GenerateToken(u.Id_user, u.Role)
	gCtx.JSON(http.StatusOK, gin.H{"token": token_user, "role": u.Role})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//функция входа/авторизации
func (a *Application) Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := ds.Users{}

	u.Login = input.Username
	hashedPassword, err := CreatePass(input.Password)
	log.Println(hashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant convert password."})
		return
	}
	u.Password = input.Password
	err = a.repo.LoginCheck(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println(u.Id_user)
	token_user, err := token.GenerateToken(u.Id_user, u.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "problem with token."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token_user, "role": u.Role})

}

//создание зашифрованного пароля
func CreatePass(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

//возвращает информацию о пользователе. Только для админов
func (a *Application) CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := a.repo.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
