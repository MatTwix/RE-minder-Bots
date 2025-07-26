package bot

type Bot interface {
	Start() error
	SendMessage(chatID int64, message string) error
	Platform() string
}
