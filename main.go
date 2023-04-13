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
		var request *tgbotapi.Message
		// Whether we got a message
		if update.Message != nil {
			request = update.Message
		} else if update.ChannelPost != nil {
			request = update.ChannelPost
		} else { continue }
		var reply string
		nein, _ := re.MatchString("^[\n ]*[Nn][Ee][Ii][Nn]", request.Text)
		doch, _ := re.MatchString("^[\n ]*[Dd][Oo][Cc][Hh]", request.Text)
		oh, _ := re.MatchString("^[\n ]*[Oo][Hh] *[!]?[\n ]*$", request.Text)
		if nein {
			reply = "Doch!"
		} else if doch {
			reply = "Oh!"
		} else if oh {
			reply = "👍"
		}
		msg := tgbotapi.NewMessage(request.Chat.ID, reply)
		msg.ReplyToMessageID = request.MessageID
		bot.Send(msg)
	}
}
