# Use the official Go image as the base image
FROM golang:1.22.2-alpine as builder

# Set the working directory to /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -v -o blog-api ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/blog-api /app/blog-api
COPY --from=builder /app/docs /app/docs

CMD ["./blog-api"]