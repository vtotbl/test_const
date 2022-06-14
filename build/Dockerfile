FROM golang:1.18-alpine

WORKDIR /app
COPY . ./
EXPOSE 8080
RUN go build ./cmd/main.go
CMD ./main