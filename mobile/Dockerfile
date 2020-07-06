FROM golang:latest

RUN apt update

RUN go get github.com/stretchr/objx/...
RUN go get github.com/clbanning/x2j/
RUN go get github.com/ugorji/go/codec
RUN go get github.com/stretchr/signature
RUN go get github.com/stretchr/gomniauth/
RUN go get gopkg.in/mgo.v2/bson/

COPY . /go/

RUN go fmt *.go
RUN go build -o gps
CMD ./gps
