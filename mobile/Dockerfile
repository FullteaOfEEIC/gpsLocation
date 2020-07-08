# create certificates
FROM ubuntu:latest AS cert
RUN apt update && apt install -y openssl
RUN openssl genrsa -out dblec.key 3072
RUN openssl req -new -key dblec.key -sha256 -out dblec.csr -subj "/CN=frt.hongo.wide.ad.jp"
RUN openssl x509 -in dblec.csr -days 365 -req -signkey dblec.key -sha256 -out dblec.crt


FROM golang:latest

RUN apt update

RUN go get github.com/stretchr/objx/...
RUN go get github.com/clbanning/x2j/
RUN go get github.com/ugorji/go/codec
RUN go get github.com/stretchr/signature
RUN go get github.com/stretchr/gomniauth/
RUN go get gopkg.in/mgo.v2/bson/
RUN go get gopkg.in/go-playground/validator.v9/
RUN go get github.com/go-sql-driver/mysql

RUN mkdir /ssl
COPY --from=cert dblec.crt /ssl
COPY --from=cert dblec.key /ssl
COPY . /go/

RUN go fmt *.go
RUN go build -o gps
CMD ./gps
