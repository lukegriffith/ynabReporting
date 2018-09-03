package main

import (
	"net/http"

	"github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
)

func main() {
	cache_url := "http://localhost:8080/"

	_ = Spending.NewCacheClient(cache_url)

	http.HandleFunc("/avgSpending", Spending.GetAverageDailySpending)
	http.HandleFunc("/last7", Spending.GetLast7Days)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
