#!/bin/bash

# エラーハンドリングの関数
handle_error() {
  local cmd="$1"
  local error_message="$2"
  echo "Error occurred in: $cmd"
  echo "Error message: $error_message"

  # Nginx のログを表示
  echo "Checking Nginx logs..."
  docker-compose logs frontend

  # コンテナを終了してからスクリプトを終了
  docker-compose down
  exit 1
}

# docker-compose.yml のパスを変数に保存
COMPOSE_FILE_PATH="-f ./app/docker-compose.yml"

# フロントエンドのヘルスチェックを待つ
echo "Waiting for frontend service to be healthy..."
max_attempts=10
attempt=0

until [ "$(docker inspect --format='{{.State.Health.Status}}' frontend)" == "healthy" ]; do
  attempt=$((attempt + 1))
  if [ $attempt -ge $max_attempts ]; then
    handle_error "Waiting for frontend readiness" "Frontend did not become healthy after $max_attempts attempts"
  fi
  echo "Waiting for frontend to be healthy (attempt $attempt of $max_attempts)..."
  sleep 5
done
echo "Frontend is healthy!"

# メインページ (index.html) にアクセスできるか確認
echo "Testing main page access..."
status_code=$(curl -o /dev/null -s -w "%{http_code}" http://localhost)
if [ "$status_code" -ne 200 ]; then
  handle_error "Testing main page access" "Expected HTTP 200 but got $status_code"
fi
echo "Main page is accessible (HTTP 200)."

# テスト完了
echo "Frontend test completed successfully!"

