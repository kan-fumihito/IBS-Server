FROM golang:latest

WORKDIR /go/src/

ADD ./src /go/src/
ADD ./include /usr/include/
ADD ./lib /usr/lib/


RUN go get -u github.com/labstack/echo/...
RUN go get -u google.golang.org/api/option
RUN go get -u firebase.google.com/go
RUN go get -u github.com/go-sql-driver/mysql