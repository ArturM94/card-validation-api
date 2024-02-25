FROM golang:1.22.0

WORKDIR ./

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY validator ./validator

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app

EXPOSE 8080

CMD ["/bin/app"]
