FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY . /app/
ENV server.port=7879 db.mysql.host=192.168.222.62 db.mysql.port=3306 db.mysql.dbname=business db.mysql.user=apper db.mysql.password=app123
RUN go build
EXPOSE 7879
CMD ["go", "run", "main.go"]