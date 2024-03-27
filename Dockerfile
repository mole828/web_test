FROM golang:latest

WORKDIR /app

ADD . .

ARG GOPROXY
ENV GOPROXY=$GOPROXY

RUN go build main.go 

CMD ["./main"]