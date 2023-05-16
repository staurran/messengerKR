package repository

import (
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	dataStruct "github.com/staurran/messengerKR.git/service/auth/pkg/data_struct"
)

type AuthRepository struct {
	db     *gorm.DB
	client *redis.Client
}

func NewRepo(db *gorm.DB, client *redis.Client) *AuthRepository {

	r := AuthRepository{db, client}
	return &r
}

func (r *AuthRepository) Login(email string, passwordInp string) (uint, error) {
	var userDB *dataStruct.User
	err := r.db.Model(&dataStruct.User{}).Where("phone = ?", email).Take(&userDB).Error
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(passwordInp))
	if err != nil {
		return 0, err
	}

	return userDB.Id, nil
}

func (r *AuthRepository) AddUser(user *dataStruct.User) (uint, error) {
	err := r.db.Create(user).Error
	return user.Id, err
}

func (r *AuthRepository) ChangeUser(user dataStruct.User) error {
	userDb := &dataStruct.User{}
	err := r.db.First(userDb, "id= ?", user.Id).Error
	if err != nil {
		return err
	}
	if user.Phone != "" {
		userDb.Phone = user.Phone
	}
	if user.Password != "" {
		userDb.Password = user.Password
	}

	if user.Bio != "" {
		userDb.Bio = user.Bio
	}

	err = r.db.Save(&userDb).Error
	return err
}

func (r *AuthRepository) GetUserIdByToken(token string) (uint, error) {
	user, err := r.client.Get(token).Result()
	if err != nil {
		return 0, err
	}
	userId, err := strconv.Atoi(user)
	if err != nil {
		return 0, err
	}
	return uint(userId), nil
}

func (r *AuthRepository) DeleteToken(token string) error {
	err := r.client.Del(token).Err()
	return err
}

func (r *AuthRepository) SaveToken(userId uint, token string) (err error) {
	userIdStr := strconv.Itoa(int(userId))
	err = r.client.Set(token, userIdStr, 120*time.Hour).Err()
	return
}

func (r *AuthRepository) CheckSession(token string) (uint, error) {
	userIdStr, err := r.client.Get(token).Result()
	if err != nil {
		return 0, err
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return 0, err
	}

	return uint(userId), nil
}
