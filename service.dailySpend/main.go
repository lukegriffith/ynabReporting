package main


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
  "time"
)

type TransactionsEnvelope struct {
	Data struct {
		Transactions []struct {
			ID                string        `json:"id"`
			Date              string        `json:"date"`
			Amount            int           `json:"amount"`
			Memo              string        `json:"memo"`
			Cleared           string        `json:"cleared"`
			Approved          bool          `json:"approved"`
			FlagColor         interface{}   `json:"flag_color"`
			AccountID         string        `json:"account_id"`
			AccountName       string        `json:"account_name"`
			PayeeID           string        `json:"payee_id"`
			PayeeName         string        `json:"payee_name"`
			CategoryID        string        `json:"category_id"`
			CategoryName      string        `json:"category_name"`
			TransferAccountID string        `json:"transfer_account_id"`
			ImportID          string        `json:"import_id"`
			Deleted           bool          `json:"deleted"`
			Subtransactions   []interface{} `json:"subtransactions"`
		} `json:"transactions"`
	} `json:"data"`
}


type DailySpending struct {

  TotalSpend, TotalTransactions *int
}

func (s *DailySpending) AddTo(spend int) {

  var spent = *s.TotalSpend + spend
  var trans = *s.TotalTransactions + 1 

  s.TotalSpend =  &spent
  s.TotalTransactions = &trans
}



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

  for k, v := range s {

    fmt.Println(k)

    fmt.Println("Average spend on day: £", (*v.TotalSpend / 1000 ) / *v.TotalTransactions) 
  }

}
