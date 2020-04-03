package main

import (
	"fmt"
	"log"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	welcome                = `Keri vaada makkale 🤗🥰😎`
	releasesCommandWelcome = "releases"
)

const (
	exit                = `Ninak vendel enikkum vendedo uvve`
	releasesCommandExit = "releases"
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
		"sugalle":   func() string { return "Parama Sugam!" },
		"stayhome":  func() string { return "#veettilirimyre" },
	}

	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
			if update.Message.NewChatMembers != nil {
				var newUsers []string
				for _, user := range *update.Message.NewChatMembers {
					newUsers = append(newUsers, "@"+getUserName(user))
				}
				joinedUsers := strings.Join(newUsers, " ")
				msg := tgbot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s", welcome, joinedUsers))
				bot.Send(msg)

				for _, user := range *update.Message.NewChatMembers {
					restrict(bot, user, update.Message.Chat.UserName)

				}

				continue
			}

			if update.Message.LeftChatMember != nil {
				user := update.Message.LeftChatMember
				msg := tgbot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s😏😤", exit, user.FirstName))
				bot.Send(msg)

				continue
			}
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

func getUserName(user tgbot.User) string {
	if user.UserName == "" {
		return user.FirstName
	}
	return user.UserName
}

func restrict(bot *tgbot.BotAPI, user tgbot.User, chatID string) {
	bot.RestrictChatMember(
		tgbot.RestrictChatMemberConfig{
			ChatMemberConfig: tgbot.ChatMemberConfig{
				UserID:          user.ID,
				ChannelUsername: chatID,
			},
		},
	)
}
