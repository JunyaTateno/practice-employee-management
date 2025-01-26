#!/bin/bash 

# エラーハンドリングのための関数
handle_error() {
  local cmd="$1"
  local error_message="$2"
  echo "Error occurred in: $cmd"
  echo "Error message: $error_message"

  # Docker Compose のログを確認
  echo "Checking Docker Compose logs for backend..."
  docker-compose $COMPOSE_FILE_PATH logs backend

  exit 1
}

# docker-compose.yml のパスを変数に保存
COMPOSE_FILE_PATH="-f ./app/docker-compose.yml"

# backend コンテナのヘルスチェックを待機（10回まで試行）
echo "Waiting for backend to be healthy..."
max_attempts=10
attempt=0

until [ "$(docker inspect --format='{{.State.Health.Status}}' backend)" == "healthy" ]; do
  attempt=$((attempt + 1))
  if [ $attempt -ge $max_attempts ]; then
    handle_error "Waiting for backend health" "Backend did not become healthy after $max_attempts attempts"
  fi
  echo "Waiting for backend to be healthy (attempt $attempt of $max_attempts)..."
  sleep 5
done
echo "Backend is healthy!"

# backend コンテナ内で Go テストを実行
echo "Running Go tests in backend container..."
error_message=$(docker-compose $COMPOSE_FILE_PATH exec -T backend sh -c "go test ./... -v" 2>&1)
if [ $? -ne 0 ]; then handle_error "Running Go tests" "$error_message"; fi
echo "Go tests completed successfully."
