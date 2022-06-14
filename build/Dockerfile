FROM golang:1.18-alpine

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
RUN #go mod download

COPY /cmd/test.txt ./
#COPY *.go ./

#RUN apk add --no-cache ca-certificates
RUN echo $PWD
RUN ls -l
RUN #go build cmd/main.go
#WORKDIR /app
#COPY /cmd/service .
#EXPOSE 8080
#ENTRYPOINT [ "/app/service" ]
