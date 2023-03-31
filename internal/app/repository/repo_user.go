package repository

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/staurran/messengerKR.git/internal/app/ds"
)

func (r *Repository) CreateUser(user *ds.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *Repository) Login(user *ds.User) error {
	user_db := ds.User{}
	err := r.db.Model(&ds.User{}).Where("username = ?", user.Username).Take(&user_db).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	user.ID = user_db.ID
	user.Role = user_db.Role
	return nil
}

func (r *Repository) GetUserByID(id uint) (*ds.User, error) {
	user := &ds.User{}
	err := r.db.First(user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetIdByUsername(username string) (uint, error) {
	user := &ds.User{}
	err := r.db.First(user, "username = ?", username).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
