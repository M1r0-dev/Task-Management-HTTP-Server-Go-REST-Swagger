FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./
RUN go mod download

COPY ../../ .

WORKDIR /app/cmd/app

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --generalInfo main.go --output ../../docs

RUN go build -o /app/server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/cmd/app/config.yaml .

EXPOSE 8080

CMD ["./server", "-config=config.yaml"]