# gqlgen-demo

### 開発環境立ち上げ

- `.env.sample` をコピーして `.env` ファイルを作成

```sh
git clone https://github.com/kmtym1998/gqlgen-demo
docker compose up
```

### 本番環境

```sh
make build-prod
docker run -p 8081:8081 --name gqlgen-demo-prod -it gqlgen-demo-prod
```

- [gqlgen](https://gqlgen.com/)
- [gorm](https://gorm.io/)
- [godotenv](https://github.com/joho/godotenv)
