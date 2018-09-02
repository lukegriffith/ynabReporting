package main


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "time"

  "github.com/lukemgriffith/ynabReporting/service.dailySpend/spending"
)






func main() {


  url := fmt.Sprintf("http://localhost:8080/")

  req, err := http.NewRequest("GET", url, nil)


	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
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

    time.Sleep(500 * time.Millisecond)

    resp, err = client.Do(req)

    status = resp.StatusCode

    if err != nil {
      log.Print("Do: ", err)
    }

  }

  fmt.Println("200 recieved")

  defer resp.Body.Close()

  var record TransactionsEnvelope

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}


  var s map[string]*DailySpending

  s = make(map[string]*DailySpending)

  for _, t := range record.Data.Transactions {
    # Only count if an outgoing.
    if t.Amount < 0 { 

      date := t.Date + "T00:00:00+00:00"

      parsedDate, e := time.Parse(time.RFC3339, date)

      if e != nil {
        log.Fatal("Error: ", e)
      }


      if val, ok := s[parsedDate.Format("Mon")]; ok {
        val.AddTo(t.Amount)
      } else {

        var amount = t.Amount
        var trans = 1
        s[parsedDate.Format("Mon")] = &DailySpending{&amount, &trans}
      }
      
    }
  }

  for k, v := range s {

    fmt.Println(k)

    fmt.Println("Average spend on day: £", (*v.TotalSpend / 1000 ) / *v.TotalTransactions) 
  }

}
