FROM golang:latest
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/Robinrrr10/ordermanagement
RUN go mod init
RUN go mod tidy
ENV server.port=7879 db.mysql.host=192.168.222.62 db.mysql.port=3306 db.mysql.dbname=business db.mysql.user=apper db.mysql.password=app123
EXPOSE 7879
CMD ["go", "run", "main.go"]
