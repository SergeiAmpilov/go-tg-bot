package main

import (
	"fmt"
	"os"
	env "tgtest/config"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var cfg env.Config

func init() {
	cfg = *env.New()
}

func main() {

	// Get Bot token from environment variables
	botToken := cfg.Token
	fmt.Println(botToken)
	//

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatId := tu.ID(update.Message.Chat.ID)

			_, _ = bot.CopyMessage(
				tu.CopyMessage(
					chatId,
					chatId,
					update.Message.MessageID,
				),
			)
		}
	}

	_, _ = bot, updates
}
