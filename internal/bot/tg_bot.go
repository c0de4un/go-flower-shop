package bot

import "sync"

type TelegramBot struct {
	config *TelegramConfig
}

var (
	telegramBotInstance *TelegramBot
	telegramBotSync     sync.Once
)

func InitTelegramBot(config *TelegramConfig) {
	telegramBotSync.Do(func() {
		telegramBotInstance = &TelegramBot{
			config: config,
		}
	})
}

func TerminateTelegramBot() {
}

func GetTelegramBot() *TelegramBot {
	return telegramBotInstance
}
