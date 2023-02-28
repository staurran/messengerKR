package constProject

type Role int

const (
	User    Role = iota // 0
	Manager             // 1
	Admin               // 2
)

type TypeChat int

const (
	Chat   TypeChat = iota //0
	Group                  // 1
	Chanel                 // 2
)

type ChatRole int

const (
	ChatAdmin ChatRole = iota //0
	ChatUser                  //1
)
