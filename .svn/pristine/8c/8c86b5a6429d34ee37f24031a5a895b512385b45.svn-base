FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/my_gin
COPY . $GOPATH/src/my_gin
RUN go build .

EXPOSE 9090
ENTRYPOINT ["./my_gin"]
