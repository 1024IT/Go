# 2023年3月時点の最新版goの軽量版「alpine」
FROM golang:1.20.2-alpine

# Install git
RUN apk update && apk add --no-cache git

# ホットリロード用にAirをインストール
RUN go install github.com/cosmtrek/air@v1.40.4

# マイグレーション用にglang-migrate（MySQL）をインストール
RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            git \
            gcc \
            musl-dev

# 作業ディレクトリの設定
WORKDIR /go/src