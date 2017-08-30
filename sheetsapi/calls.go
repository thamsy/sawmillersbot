package sheetsapi

import (
	"fmt"
	"log"
	"time"
	"saw_millers_bot/secret"
)

var (
	dinnerDutyTypes []string = []string{"Chef 1", "Chef 2", "Wash"}
	dinnerDutyCols  []string = []string{"", "B", "C", "D", "E", "F", "G"}
)

func GetDinnerDuty(weekday time.Weekday) string {
	if weekday == 0 {
		return "It's Sunday, no duty!"
	}

	var res string
	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	col := dinnerDutyCols[weekday]
	readRange := "Cooking!" + col + "2:4"
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		for i, row := range resp.Values {
			res += dinnerDutyTypes[i] + fmt.Sprintf(": %s\n", row[0])
		}
	} else {
		fmt.Print("No data found.")
	}

	return res
}

func GetDinnerDutyTmr(weekday time.Weekday) string {
	return GetDinnerDuty((weekday + 1) % 7)
}

func GetTrashDuty() string {
	var res string
	readRange := "House Cleaning!A12"
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			res += "Trash Duty: " + fmt.Sprintf("%s\n", row[0])
		}
	} else {
		fmt.Print("No data found.")
	}
	return res
}

func GetCleaningDuty() string {
	var res string
	readRange := "House Cleaning!A5:B12"
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			res += fmt.Sprintf("%s - %s\n", row[1], row[0])
		}
	} else {
		fmt.Print("No data found.")
	}
	return res
}

func GetNextCleaningDate() string {
	var res string
	readRange := "House Cleaning!B1:C1"
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			res += "Next Cleaning Date: " + fmt.Sprintf("%s %s\n", row[0], row[1])
		}
	} else {
		fmt.Print("No data found.")
	}
	return res
}

func GetHelp() string {
	return `/dinnerduty - Who's on Dinner Duty today?
	/dinnerdutytmr - Who's on Dinner Duty tomorrow?
	/trashduty - Who's supposed to take out the trash this week?
	/cleaningduty - Who's on what cleaning duty?
	/nextcleaningdate - When's the next date for cleaning the house?`
}
