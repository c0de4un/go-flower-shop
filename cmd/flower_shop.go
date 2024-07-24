package main

import (
	"fmt"

	"github.com/c0de4un/go-flower-shop/internal/bot"
	"github.com/c0de4un/go-flower-shop/internal/logging"
)

func main() {
	logging.InitializeLogger()

	botCfg, err := bot.LoadTGConfig("configs/tg.xml")
	if err != nil {
		panic(fmt.Sprintf("failed to load tg config: %v\n", err))
	}

	bot.InitTelegramBot(botCfg)
	bot.TerminateTelegramBot()

	fmt.Println("Hello World !")
}
