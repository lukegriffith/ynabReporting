package main


import (
  "github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
)






func main() {
  Spending.Process("http://localhost:8080/")
}
