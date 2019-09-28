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
		fmt.Println("fs - create financial statements (balance sheet and income statement)")
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
				tbPointer.writeTrialBalance()
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			trialBalanceExists = true
			// Entering Adjusting Entries
		} else if input == "ae" {
			if journalExists {
				journalPointer.enterEntries()
				// make adjustments
				// create adjusted trial balance
				tbPointer.createTrialBalance(myJournal)
				tbPointer.writeTrialBalance()
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			// Create Balance Sheet
		} else if input == "fs" {
			if trialBalanceExists {
				// create a balance sheet
			} else {
				fmt.Println("No trial balance exists. Balance sheet could not be created.")
			}
		} else if input == "q" {
			os.Exit(0)
		} else {
			fmt.Println("Unaccepted input, please try again :)")
		}
	}
}
