# ビルドステージ
FROM golang:1.23.2-alpine3.20 AS builder

## 必要なパッケージのインストール
#RUN apk --no-cache add gcc musl-dev
#
# 作業ディレクトリの設定
WORKDIR /app

# go.mod, go.sum, ソースコードのコピー
# Dockerfileと同じディレクトリから相対パスでファイルをコピーする
COPY go.mod go.sum /app/
COPY cmd/ /app/cmd/
COPY internal/ /app/internal/

# 依存関係のインストール
RUN go mod download

ARG APP_VERSION=dev
ARG CMMIT_HASH=none
# Goアプリケーションのビルド
# go build -ldflags "-X 'main.Version=v1.0.0' -X 'main.Commit=abcdef' -X 'main.BuiltBy=docker'"
RUN go build -ldflags "-X 'main.Version=${APP_VERSION}' -X 'main.Commit=${COMMIT_HASH}' -X 'main.BuiltBy=docker'" -o /app/api /app/cmd/server/main.go

# 実行ステージ
FROM alpine:latest

# 作業ディレクトリの設定
WORKDIR /root/

# ビルドされたバイナリをコピー
COPY --from=builder /app/api .

ARG ENV_NAME=.env.prod
COPY envs/${ENV_NAME} /api/.env

ARG APP_PORT=8080
EXPOSE ${APP_PORT}

# アプリケーションを実行
CMD ["./api", "-e", "/api/.env"]
