# Build stage
FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY config.env ./
COPY database.db ./
RUN apk add build-base
RUN go mod download

COPY . .
RUN go build -o main main.go

EXPOSE 8080
CMD [ "/app/main" ]
