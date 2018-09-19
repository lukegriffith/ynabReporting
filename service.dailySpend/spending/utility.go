package spending

import (
	"log"
	"time"
)

func getDate(rawdate string) time.Time {

	date := rawdate + "T00:00:00+00:00"

	parsedDate, e := time.Parse(time.RFC3339, date)

	if e != nil {
		log.Fatal("Error: ", e)
	}

	return parsedDate

}
