FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /golang-mysql

EXPOSE 8080

CMD [ "/golang-mysql" ]