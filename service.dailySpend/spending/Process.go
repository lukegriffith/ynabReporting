package Spending

import (
	"errors"
	"log"
	"math"
	"time"

)

func Get() (map[string]float64, error) {

	cacheClient, err := Spending.GetCacheClient()

	if err != nil {
		log.Fatal(err)
	}

	record, err := cacheClient.queryCache()

	if err != nil {
		return nil, errors.New("Unable to query cache")
	}

	var s map[string]*dailySpending

	s = make(map[string]*dailySpending)

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
				s[parsedDate.Format("Mon")] = &dailySpending{&amount, &trans}

			}
		}
	}

	var result map[string]float64

	result = make(map[string]float64)

	for k, v := range s {
		// determine average and add to result map.
		avg := math.Abs(float64((*v.TotalSpend / 1000) / *v.TotalTransactions))
		result[k] = avg
	}

	return result, nil

}
