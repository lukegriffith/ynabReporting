FROM golang:1.11

RUN go get github.com/lukemgriffith/ynabReporting/service.dailySpend/spending
WORKDIR /go/src/github.com/lukemgriffith/ynabReporting/service.dailySpend
RUN make build 