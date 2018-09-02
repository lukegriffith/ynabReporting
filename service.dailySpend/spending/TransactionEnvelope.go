package Spending

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
