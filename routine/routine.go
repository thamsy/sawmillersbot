package routine

import (
	"sawmillersbot/secret"
	"sawmillersbot/sheetsapi"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	bot *tgbotapi.BotAPI
)

func StartRoutine(b *tgbotapi.BotAPI) {
	bot = b
	trashDutyReminder()
}

func trashDutyReminder() {
	ticker := time.NewTicker(5 * time.Second)
	var msg tgbotapi.MessageConfig
	msg = tgbotapi.NewMessage(secret.DeveloperChatId, sheetsapi.GetTrashDuty())
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				t := time.Now()
				if t.Weekday() == time.Saturday && t.Hour() == 12 {
					//if t.Weekday() == time.Tuesday && t.Hour() == 22 {
					bot.Send(msg)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
