package spending

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
)

// GetAverageDailySpending is the controller  for responding to HTTP requests.
// at /avgSpending
func GetAverageDailySpending(w http.ResponseWriter, r *http.Request) {

	s, err := averageDailySpending()

	if err != nil {
		log.Fatal("averageDailySpending failed: ", err)
	}

	jsonString, err := json.Marshal(s)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(r.Method, " at ", r.RequestURI, " from ", r.RemoteAddr)

	w.Write(jsonString)

}

// Queries cache and calculates the average daily spend.
func averageDailySpending() (map[string]float64, error) {

	cacheClient, err := GetCacheClient()

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
			parsedDate := getDate(t.Date)

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
