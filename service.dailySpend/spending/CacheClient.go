package spending

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// CacheClient structure used for querying the cache.
type CacheClient struct {
	CacheURL     string
	AccountName  string
	CheckAccount bool
}

// Singleton variable
var client *CacheClient

// GetCacheClient singleton get
func GetCacheClient() (*CacheClient, error) {

	if client == nil {
		return &CacheClient{}, errors.New("no cache client initiated")
	}

	return client, nil

}

// NewCacheClient singleton new
func NewCacheClient(url string, account string) *CacheClient {

	client = &CacheClient{url, account, false}

	return client
}

// QueryCache queries the cache
func (c *CacheClient) QueryCache() ([]Transaction, error) {

	req, err := http.NewRequest("GET", c.CacheURL, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Print("Do: ", err)
	}

	status := resp.StatusCode

	for {
		if status == 200 {
			break
		}

		if status > 300 {
			return nil, errors.New("non 20* response from cache")
		}

		time.Sleep(3 * time.Second)

		resp, err = client.Do(req)

		status = resp.StatusCode

		if err != nil {
			log.Print("Do: ", err)
			return nil, err
		}

	}

	defer resp.Body.Close()

	var record transactionsEnvelope

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	var transactions []Transaction

	for _, e := range record.Data.Transactions {
		if c.CheckAccount && e.AccountName == c.AccountName || !c.CheckAccount {

			transactions = append(transactions, Transaction(e))
		}
	}

	return transactions, nil
}

func (c *CacheClient) getAccounts() []string {

	var accMap []string

	c.CheckAccount = false

	transactions, err := c.QueryCache()

	if err != nil {
		log.Panic("Unable to query cache.")
	}

	// Iterate all transactions, to determine the set of account names.
	for _, t := range transactions {
		// Use boolean to determine if loop should be continued
		found := false

		for _, a := range accMap {
			// Check each account in accMap, compare to current transaction account name.
			if a == t.AccountName {
				// If found, set found to true and break loop.
				found = true
				break
			}

		}

		if found {
			// Continue to next transaction if AccountName is found.
			continue
		}

		accMap = append(accMap, t.AccountName)
	}

	return accMap
}
