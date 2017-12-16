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
	ticker := time.NewTicker(time.Hour)
	msg := tgbotapi.NewMessage(secret.ChatId, sheetsapi.GetTrashDuty())
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				t := time.Now()
				if t.Weekday() == time.Tuesday && t.Hour() == 22 {
					bot.Send(msg)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
