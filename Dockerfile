FROM golang:1.14
RUN mkdir /app
RUN apt-get install curl
COPY objstore/ /app/objstore
WORKDIR /app/objstore
ENV GOPROXY="https://goproxy.cn,direct"
ENV GO111MODULE="on"
RUN go build -o main main.go
ENTRYPOINT ["/app/objstore/main"]