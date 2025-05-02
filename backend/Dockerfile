FROM golang:1.24.2-alpine

# 必要なツールのインストール
RUN apk add --no-cache git bash

# airのインストール（PATH通す）
ENV GO111MODULE=on \
    PATH=$PATH:/go/bin

RUN go install github.com/air-verse/air@latest

# アプリの作業ディレクトリ
WORKDIR /app

# go.mod, go.sum を先にコピー（キャッシュ効かせるため）
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# go mod tidyを実行して依存関係を整理
RUN go mod tidy

# ポート公開（例：8080を使う場合）
EXPOSE 8080

# airでアプリ起動
CMD ["air"]
