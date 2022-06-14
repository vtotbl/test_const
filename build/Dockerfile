FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY ./cmd/main.go ./
#COPY go.sum ./
#RUN go mod download

RUN go get github.com/vtotbl/test_const.git
RUN go build github.com/vtotbl/test_const.git
#RUN go mod tidy
#RUN go build ./main.go
#COPY *.go ./

#RUN apk add --no-cache ca-certificates
RUN echo $PWD
RUN ls -l
RUN #go build cmd/main.go
#WORKDIR /app
#COPY /cmd/service .
#EXPOSE 8080
#ENTRYPOINT [ "/app/service" ]
