package main

import (
	"os"
	"log"
	re "regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("NEIN_DOCH_OH_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			var reply string
			nein, _ := re.MatchString("^[\n ]*[Nn][Ee][Ii][Nn]", update.Message.Text)
			doch, _ := re.MatchString("^[\n ]*[Dd][Oo][Cc][Hh]", update.Message.Text)
			oh, _ := re.MatchString("^[\n ]*[Oo][Hh] *[!]?[\n ]*$", update.Message.Text)
			if nein {
				reply = "Doch!"
			} else if doch {
				reply = "Oh!"
			} else if oh {
				reply = "👍"
			}
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
