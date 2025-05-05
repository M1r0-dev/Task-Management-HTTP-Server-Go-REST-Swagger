# ── Стадия сборки ─────────────────────────────────────────────────────────────
FROM golang:1.23.6-alpine AS builder
RUN apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH=/go/bin:$PATH

WORKDIR /proj
COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR /proj

RUN swag init \
    --generalInfo cmd/app/main.go \
    --dir . \                     
    --output docs


RUN mv docs/swagger.json docs/doc.json

# затем собираем бинарник
WORKDIR /proj/cmd/app
RUN go build -o /proj/server .


# ── Финальный образ ───────────────────────────────────────────────────────────
FROM alpine:latest
WORKDIR /proj

COPY --from=builder /proj/server .
COPY --from=builder /proj/cmd/app/config.yaml .
COPY --from=builder /proj/docs ./docs

EXPOSE 8080

CMD ["./server", "-config=config.yaml"]
