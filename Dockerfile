# 本番環境ではビルドされたファイルを実行するだけの単純なイメージにする

# バイナリファイルをビルドする用の中間イメージ
FROM golang:1.16.5 as builder
WORKDIR /app
COPY .  ./
RUN go mod download
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/server \
    -ldflags '-s -w'

# 本番実行環境
FROM scratch as runner
COPY --from=builder /go/bin/server /app/server
COPY ./.env /
EXPOSE 8081
ENTRYPOINT ["/app/server"]
