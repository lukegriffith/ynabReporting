# Ynab Experiment

A number of services relating to my YNAB budgets. Creating new views from the data stored in my finances.


## service.ynabCache
Caching service that stores a json blob in memory, refreshes on certain logic.

Mainly used to stop hammering the YNAB api's 

## service.dailySpend
Calculates the daily spendings to show me how I on average spend my money during the week.
