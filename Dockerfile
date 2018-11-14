FROM golang:1.11 as builder

WORKDIR /usr/src

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o daemon .

FROM alpine:latest

RUN apk update && apk add --no-cache --virtual ca-certificates

WORKDIR /usr/app
COPY --from=builder /usr/src/ .

CMD ["./daemon", "config.json"]
