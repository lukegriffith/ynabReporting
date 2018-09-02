package main


import (
  "fmt"
  "github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
)






func main() {
    
  s := Spending.Process("http://localhost:8080/")

  for k, v := range s {
  
	  fmt.Println(k)
	  fmt.Println("Average spend on day: £", (*v.TotalSpend / 1000 ) / *v.Totaltransactions) 
	}
}
