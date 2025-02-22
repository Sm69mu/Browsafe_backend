FROM golang:1.24.0-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod tidy

# Copy .env file first and set proper permissions
COPY .env firebase.json ./
RUN chmod 644 .env firebase.json

# Then copy everything else
COPY . .

EXPOSE 3000

# Verify .env exists before running
RUN ls -la .env

CMD ["air", "-c", ".air.toml"]