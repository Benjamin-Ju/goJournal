package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	unacceptedInput := true
	journalExists := false
	trialBalanceExists := false

	for unacceptedInput {
		fmt.Println("Enter a command")
		fmt.Println("j - start entering journal entries")
		fmt.Println("p - post journal entries")
		fmt.Println("tb - create a trial balance")
		fmt.Println("w - create a worksheet")
		fmt.Println("ae - create adjusting entries")
		fmt.Println("bs - create a balance sheet")
		fmt.Println("fs - create a financial statement")
		fmt.Println("c - close books")
		input, _ := fmt.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		fmt.Println(input)
		if input == "j" {
			// go to journaling interface
			journalExists = true
			unacceptedInput = false

		} else if input == "p" {
			if journalExists {
				// post
			} else {
				fmt.Println("No journal entries exist. Posting could be not happen.")
			}
			// post journal
			unacceptedInput = false
		} else if input == "tb" {
			if journalExists {
				// create a trial balance
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			unacceptedInput = false
			trialBalanceExists = true
		} else if input == "w" {
			if journalExists {
				// create a worksheet
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			// create a worksheet
			unacceptedInput = false
		} else if input == "ae" {
			if journalExists {
				// go into journal interface
			} else {
				fmt.Println("No journal entries exist. Trial balance could not be created.")
			}
			// make adjustments
			// create adjusted trial balance
			unacceptedInput = false
			trialBalanceExists = true
		} else if input == "bs" {
			if trialBalanceExists {
				// create a balance sheet
			} else {
				fmt.Println("No trial balance exists. Balance sheet could not be created.")
			}

			unacceptedInput = false
		} else if input == "fs" {
			if trialBalanceExists {

			} else {
				fmt.Println("No trial balance exists. Financial statement could not be created.")
			}
			// create a financial statement
			unacceptedInput = false
		} else if input == "c" {
			if trialBalanceExists {
				//
			} else {
				fmt.Println("No trial balance exists. Financial statement could not be created.")
			}
			// closes accounts
			unacceptedInput = false
			journalExists = false
			trialBalanceExists = false
		}
	}
}
