WORKDIR /app

ADD *

RUN go build main.go 

CMD main 