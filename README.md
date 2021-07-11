# gqlgen-demo

Go で GraphQL サーバをたてるためのテンプレートリポジトリ

下記のパッケージを使ってます

- [gqlgen](https://gqlgen.com/)
- [gorm](https://gorm.io/)
- [godotenv](https://github.com/joho/godotenv)

### 開発環境立ち上げ

```shell
git clone https://github.com/kmtym1998/gqlgen-demo
docker compose up
```

### 開発方針

**① スキーマファイル(`schema.graphqls`)を修正**

**② スキーマファイルをもとにコード生成**

```shell
make generate
```

**③ API を書く**

`graph/generated`, `graph/model`, `schema.resolvers.go` にコードが生成される
`graph/generated`, `graph/model`は修正しちゃだめ
`schema.resolvers.go`を修正していく

### 本番環境

```shell
make build-prod
docker run -p 8080:8080 --name gqlgen-demo-prod -it gqlgen-demo-prod
```
