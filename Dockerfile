FROM golang:1.20-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
# RUN go mod tidy
RUN go run main.go update
RUN go build -o /app/main main.go

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD ["/app/main"]
