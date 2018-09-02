package main

import (
	"fmt"
	"github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
  "log"
  "encoding/json"
)

func main() {

	s, err := Spending.Process("http://localhost:8080/")

	if err != nil {
		log.Fatal("Process failed: ", err)
	}

  jsonString, err := json.Marshal(s)

  if err != nil {
    log.Fatal("Unable to seralize json ", err)
  }

  fmt.Println(jsonString)

}
