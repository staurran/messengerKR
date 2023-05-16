package usecase

import (
	"github.com/staurran/messengerKR.git/internal/app/ds"
	"github.com/staurran/messengerKR.git/internal/pkg/user"
)

type UserUseCase struct {
	UserRepo user.IUserRepository
}

func NewUserUseCase(
	userRepo user.IUserRepository,
) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateContact(userId uint, phone string) error {
	contactId, err := uc.UserRepo.GetUserIdByPhone(phone)
	if err != nil {
		return err
	}
	contactRow := ds.Contact{UserID: userId, ContactID: contactId}
	err = uc.UserRepo.CreateContact(contactRow)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) DeleteContact(userId uint, contactId uint) error {
	err := uc.UserRepo.DeleteContact(userId, contactId)
	return err
}

func (uc *UserUseCase) GetContacts(userId uint) ([]user.Contact, error) {
	contacts, err := uc.UserRepo.GetContacts(userId)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (uc *UserUseCase) GetUserById(userId uint) (ds.User, error) {
	userInfo, err := uc.UserRepo.GetUserById(userId)
	if err != nil {
		return ds.User{}, err
	}
	return userInfo, nil
}
