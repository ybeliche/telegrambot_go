ARG image_name=ybeliche_telegrambot
ARG build_image=golang:latest
ARG base_image=alpine:3.14

FROM ${build_image} AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

ARG image_name
RUN GOOS=linux GOARCH=amd64 go build -o image_name main.go

ARG base_image
FROM ${base_image}

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/telegrambot /usr/local/bin/telegrambot

CMD ["telegrambot"]
