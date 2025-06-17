FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

# ランタイム用に軽量イメージを用意
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/server .

ENV PORT=8080
EXPOSE 8080

CMD [ "./server" ]
