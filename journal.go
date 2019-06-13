package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type journalEntry struct {
	date           string
	debitEntries   []string
	creditEntries  []string
	debitBalances  []int
	creditBalances []int
}

type journal []journalEntry

func enterEntries() {
	reader := bufio.NewReader(os.Stdin)
	stay := true
	for stay {
		// date gets entered
		fmt.Println("Day of entry:")
		date, _ := reader.ReadString('\n')
		date = strings.Replace(date, "\n", "", -1)
		// declarations required for accounts
		var debitAccounts []string
		var creditAccounts []string
		var debitValues []int
		var creditValues []int
		stillDebit := true
		stillCredit := true
		// entering debit entries
		fmt.Println("Enter your debit entries, ")
		for stillDebit {
			fmt.Println("Enter account:")
			account, _ := reader.ReadString('\n')
			account = strings.Replace(account, "\n", "", -1)
			fmt.Println("Enter value:")
			value, _ := reader.ReadString('\n')
			value = strings.Replace(value, "\n", "", -1)
			valueInt, _ := strconv.Atoi(value)
			// if a space exists as a prefix, it is now considered a credit entry
			if strings.HasPrefix(account, " ") {
				stillDebit = false
				creditAccounts = append(creditAccounts, account)
				creditValues = append(creditValues, valueInt)
			} else {
				debitAccounts = append(debitAccounts, account)
				debitValues = append(debitValues, valueInt)
			}
		}
		for stillCredit {
			fmt.Println("Enter account: (enter q to quit)")
			account, _ := reader.ReadString('\n')
			account = strings.Replace(account, "\n", "", -1)
			if account == "q" {
				stillCredit = false
				break
			}
			creditAccounts = append(creditAccounts, account)
			fmt.Println("Enter value:")
			value, _ := reader.ReadString('\n')
			value = strings.Replace(value, "\n", "", -1)
			valueInt, _ := strconv.Atoi(value)
			creditValues = append(creditValues, valueInt)
		}
		temp := journalEntry{
			date:           date,
			debitEntries:   debitAccounts,
			debitBalances:  debitValues,
			creditEntries:  creditAccounts,
			creditBalances: creditValues,
		}
		// j = append(j, temp)
		fmt.Println(temp)
	}

}
