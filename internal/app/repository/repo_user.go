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

func (r *Repository) ChangeUser(userInp ds.User) error {
	userDB := &ds.User{}
	err := r.db.First(userDB, "id = ?", userInp.ID).Error // find product with code D42
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

func (r *Repository) GetContacts(userId uint) (contacts []ds.User, err error) {
	err = r.db.Find(&contacts, "user_id = ?", userId).Error
	return
}

func (r *Repository) GetUserIdByPhone(phone string) (userId uint, err error) {
	user := &ds.User{}
	err = r.db.First(user, "phone = ?", phone).Error

	return user.ID, err
}

func (r *Repository) CreateContact(contact ds.Contact) error {
	err := r.db.Create(contact).Error
	return err
}

func (r *Repository) DeleteContact(userId, contactId uint) error {
	contact := &ds.Contact{}
	err := r.db.First(contact, "user_id = ? contact_id = ?", userId, contactId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(contact).Error
	return err
}
