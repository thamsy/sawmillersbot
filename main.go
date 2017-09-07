package main

import (
	"log"
	"saw_millers_bot/sheetsapi"
	"time"
	"math/rand"
	
	"saw_millers_bot/secret"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	sheetsapi.Init()

	bot, err := tgbotapi.NewBotAPI(secret.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Chat.ID != secret.ChatId && update.Message.Chat.ID != secret.DeveloperChatId { // for security
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var msg tgbotapi.MessageConfig
		command := update.Message.Command()
		if command == "dinnerduty" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetDinnerDuty(time.Now().Weekday()))
		} else if command == "dinnerdutytmr" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetDinnerDuty(time.Now().Add(time.Hour * 24).Weekday()))
		} else if command == "trashduty" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetTrashDuty())
		} else if command == "cleaningduty" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetCleaningDuty())
		} else if command == "nextcleaningdate" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetNextCleaningDate())
		} else if command == "help" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetHelp())
		} else if command == "flipcoin" {
 			var coin int = Intn(1) // 1 is heads, 0 is tails
 			var pmsg string
 			if coin {
 				pmsg = "Heads"
 			} else {
 				pmsg = "Tails"
 			}
 			msg = tgbotapi.NewMessage(update.Message.Chat.ID, pmsg)
		}

		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}
}
