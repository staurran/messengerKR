package repository

type LastMessage struct {
	Username string
	Content  string
}

type ChatStruct struct {
	Id       uint `json:"id"`
	Name     string
	Avatar   string
	CountMes int64 `json:"count_mes"`
}
