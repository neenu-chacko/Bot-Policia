package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotAPI("906142594:AAGp2PgkOUWFNmcFq0fazTU8APVmGSgpaPk")
	if err != nil {
		panic(err)
	}

	fmt.Print("Bot Connected Successfully!\n")
	log.Printf("Authorized on account %s", bot.Self.UserName)

	isWelcome := [10]string{"Keri vaada makkale 🤗🥰😎",
		"Ente moneii ni ntha ithra vaikye",
		"Ithevdarnn",
		"Ini nammal oru poli polikkm",
		"Pani Varunnund Avaracha",
		"Hello there! We're extremely happy to have you on board!",
		"Upadravikkaruth",
		"Sneham Matharam",
		"Yes!You're in",
		"Biju Pls!!!"}
	isExit := [10]string{
		"Ninak vendel enikkum vendedo uvve",
		"Good Bye,Don't cry. We won't!",
		" We'll miss you! I'm Joking...You're dead to us!",
		"See you somewhere else ",
		"You'll be greatly missed",
		" Crap!! You're leaving!!!!",
		" Why bruh why?!",
		"Fine! Go!",
		"Bon Voyage",
		"Wokeiii.......!!",
	}
	n := rand.Intn(10)

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

				msg := tgbot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s", isWelcome[n], joinedUsers))
				bot.Send(msg)

				for _, user := range *update.Message.NewChatMembers {
					restrict(bot, user, update.Message.Chat.UserName)
				}

				continue
			}

			if update.Message.LeftChatMember != nil {
				user := update.Message.LeftChatMember
				msg := tgbot.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s\n%s😏😤", isExit[n], user.FirstName))
				bot.Send(msg)

				continue
			}
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
				UserID:             user.ID,
				SuperGroupUsername: chatID,
			},
			CanAddWebPagePreviews: true,
		},
	)
}
