package bot

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/c0de4un/go-flower-shop/internal/logging"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type TelegramBot struct {
	config    *TelegramConfig
	isEnabled atomic.Bool
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

		telegramBotInstance.start()
	})
}

func TerminateTelegramBot() {
	GetTelegramBot().isEnabled.Store(false)
}

func GetTelegramBot() *TelegramBot {
	return telegramBotInstance
}

func (b *TelegramBot) start() {
	bot, err := telego.NewBot(b.config.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		logging.GetLogger().Error(
			fmt.Sprintf("TelegramBot::start:NewBot: %v", err),
		)

		panic(err)
	}

	botUser, err := bot.GetMe()
	if err != nil {
		logging.GetLogger().Error(
			fmt.Sprintf("TelegramBot::start:GetMe: %v", err),
		)

		panic(err)
	}

	logging.GetLogger().Debug(
		fmt.Sprintf("Bot user: %v\n", botUser),
	)

	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer bh.Stop()

	// Stop getting updates
	defer bot.StopLongPolling()

	// Register new handler with match on command `/start`
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("Catalog").WithWebApp(
					&telego.WebAppInfo{
						URL: b.config.AppUrl,
					},
				),
			),
		)

		msg := tu.Message(message.Chat.ChatID(), "Welcome !").
			WithReplyMarkup(keyboard).
			WithProtectContent()

		bot.SendMessage(msg)
	}, th.CommandEqual("start"))

	// Start handling updates
	bh.Start()
}
