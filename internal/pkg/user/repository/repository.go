package repository

import (
	"gorm.io/gorm"

	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/pkg/user"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {

	r := UserRepository{db}
	return &r
}

func (r *UserRepository) GetUserById(id uint) (ds.User, error) {
	user := ds.User{}
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return ds.User{}, err
	}
	return user, nil
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
	if userInp.Bio != "" {
		userDB.Bio = userInp.Bio
	}
	if userInp.Phone != "" {
		userDB.Password = userInp.Password
	}
	err = r.db.Save(&userDB).Error
	return err
}

func (r *UserRepository) GetContacts(userId uint) (contacts []user.Contact, err error) {
	err = r.db.Table("contacts c").
		Select("c.contact_id as user_id, u.username, u.phone, u.avatar").
		Where("c.user_id=?", userId).
		Joins("Join users u on u.id=c.contact_id").Find(&contacts).Error
	return
}

func (r *UserRepository) CreateContact(contact ds.Contact) error {
	err := r.db.Create(&contact).Error
	return err
}

func (r *UserRepository) DeleteContact(userId, contactId uint) error {
	contact := &ds.Contact{}
	err := r.db.First(contact, "user_id = ? AND contact_id = ?", userId, contactId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(contact).Error
	return err
}

func (r *UserRepository) GetUserIdByPhone(phone string) (uint, error) {
	user := &ds.User{}
	err := r.db.First(user, "phone = ?", phone).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
