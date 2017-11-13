package sheetsapi

import (
	"fmt"
	"log"
	"sawmillersbot/secret"
	"time"

	"strings"

	"google.golang.org/api/sheets/v4"
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

func GetQuoteOfTheDay() string {
	var res string
	readRange := "Quotes!A1:B1" // sheet and range
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			res += "Quote of the Day:\n\n" + fmt.Sprintf("\"%s\" (%s, 2017)\n", row[0], row[1])
		}
	} else {
		fmt.Print("No data found.")
	}
	return res
}

// Contract Bridge Functions

func GetBridgePlayers() []string {
	var res []string
	readRange := "Bridge Scoring!J2:O2" // sheet and range
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			for _, name := range row {
				res = append(res, name.(string))
			}
		}
	} else {
		fmt.Print("No data found.")
	}
	return res
}

func GetCurrBridgeHistRow() string {
	var num string
	readRange := "Bridge Scoring!J1:J1" // sheet and range
	resp, err := srv.Spreadsheets.Values.Get(secret.SpreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			num = row[0].(string)
		}
	} else {
		fmt.Print("No data found.")
	}
	return num
}

func WriteBid(data string) {
	writeBridgeScoring(data, "A")
}

func WriteWinner1(data string) {
	writeBridgeScoring(data, "B")
}

func WriteWinner2(data string) {
	writeBridgeScoring(data, "C")
}

func WriteLoser1(data string) {
	writeBridgeScoring(data, "D")
}

func WriteLoser2(data string) {
	writeBridgeScoring(data, "E")
}

func WriteWinDoubleStatus(data string) {
	writeBridgeScoring(data, "F")
}

func WriteVul(data string) {
	writeBridgeScoring(data, "G")
}

func writeBridgeScoring(data string, col string) {
	data = strings.Split(data, "_")[1]
	cell := col + GetCurrBridgeHistRow()

	rb := &sheets.ValueRange{}
	rb.Values = [][]interface{}{{data}}

	resp, err := srv.Spreadsheets.Values.Update(secret.SpreadsheetId, "Bridge Scoring!"+cell+":"+cell, rb).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", resp)
}
