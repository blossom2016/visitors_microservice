#compile stage
FROM golang:1.14.2 AS build-env

#Enable the vendoring mode
ENV GOFLAGS="-mod=vendor"

ADD . /dockerdev
WORKDIR /dockerdev

RUN go build -o /server


#Final stage
FROM debian:buster

EXPOSE 8080

WORKDIR /
COPY --from=build-env /server /

CMD ["/server"]