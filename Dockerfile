FROM golang:1.24.0-alpine
WORKDIR /app
COPY go.* ./
RUN go mod tidy
COPY . .
RUN go build -o main cmd/main.go
EXPOSE 8080
CMD [ "./main" ]