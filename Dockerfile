FROM golang:1.9.2
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get github.com/Robinrrr10/ordermanagement
RUN go install
ENV server.port=7879 db.mysql.host=192.168.222.62 db.mysql.port=3306 db.mysql.dbname=business db.mysql.user=apper db.mysql.password=app123
EXPOSE 7879
ENTRYPOINT ["/go/src/app"]