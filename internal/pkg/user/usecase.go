package user

import "github.com/staurran/messengerKR.git/internal/app/ds"

type UseCase interface {
	CreateContact(userId uint, phone string) error
	DeleteContact(userId uint, contact uint) error
	GetContacts(userId uint) ([]Contact, error)
	GetUserById(userId uint) (ds.User, error)
}
