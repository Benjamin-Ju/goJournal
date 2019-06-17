package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (t *trialBalance) createFinancialStatements() {
	reader := bufio.NewReader(os.Stdin)
	balanceSheet := make(map[string]int)
	incomeStatement := make(map[string]int)
	for account, value := range *t {
		temp := []string{"Which statement does ", account, " belong to?"}
		fmt.Println(strings.Join(temp, ""))
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		if input == "i" {
			incomeStatement[account] = value
		} else if input == "b" {
			balanceSheet[account] = value
		}
	}
}
