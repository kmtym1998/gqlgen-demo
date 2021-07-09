# https://qiita.com/Yuuki557/items/3d088de91ab86bc71600
# https://qiita.com/po3rin/items/8b57e6c22f2b34751333

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
FROM scratch as prod
COPY --from=builder /go/bin/server /app/server
COPY ./.env /
EXPOSE 8081
ENTRYPOINT ["/app/server"]

# 開発環境ホットリロード
FROM golang:1.16.5 as dev
EXPOSE 8081
WORKDIR /app
ENV GO111MODULE=on
COPY . ./
RUN go get github.com/pilu/fresh
CMD ["fresh"]d