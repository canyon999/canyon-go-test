FROM golang:1.14 as builder

WORKDIR /build
COPY . /build


RUN go build -a -o basic

# 表示依赖 alpine 最新版
FROM alpine:latest

WORKDIR /data

COPY --from=builder /build/basic .

EXPOSE 8080
ENTRYPOINT ["./basic"]