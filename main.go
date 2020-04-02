package main

import (
	"fmt"
	"strings"
    "log"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("906142594:AAGp2PgkOUWFNmcFq0fazTU8APVmGSgpaPk")
	if err != nil {
		panic(err)
	}
	fmt.Print("Bot Connected Successfully!\n")
    	log.Printf("Authorized on account %s", bot.Self.UserName)

	responses := map[string]func() string{
		"/start":    func() string { return "Nice to meet you!" },
		"hi":        func() string { return "Hi!" },
		"poli":      func() string { return "sanam" },
		"go corona": func() string { return "Corona Go" },
		"sugalle" :func() string { return "Parama Sugam!" },
		"stayhome":  func() string { return "#veettilirimyre" },
	}

	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		str := update.Message.Text
		resp, ok := responses[strings.ToLower(str)]
		if !ok {
			msg := tgbot.NewMessage(update.Message.Chat.ID, "Jaba jaba, Jabaabi jaba!")
			bot.Send(msg)
			continue
		}
		msg := tgbot.NewMessage(update.Message.Chat.ID, resp())
		bot.Send(msg)
	}
}
