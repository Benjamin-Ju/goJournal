package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type journalEntry struct {
	var date           string
	var debitEntries   []string
	var creditEntries  []string
	var debitBalances  []int
	var creditBalances []int
}

type journal []journalEntry

func (j journal) enterEntries() {
	reader := bufio.NewReader(os.Stdin)
	stay := true
	for stay {
		fmt.Println("Day of entry:")
		date, _ := reader.ReadString('\n')
		date = strings.Replace(date, "\n", "", -1)
		var debitAccounts []string
		var creditAccounts []string
		var debitValues []int
		var creditValues []int
		stillDebit := true
		stillCredit := true
		fmt.Println("Enter your debit entries, ")
		for stillDebit {
			fmt.Println("Enter account:")
			account, _ := reader.ReadString('\n')
			account = strings.Replace(account, "\n", "", -1)
			fmt.Println("Enter value:")
			value, _ := reader.ReadString('\n')
			value = strings.Replace(value, "\n", "", -1)
			valueInt = int(value)
		}
	}

}
