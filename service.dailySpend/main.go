package main

import (
	"fmt"
	"github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
	"log"
)

func main() {

	s, err := Spending.Process("http://localhost:8080/")

	if err != nil {
		log.Fatal("Process failed: ", err)
	}

	for k, v := range s {

		fmt.Println(k)
		fmt.Println("Average spend on day: £", (*v.TotalSpend/1000) / *v.Totaltransactions)
	}
}
