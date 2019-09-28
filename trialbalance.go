package main

import (
	"fmt"
	"sort"
	"strconv"
)

type trialBalance map[string]int

func (t *trialBalance) createTrialBalance(j journal) {
	maxLengthCredit := ""
	maxLengthDebit := ""
	tempTrialBalance := trialBalance{}
	for _, entry := range j {
		for index, debit := range entry.debitEntries {
			if _, ok := tempTrialBalance[debit]; ok {
				tempTrialBalance[debit] += entry.debitBalances[index]
			} else {
				tempTrialBalance[debit] = entry.debitBalances[index]
			}
			if len(debit) > len(maxLengthDebit) {
				maxLengthDebit = debit
			}
		}
		for index, credit := range entry.creditEntries {
			if _, ok := tempTrialBalance[credit]; ok {
				tempTrialBalance[credit] -= entry.creditBalances[index]
			} else {
				tempTrialBalance[credit] = -entry.creditBalances[index]
			}
			if len(credit) > len(maxLengthCredit) {
				maxLengthCredit = credit
			}
		}

	}
	(*t) = tempTrialBalance
}

func (t *trialBalance) writeTrialBalance() {
	// The Sort
	n := map[int][]string{}
	var a [] int
	for k, v := range *t {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// The Printing
	spaceString := "                                 "
	for _, k := range a {
		for _, s := range n[k] {
			newLine := ""
			if k < 0 {
				newLine = s + spaceString[len(s)-1:] + "    " + strconv.Itoa(k*-1)
			} else {
				newLine = s + spaceString[len(s)-1:] + strconv.Itoa(k)
			}
			fmt.Println(newLine)
		}
	}
}