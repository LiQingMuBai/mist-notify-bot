package domain

type MessageToSend struct {
	ChatId int64
	Text   string
	Images []string
}
