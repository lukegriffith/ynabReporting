package spending

// Transaction represents a single transaction from YNAB.
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
