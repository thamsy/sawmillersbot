package main

import (
	"log"
	"sawmillersbot/contbridg"
	"sawmillersbot/sheetsapi"
	"time"

	"sawmillersbot/otherfunctions"
	"sawmillersbot/secret"

	"strings"

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

		var msg tgbotapi.MessageConfig
		if update.Message != nil && update.Message.IsCommand() {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if update.Message.Chat.ID != secret.ChatId && update.Message.Chat.ID != secret.DeveloperChatId { // for security
				continue
			}
			command := update.Message.Command()
			if command == "dinnerduty" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetDinnerDuty(time.Now().Weekday()))
			} else if command == "dinnerdutytmr" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetDinnerDuty(time.Now().Add(time.Hour*24).Weekday()))
			} else if command == "trashduty" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetTrashDuty())
			} else if command == "cleaningduty" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetCleaningDuty())
			} else if command == "nextcleaningdate" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetNextCleaningDate())
			} else if command == "incrementschedule" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.IncrementSchedule())
			} else if command == "help" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, otherfunctions.GetHelp())
			} else if command == "flipcoin" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, otherfunctions.GetFlippedCoin())
			} else if command == "qotd" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, sheetsapi.GetQuoteOfTheDay())
			} else if command == "inputcontbridgscore" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Select Winning Bid:")
				msg.ReplyMarkup = contbridg.ChooseBid()
			}
		} else {
			callbackQuery := update.CallbackQuery
			if callbackQuery.Message.ReplyToMessage.IsCommand() {
				switch command := callbackQuery.Message.ReplyToMessage.Command(); command {
				case "inputcontbridgscore":
					data := callbackQuery.Data
					if data == "cancel" {
						var deletemsg tgbotapi.DeleteMessageConfig
						deletemsg.ChatID = callbackQuery.Message.Chat.ID
						deletemsg.MessageID = callbackQuery.Message.MessageID
						bot.DeleteMessage(deletemsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseBidPrefix) {
						go sheetsapi.WriteBid(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "Choose Winner 1 of Bid"
						editmsg.ReplyMarkup = contbridg.ChooseBidPlayers(contbridg.ChooseBidWinners1Prefix)
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseBidWinners1Prefix) {
						go sheetsapi.WriteWinner1(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "Choose Winner 2 of Bid"
						editmsg.ReplyMarkup = contbridg.ChooseBidPlayers(contbridg.ChooseBidWinners2Prefix)
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseBidWinners2Prefix) {
						go sheetsapi.WriteWinner2(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "Choose Loser 1 of Bid"
						editmsg.ReplyMarkup = contbridg.ChooseBidPlayers(contbridg.ChooseBidLosers1Prefix)
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseBidLosers1Prefix) {
						go sheetsapi.WriteLoser1(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "Choose Loser 2 of Bid"
						editmsg.ReplyMarkup = contbridg.ChooseBidPlayers(contbridg.ChooseBidLosers2Prefix)
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseBidLosers2Prefix) {
						sheetsapi.WriteLoser2(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "Choose Win Status"
						editmsg.ReplyMarkup = contbridg.ChooseWinStatus()
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseWinStatusPrefix) {
						go sheetsapi.WriteWinDoubleStatus(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "How many Over or Under?"
						editmsg.ReplyMarkup = contbridg.ChooseOverUnder()
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseOverUnderPrefix) {
						go sheetsapi.WriteOverUnder(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "isVulnerable?"
						editmsg.ReplyMarkup = contbridg.ChooseIsVulnerable()
						bot.Send(editmsg)
						continue
					} else if strings.HasPrefix(data, contbridg.ChooseIsVulnerablePrefix) {
						go sheetsapi.WriteVul(data)
						var editmsg tgbotapi.EditMessageTextConfig
						editmsg.ChatID = callbackQuery.Message.Chat.ID
						editmsg.MessageID = callbackQuery.Message.MessageID
						editmsg.Text = "Entry Recorded!"
						bot.Send(editmsg)
						continue
					}
				}
			}

		}

		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}
}
