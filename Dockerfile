# Compile stage
FROM golang:1.16.6 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /server

# Final stage
FROM debian:buster

EXPOSE 8080

WORKDIR /
#COPY .env .env
COPY --from=build-env /server /

CMD ["/server"]