package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
)

func main() {
	cache_url := "http://localhost:8080/"

	_ = Spending.NewCacheClient(cache_url)

	http.HandleFunc("/", serveSpending)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}

func serveSpending(w http.ResponseWriter, r *http.Request) {

	s, err := Spending.Get()

	if err != nil {
		log.Fatal("Process failed: ", err)
	}

	jsonString, err := json.Marshal(s)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(r.Method, " at ", r.RequestURI, " from ", r.RemoteAddr)

	w.Write(jsonString)

}
