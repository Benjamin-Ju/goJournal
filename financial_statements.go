package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (t *trialBalance) createFinancialStatements() {
	reader := bufio.NewReader(os.Stdin)
	balanceSheet := make(map[string]map[string]int)
	incomeStatement := make(map[string]map[string]int)
	for account, value := range *t {
		temp := []string{"Which statement does ", account, " belong to?"}
		fmt.Println(strings.Join(temp, ""))
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if input == "i" {
			fmt.Println("Is it income (i) or an expense? (e)?")
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			if input == "i" {
				incomeStatement["income"][account] = value
			}
			if input == "e" {
				incomeStatement["expense"][account] = value
			}
		} else if input == "b" {
			fmt.Println("Is it an asset (a), liability (l), or part of owner's equity (e)?")
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			if input == "a" {
				balanceSheet["assets"][account] = value
			}
			if input == "l" {
				balanceSheet["liability"][account] = value
			}
			if input == "e" {
				balanceSheet["equity"][account] = value
			}
		}
	}
}
