package spending

import (
	"github.com/mitchellh/mapstructure"
)

type transactionsEnvelope struct {
	Data struct {
		Transactions []struct {
			ID                string        `json:"id"`
			Date              string        `json:"date"`
			Amount            int           `json:"amount"`
			Memo              string        `json:"memo"`
			Cleared           string        `json:"cleared"`
			Approved          bool          `json:"approved"`
			FlagColor         interface{}   `json:"flag_color"`
			AccountID         string        `json:"account_id"`
			AccountName       string        `json:"account_name"`
			PayeeID           string        `json:"payee_id"`
			PayeeName         string        `json:"payee_name"`
			CategoryID        string        `json:"category_id"`
			CategoryName      string        `json:"category_name"`
			TransferAccountID string        `json:"transfer_account_id"`
			ImportID          string        `json:"import_id"`
			Deleted           bool          `json:"deleted"`
			Subtransactions   []interface{} `json:"subtransactions"`
		} `json:"transactions"`
	} `json:"data"`
}

type Transaction struct {
	ID                string
	Date              string
	Amount            int
	Memo              string
	Cleared           string
	Approved          bool
	FlagColor         interface{}
	AccountID         string
	AccountName       string
	PayeeID           string
	PayeeName         string
	CategoryID        string
	CategoryName      string
	TransferAccountID string
	ImportID          string
	Deleted           bool
	Subtransactions   []interface{}
}

func (t *transactionsEnvelope) getAccounts() []string {

	var accMap []string

	for _, r := range t.Data.Transactions {

		for _, a := range accMap {

			if a == r.AccountName {
				continue
			}

			accMap = append(accMap, r.AccountName)
		}
	}

	return accMap
}

func (t *transactionsEnvelope) getTransactionsForAccount(AccountName string) []Transaction {

	var transactions []Transaction

	for _, e := range t.Data.Transactions {
		if e.AccountName == AccountName {

			var trn Transaction

			if err := mapstructure.Decode(e, trn); err != nil {
				panic(err)
			}

			transactions = append(transactions, trn)
		}
	}

	return transactions
}
