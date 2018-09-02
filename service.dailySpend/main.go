package main

import (
	"fmt"
	"github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
  "log"
  "encoding/json"
  "net/http"
)

func main() {

  http.HandleFunc("/", serveSpending)
  if err := http.ListenAndServe(":8081", nil); err != nil {
    panic(err)
  }



}

func serveSpending(w http.ResponseWriter, r *http.Request) {
  cache_url := "http://localhost:8080/"

  s, err := Spending.Get(cache_url)

	if err != nil {
		log.Fatal("Process failed: ", err)
	}

  jsonString, err := json.Marshal(s)

  if err != nil {
    fmt.Println(err)
    log.Fatal(err)
  }

  log.Print(r.Method, " at ", r.RequestURI, " from ", r.RemoteAddr)

  w.Write(jsonString)

}