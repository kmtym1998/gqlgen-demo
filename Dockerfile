# https://zenn.dev/iris_ohyaman/articles/73d4472a3e2f94c55664
FROM golang:1.16.5

ENV APP_ROOT /app

RUN mkdir $APP_ROOT
WORKDIR $APP_ROOT

COPY ./ $APP_ROOT

EXPOSE 8081

CMD ["go", "run", "server.go"]