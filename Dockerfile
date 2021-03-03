FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY . /app/
RUN go build
CMD ["/app/ordermanagement"]