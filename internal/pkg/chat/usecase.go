package chat

type UseCase interface {
	CreateChat(chat ChatInp, userId uint) (chatId uint, err error)
	DeleteChat(chatId uint, userId uint) error
	ChangePhoto(num int, photoId uint, userId uint) error

	GetAllPhotos(userId uint) ([]uint, error)
	GetAvatar(userId uint) (uint, error)
}
