FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

RUN go build -o api

EXPOSE 3200
ENTRYPOINT ["/var/www/api"]