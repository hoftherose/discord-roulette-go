FROM golang:1.19.0-buster as builder
WORKDIR /usr/src/app

COPY go.mod .
RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 go build ./main.go

FROM alpine:3.16.2

COPY --from=builder /usr/src/app/main /usr/bin

CMD ["main"]