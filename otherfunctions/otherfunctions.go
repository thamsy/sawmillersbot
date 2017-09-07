package otherfunctions

import "math/rand"

func GetFlippedCoin() string {
	var coin int = rand.Intn(2)
	var pmsg string
	if coin == 1 {
		pmsg = "Heads"
	} else {
		pmsg = "Tails"
	}
	return pmsg
}

func GetHelp() string {
	return `/dinnerduty - Who's on Dinner Duty today?
	/dinnerdutytmr - Who's on Dinner Duty tomorrow?
	/trashduty - Who's supposed to take out the trash this week?
	/cleaningduty - Who's on what cleaning duty?
	/nextcleaningdate - When's the next date for cleaning the house?
	/flipcoin - 50% chance heads, 50% tails, or is it?`
}
