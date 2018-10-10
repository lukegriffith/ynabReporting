package main

import (
	"net/http"

	"github.com/lukemgriffith/ynabReporting/service.dailySpend/spending"
)

func main() {
	cacheURL := "http://localhost:8080/"

	_ = spending.NewCacheClient(cacheURL)

	http.HandleFunc("/avgSpending", spending.GetAverageDailySpending)
	http.HandleFunc("/last7", spending.GetLast7Days)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
