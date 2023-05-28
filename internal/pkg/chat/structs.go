package chat

import (
	"time"
)

type ChatInp struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Users       []uint `json:"users"`
}

type ChatRepoStruct struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	CountMes int64  `json:"count_mes"`
}

type ChatStruct struct {
	Id                  uint   `json:"id"`
	Name                string `json:"name"`
	Avatar              string `json:"avatar"`
	QuantityUnshownMess int64  `json:"quantityUnshownMess"`
	LastMessage
}

type LastMessage struct {
	Username string    `json:"username"`
	Content  string    `json:"content"`
	Time     time.Time `json:"time"`
}

type MessageTemp struct {
	ID          uint          `json:"id"`
	Content     string        `json:"content"`
	UserId      uint          `json:"userFrom"`
	Attachment  []string      `json:"attachment"`
	Photos      []string      `json:"photo"`
	Audio       string        `json:"audio"`
	TimeCreated time.Time     `json:"timeCreated"`
	Reactions   []ReactionMes `json:"reactions"`
	Shown       bool          `json:"shown"`
}

type Message struct {
	ID          uint          `json:"id"`
	Content     string        `json:"content"`
	UserFrom    UserFrom      `json:"userFrom"`
	Attachment  []string      `json:"attachment"`
	Photos      []string      `json:"photo"`
	Audio       string        `json:"audio"`
	TimeCreated time.Time     `json:"timeCreated"`
	Reactions   []ReactionMes `json:"reactions"`
	Shown       bool          `json:"shown"`
}

type UserFrom struct {
	UserId   uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type ReactionMes struct {
	Reaction string `json:"reaction"`
	Quantity int    `json:"quantity"`
}

type MessageInp struct {
	Content    string   `json:"content"`
	ChatId     uint     `json:"chatId"`
	Attachment []string `json:"attachment"`
	Photos     []string `json:"photo"`
	Audio      string   `json:"audio"`
}

type ReactionList struct {
	Reaction string
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type ChatInfo struct {
	Id          uint       `json:"id"`
	Name        string     `json:"name"`
	Users       []UserFrom `json:"users"`
	Description string     `json:"description"`
	Avatar      string     `json:"avatar"`
}
