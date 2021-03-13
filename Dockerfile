FROM golang:latest
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/Robinrrr10/ordermanagement
RUN go mod init
RUN go mod tidy
ENV server.port=8080 db.mysql.host=localhost db.mysql.port=3306 db.mysql.dbname=business db.mysql.user=root db.mysql.password=root
EXPOSE 8080
CMD ["go", "run", "main.go"]
