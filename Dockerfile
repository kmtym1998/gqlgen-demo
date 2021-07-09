# https://qiita.com/Yuuki557/items/3d088de91ab86bc71600
FROM golang:1.16.5 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/server \
    -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/bin/server /app/server

EXPOSE 8081

ENTRYPOINT ["/app/server"]