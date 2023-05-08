package app2

////Users

type CurrentUserResult struct {
	Username string
	Avatar   string
	Phone    string
	Bio      string
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone" binding:"required"`
	Bio      string `json:"bio"`
}

type ContactUser struct {
	Phone string
}

///Chats
