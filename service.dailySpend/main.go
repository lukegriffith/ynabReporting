package main


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "time"

  "github.com/lukemgriffith/ynabReporting/service.dailySpend/Spending"
)






func main() {
  Spending.Process("http://localhost:8080/")
}
