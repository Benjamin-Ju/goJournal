package main

func (j journal) createTrialBalance() {
	trialBalance := make(map[string]int)
	for _, entry := range j {
		for index, debit := range entry.debitEntries {
			if _, ok := trialBalance[debit]; ok {
				trialBalance[debit] += entry.debitBalances[index]
			} else {
				trialBalance[debit] = entry.debitBalances[index]
			}
		}
		for index, credit := range entry.creditEntries {
			if _, ok := trialBalance[credit]; ok {
				trialBalance[credit] -= entry.creditBalances[index]
			} else {
				trialBalance[credit] = -entry.creditBalances[index]
			}
		}

	}
}
