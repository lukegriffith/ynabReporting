package Spending

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"errors"
	"math"
)

func Process(url string) (map[string]int, error) {

	req, err := http.NewRequest("GET", url, nil)

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
			return nil, errors.New("Non 20* response from cache.")
		}

		time.Sleep(500 * time.Millisecond)

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

	var s map[string]*DailySpending

	s = make(map[string]*DailySpending)

	for _, t := range record.Data.Transactions {
		// Only count if an outgoing.
		if t.Amount < 0 {
			// Ynab date structure does not track time.
			date := t.Date + "T00:00:00+00:00"

			parsedDate, e := time.Parse(time.RFC3339, date)

			if e != nil {
				log.Fatal("Error: ", e)
			}

			if val, ok := s[parsedDate.Format("Mon")]; ok {
				val.AddTo(t.Amount)
			} else {
				var amount = 0 + t.Amount
				var trans = 1
				s[parsedDate.Format("Mon")] = &DailySpending{&amount, &trans}

			}
		}
	}

	var result map[string]int

	result = make(map[string]int)

	for k, v := range s {
		// determine average and add to result map.
		avg := math.Abs((*v.TotalSpend/1000) / *v.TotalTransactions)
		result[k] = avg
	}

	return result, nil

}