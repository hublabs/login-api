FROM golang:1.13 AS builder

RUN go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/github.com/hublabs/login-api
ADD go.mod go.sum ./
RUN go mod download
ADD . /go/src/github.com/hublabs/login-api
ENV CGO_ENABLED=0
RUN go build -o login-api

FROM pangpanglabs/alpine-ssl
WORKDIR /go/src/github.com/hublabs/login-api
COPY --from=builder /go/src/github.com/hublabs/login-api ./

EXPOSE 8002

CMD ["./login-api", "api-server"]