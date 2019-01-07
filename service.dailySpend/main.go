package main

import (
	"fmt"
	"net/http"

	"github.com/lukemgriffith/ynabReporting/service.dailySpend/spending"
)

func main() {
	cacheURL := "http://localhost:32777/"

	port := ":8081"

	fmt.Println("Webserver running on port " + port)
	_ = spending.NewCacheClient(cacheURL, "", false)

	http.HandleFunc("/avgSpending", spending.GetAverageDailySpending)
	http.HandleFunc("/last7", spending.GetLast7Days)
	http.HandleFunc("/account", spending.Account)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
