ARG image_name=ybeliche_telegrambot
ARG build_image=golang:latest
ARG base_image=alpine:3.14

FROM ${build_image} AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

ARG image_name
RUN GOOS=linux GOARCH=amd64 go build -o ${image_name} main.go

FROM ${base_image}

RUN apk add --no-cache ca-certificates

ARG image_name
COPY --from=builder /app/${image_name} /usr/local/bin/${image_name}

CMD ["/usr/local/bin/${image_name}"]
