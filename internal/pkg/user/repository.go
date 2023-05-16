package user

import "github.com/staurran/messengerKR.git/internal/app/ds"

type IUserRepository interface {
	GetUserIdByPhone(phone string) (uint, error)
	CreateContact(contact ds.Contact) error
	DeleteContact(userId, contactId uint) error
	GetContacts(userId uint) ([]Contact, error)
	GetUserById(userId uint) (UserInfo, error)
}
