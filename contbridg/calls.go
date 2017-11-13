package contbridg

import (
	"strconv"

	"sawmillersbot/sheetsapi"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	ChooseBidPrefix          = "bid_"
	ChooseBidWinners1Prefix  = "bidwinners1_"
	ChooseBidWinners2Prefix  = "bidwinners2_"
	ChooseBidLosers1Prefix   = "bidlosers1_"
	ChooseBidLosers2Prefix   = "bidlosers2_"
	ChooseWinStatusPrefix    = "winstatus_"
	ChooseOverUnderPrefix    = "overunder_"
	ChooseIsVulnerablePrefix = "vul_"
)

var (
	suits       = []string{"C", "D", "H", "S", "NT"}
	suits_emoji = []string{"\xE2\x99\xA3", "\xE2\x99\xA6", "\xE2\x99\xA5", "\xE2\x99\xA0", " NT"}
	win_status  = []string{"Win", "Lose"}
	doub_status = []string{"", "(Doubled)", "(ReDoub)"}
)

func ChooseBid() tgbotapi.InlineKeyboardMarkup {
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()
	for num := 1; num <= 7; num++ {
		inlineKeyboardRows := tgbotapi.NewInlineKeyboardRow()
		strNum := strconv.Itoa(num)
		for idx, suit := range suits {
			inlineKeyboardButton := tgbotapi.NewInlineKeyboardButtonData(strNum+suits_emoji[idx], ChooseBidPrefix+strNum+suit)
			inlineKeyboardRows = append(inlineKeyboardRows, inlineKeyboardButton)
		}
		inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, inlineKeyboardRows)
	}
	finalRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Pass", "pass"),
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	)
	inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, finalRow)
	return inlineKeyboardMarkup
}

func ChooseBidPlayers(prefix string) *tgbotapi.InlineKeyboardMarkup {
	var players = sheetsapi.GetBridgePlayers()
	var inlineKeyboardRow = tgbotapi.NewInlineKeyboardRow()

	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()
	for idx, player := range players {
		if idx%2 == 0 {
			inlineKeyboardRow = tgbotapi.NewInlineKeyboardRow()
		}
		inlineKeyboardRow = append(inlineKeyboardRow, tgbotapi.NewInlineKeyboardButtonData(player, prefix+player))
		if idx%2 == 1 || idx == len(players)-1 {
			inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, inlineKeyboardRow)
		}
	}
	finalRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	)
	inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, finalRow)
	return &inlineKeyboardMarkup
}

func ChooseWinStatus() *tgbotapi.InlineKeyboardMarkup {
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()
	for _, winstate := range win_status {
		ikr := tgbotapi.NewInlineKeyboardRow()
		for _, doubstate := range doub_status {
			ikb := tgbotapi.NewInlineKeyboardButtonData(winstate+doubstate, ChooseWinStatusPrefix+winstate+doubstate)
			ikr = append(ikr, ikb)
		}
		inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, ikr)
	}
	finalRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	)
	inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, finalRow)
	return &inlineKeyboardMarkup
}

func ChooseOverUnder() *tgbotapi.InlineKeyboardMarkup {
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()

	var inlineKeyboardRow = tgbotapi.NewInlineKeyboardRow()
	for i := 0; i <= 7; i++ {
		if i%2 == 0 {
			inlineKeyboardRow = tgbotapi.NewInlineKeyboardRow()
		}
		inlineKeyboardRow = append(inlineKeyboardRow, tgbotapi.NewInlineKeyboardButtonData(
			strconv.Itoa(i), ChooseOverUnderPrefix+strconv.Itoa(i)))
		if i%2 == 1 {
			inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, inlineKeyboardRow)
		}
	}
	finalRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	)
	inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, finalRow)
	return &inlineKeyboardMarkup
}

func ChooseIsVulnerable() *tgbotapi.InlineKeyboardMarkup {
	ikr := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Yes", ChooseIsVulnerablePrefix+"Yes"),
		tgbotapi.NewInlineKeyboardButtonData("No", ChooseIsVulnerablePrefix+"No"),
	)
	finalRow := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cancel", "cancel"),
	)
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(ikr, finalRow)
	return &inlineKeyboardMarkup
}
