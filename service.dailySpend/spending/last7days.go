package spending

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"time"
)

// GetLast7Days is the controller for /7days
func GetLast7Days(w http.ResponseWriter, r *http.Request) {

	s, err := last7Days()

	if err != nil {
		log.Fatal("last7Days failed: ", err)
	}

	jsonString, err := json.Marshal(s)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(r.Method, " at ", r.RequestURI, " from ", r.RemoteAddr)

	w.Write(jsonString)

}

func last7Days() (map[string]float64, error) {

	cacheClient, err := GetCacheClient()

	if err != nil {
		log.Fatal(err)
	}

	record, err := cacheClient.queryCache()

	if err != nil {
		return nil, errors.New("Unable to query cache")
	}

	var s map[string]float64

	s = make(map[string]float64)

	// Horrible hack, YNAB does not do hours, aslong as its not midnight on the dot
	// 8 days will work for 7.
	lastWeek := time.Now().AddDate(0, 0, -8)
	now := time.Now()

	for _, t := range record.Data.Transactions {
		// Only count if an outgoing.
		if t.Amount < 0 {
			// Ynab date structure does not track time.
			date := t.Date + "T00:00:00+00:00"

			parsedDate, e := time.Parse(time.RFC3339, date)

			if e != nil {
				log.Fatal("Error: ", e)
			}

			if parsedDate.After(lastWeek) && parsedDate.Before(now) {
				if val, ok := s[parsedDate.Format("02-01-2006")]; ok {
					s[parsedDate.Format("02-01-2006")] = val + math.Abs((float64(t.Amount) / 1000))
				} else {
					s[parsedDate.Format("02-01-2006")] = math.Abs((float64(t.Amount) / 1000))
				}
			}

		}
	}

	return s, nil

}
