package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/ds"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) CreateUser(user *ds.User) error {
	err := r.db.Create(user).Error
	return err
}

func (r *UserRepository) Login(user *ds.User) error {
	user_db := ds.User{}
	err := r.db.Model(&ds.User{}).Where("username = ?", user.Username).Take(&user_db).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	user.Id = user_db.Id
	return nil
}

func (r *UserRepository) GetUserByID(id uint) (*ds.User, error) {
	user := &ds.User{}
	err := r.db.First(user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetIdByUsername(username string) (uint, error) {
	user := &ds.User{}
	err := r.db.First(user, "username = ?", username).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *UserRepository) ChangeUser(userInp ds.User) error {
	userDB := &ds.User{}
	err := r.db.First(userDB, "id = ?", userInp.Id).Error // find product with code D42
	if err != nil {
		return err
	}
	if userInp.Username != "" {
		userDB.Username = userInp.Username
	}
	if userInp.Avatar != "" {
		userDB.Avatar = userInp.Avatar
	}
	if userInp.Bio != "" {
		userDB.Bio = userInp.Bio
	}
	if userInp.Phone != "" {
		userDB.Password = userInp.Password
	}
	err = r.db.Save(&userDB).Error
	return err
}

func (r *UserRepository) GetContacts(userId uint) (contacts []ds.User, err error) {
	err = r.db.Find(&contacts, "user_id = ?", userId).Error
	return
}

func (r *UserRepository) GetUserIdByPhone(phone string) (userId uint, err error) {
	user := &ds.User{}
	err = r.db.First(user, "phone = ?", phone).Error

	return user.Id, err
}

func (r *UserRepository) CreateContact(contact ds.Contact) error {
	err := r.db.Create(contact).Error
	return err
}

func (r *UserRepository) DeleteContact(userId, contactId uint) error {
	contact := &ds.Contact{}
	err := r.db.First(contact, "user_id = ? contact_id = ?", userId, contactId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(contact).Error
	return err
}
