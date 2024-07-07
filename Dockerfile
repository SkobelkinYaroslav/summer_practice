FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/summer_practice

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/app /app/bin/app
EXPOSE 8080

CMD ["/app/bin/app"]
