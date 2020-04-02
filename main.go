package main

import (
	"fmt"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("906142594:AAGp2PgkOUWFNmcFq0fazTU8APVmGSgpaPk")
	if err != nil {
		panic(err)
	}
	fmt.Print("Bot Connected Successfully!\n")

	responses := map[string]func() string{
		"/start":    func() string { return "Nice to meet you!" },
		"hi":        func() string { return "Hi!" },
		"poli":      func() string { return "sanam" },
		"Go Corona": func() string { return "Corona Go" },
		"stayhome":  func() string { return "#veettilirimyre" },
	}

	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		resp, ok := responses[update.Message.Text]
		if !ok {
			msg := tgbot.NewMessage(update.Message.Chat.ID, "Jaba jaba, Jabaabi jaba!")
			bot.Send(msg)
			continue
		}
		msg := tgbot.NewMessage(update.Message.Chat.ID, resp())
		bot.Send(msg)
	}
}
