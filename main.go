
package main
import (
		tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
		"fmt"
)

func main() {
		bot, err := tgbot.NewBotAPI("906142594:AAGp2PgkOUWFNmcFq0fazTU8APVmGSgpaPk")
		if err != nil {
			panic(err)
			return
		}
	  fmt.Print("Bot Connected Successfully!\n")
	  u := tgbot.NewUpdate(0)
	  updates,err := bot.GetUpdatesChan(u)

	  for update := range updates{
		  if update.Message == nil{
			  continue
		  }
		  msg := tgbot.NewMessage(update.Message.Chat.ID,"Aaha itharith!! Keri Vaa")
		  bot.Send(msg)
	  }
}