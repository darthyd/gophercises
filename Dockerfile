FROM golang:1.17

WORKDIR /go/app

COPY . ./go/app
CMD ["tail", "-f", "/dev/null"]