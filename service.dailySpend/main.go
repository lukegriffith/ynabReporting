package main

import (
	"fmt"
	"github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
  "log"
  "encoding/json"
)

func main() {

  cache_url := "http://localhost:8080/"

	s, err := Spending.Process(cache_url)

	if err != nil {
		log.Fatal("Process failed: ", err)
	}

  jsonString, err := json.Marshal(s)

  if err != nil {
    fmt.Println(err)
    log.Fatal("Unable to seralize json ", err)
  }

  fmt.Println(string(jsonString))

}
