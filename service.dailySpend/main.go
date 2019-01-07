package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lukemgriffith/ynabReporting/service.dailySpend/spending"
)

func testClient(w http.ResponseWriter, r *http.Request) {
	cacheClient, err := spending.GetCacheClient()

	if err != nil {
		log.Fatal(err)
	}

	s, err := cacheClient.QueryCache()

	jsonString, err := json.Marshal(s)	

	fmt.Println(s)

	w.Write(jsonString)

}

func main() {
	cacheURL := "http://localhost:32776/"

	port := ":8081"

	fmt.Println("Webserver running on port " + port)
	_ = spending.NewCacheClient(cacheURL, "")

	http.HandleFunc("/avgSpending", spending.GetAverageDailySpending)
	http.HandleFunc("/last7", spending.GetLast7Days)
	http.HandleFunc("/test", testClient)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
