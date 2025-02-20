FROM golang:1.8

WORKDIR /go/src/app
COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

RUN go build ./cmd/main.go

CMD ["./main"]
