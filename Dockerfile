FROM golang:1.17-alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

EXPOSE 8000:8000

WORKDIR /app

COPY ./httpserver ./httpserver
COPY ./go.mod  .
RUN go build -o ./main ./httpserver

CMD ["./main"]