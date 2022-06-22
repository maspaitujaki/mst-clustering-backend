FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o mst-api

EXPOSE 9000

CMD ./mst-api