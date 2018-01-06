FROM golang:latest

LABEL maintainer="tforrest"

ADD . /app/
WORKDIR /app

RUN go build -o server .
EXPOSE 443
CMD ["/app/server"]