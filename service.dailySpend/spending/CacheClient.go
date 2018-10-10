package spending

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
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

// query function
func (c *CacheClient) queryCache() ([]Transaction, error) {

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
		if c.CheckAccount && e.AccountName == c.AccountName {

			var trn Transaction

			if err := mapstructure.Decode(e, trn); err != nil {
				panic(err)
			}

			transactions = append(transactions, trn)
		}
	}

	return transactions, nil
}

func (c *CacheClient) getAccounts() []string {

	var accMap []string

	c.CheckAccount = false

	accs, err := c.queryCache()

	if err != nil {
		log.Panic("Unable to query cache.")
	}

	// Not sure this works. Meant to be iterating through accounts and adding to an array if it doesn't exist.
	for _, r := range accs {

		found := false

		for _, a := range accMap {

			if a == r.AccountName {
				found = true
				break
			}

		}

		accMap = append(accMap, r.AccountName)
	}

	return accMap
}
