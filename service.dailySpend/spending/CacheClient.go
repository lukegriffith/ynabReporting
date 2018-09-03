package Spending

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// Structure used for querying the cache.
type CacheClient struct {
	CacheUrl string
}

// Singleton variable
var client *CacheClient

// Singleton get
func GetCacheClient() (*CacheClient, error) {

	if client == nil {
		return &CacheClient{}, errors.New("No cache client initiated.")
	} else {
		return client, nil
	}
}

// Singleton new
func NewCacheClient(url string) *CacheClient {
	client = &CacheClient{url}

	return client
}

// query function
func (c *CacheClient) queryCache() (transactionsEnvelope, error) {

	req, err := http.NewRequest("GET", c.CacheUrl, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return transactionsEnvelope{}, err
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
			return transactionsEnvelope{}, errors.New("Non 20* response from cache.")
		}

		time.Sleep(3 * time.Second)

		resp, err = client.Do(req)

		status = resp.StatusCode

		if err != nil {
			log.Print("Do: ", err)
			return transactionsEnvelope{}, err
		}

	}

	defer resp.Body.Close()

	var record transactionsEnvelope

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record, nil

}
