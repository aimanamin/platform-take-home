FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make gcc musl-dev
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/server/main.go

FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache postgresql-client
COPY --from=builder /app/main /app/main

EXPOSE 8080 9008 8081

CMD ["/app/main"]