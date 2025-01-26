# Go の公式イメージを使用
FROM golang:1.23.3

# 作業ディレクトリの設定
WORKDIR /app

# モジュールファイルをコピーして依存関係をインストール
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションコードをコピー
COPY . .

# curl をインストール
RUN apt-get update && apt-get install -y curl

# アプリケーションのビルド
RUN go build -o main.exe .

# ポートを明示
EXPOSE 8080

# アプリケーションの実行
CMD ["./main.exe"]
