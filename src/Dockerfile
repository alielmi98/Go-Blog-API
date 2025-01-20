FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o blog-api cmd/main.go
CMD ["/app/blog-api"]
