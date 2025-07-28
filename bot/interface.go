package bot

type Bot interface {
	Start() error
	SendMessage(chatID string, message string) error
	Platform() string
}
