package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	journalExists := false
	trialBalanceExists := false
	myJournal := journal{}
	journalPointer := &myJournal
	myTrialBalance := trialBalance{}
	tbPointer := &myTrialBalance

	for true {
		fmt.Println("Enter a command")
		fmt.Println("j - start entering journal entries")
		fmt.Println("tb - create a trial balance")
		fmt.Println("ae - create adjusting entries")
		fmt.Println("bs - create a balance sheet")
		fmt.Println("is - create an income statement")
		fmt.Println("q - quit program")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		fmt.Println(input)
		// Entering Journalling Entries
		if input == "j" {
			// go to journaling interface
			journalExists = true
			journalPointer.enterEntries()
			// Posting to Ledger
		} else if input == "tb" {
			if journalExists {
				tbPointer.createTrialBalance(myJournal)
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			trialBalanceExists = true
			// Entering Adjusting Entries
		} else if input == "ae" {
			if journalExists {
				// go into journal interface
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			// make adjustments
			// create adjusted trial balance
			trialBalanceExists = true
			// Create Balance Sheet
		} else if input == "bs" {
			if trialBalanceExists {
				// create a balance sheet
			} else {
				fmt.Println("No trial balance exists. Balance sheet could not be created.")
			}
			// Create Income Statement
		} else if input == "is" {
			if trialBalanceExists {

			} else {
				fmt.Println("No trial balance exists. Income statement could not be created.")
			}
			// create a income statement
			// Close Books
		} else if input == "q" {
			os.Exit(0)
		} else {
			fmt.Println("Unaccepted input, please try again :)")
		}
	}
}
