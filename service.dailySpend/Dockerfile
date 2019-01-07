FROM golang:1.11 as builder

RUN go get github.com/lukemgriffith/ynabReporting/service.dailySpend/spending
WORKDIR /go/src/github.com/lukemgriffith/ynabReporting/service.dailySpend
RUN make


FROM alpine:latest

COPY --from=builder /go/bin/dailySpend /go/bin/dailySpend
CMD ["/go/bin/dailySpend"]