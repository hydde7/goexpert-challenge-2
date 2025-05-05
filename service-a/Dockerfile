FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init || true
RUN go build -o service-a ./main.go
EXPOSE 8080
CMD ["./service-a"]