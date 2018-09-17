package spending

type account struct {
	Accounts   [][]dailySpending
	AccountMap []string
}

func (acc *account) getAccounts() []string {
	var accs []string

	for _, v := range acc.AccountMap {
		accs = append(accs, v)
	}

	return accs
}

func (acc *account) getAccountSpendings(n int) []dailySpending {
	return acc.Accounts[n]
}
