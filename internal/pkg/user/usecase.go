package user

type UseCase interface {
	CreateContact(userId uint, phone string) error
	DeleteContact(userId uint, contact uint) error
	GetContacts(userId uint) ([]Contact, error)
	GetUserById(userId uint) (UserInfo, error)
}
