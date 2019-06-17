package main

type trialBalance map[string]int

func (t *trialBalance) createTrialBalance(j journal) {
	tempTrialBalance := trialBalance{}
	for _, entry := range j {
		for index, debit := range entry.debitEntries {
			if _, ok := tempTrialBalance[debit]; ok {
				tempTrialBalance[debit] += entry.debitBalances[index]
			} else {
				tempTrialBalance[debit] = entry.debitBalances[index]
			}
		}
		for index, credit := range entry.creditEntries {
			if _, ok := tempTrialBalance[credit]; ok {
				tempTrialBalance[credit] -= entry.creditBalances[index]
			} else {
				tempTrialBalance[credit] = -entry.creditBalances[index]
			}
		}

	}
	(*t) = tempTrialBalance
}
