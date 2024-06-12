FROM golang:latest

RUN mkdir -p /go/src/karyavideo

WORKDIR /go/src/karyavideo

COPY . .

##RUN go get -d -v ./...
RUN go build -o main .



CMD ["/go/src/karyavideo/main"]

EXPOSE 8083