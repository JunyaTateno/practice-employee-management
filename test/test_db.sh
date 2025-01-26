#!/bin/bash 

# エラーハンドリングのための関数
handle_error() {
  local cmd="$1"
  local error_message="$2"
  echo "Error occurred in: $cmd"
  echo "Error message: $error_message"

  # MySQLログを確認
  echo "Checking MySQL logs..."
  docker-compose $COMPOSE_FILE_PATH logs db

  # コンテナを終了してからスクリプトを終了
  docker-compose $COMPOSE_FILE_PATH down
  exit 1
}

# docker-compose.yml のパスを変数に保存 (/app/ 内にある場合)
COMPOSE_FILE_PATH="-f ./app/docker-compose.yml"

# MySQL接続情報
DB_HOST="db"
DB_PORT="3306"
DB_USER="youruser"
DB_PASS="yourpassword"
DB_NAME="company"

# MySQLが準備完了になるまで待機（10回まで試行）
echo "Waiting for MySQL to be ready..."
max_attempts=5  # 最大試行回数
attempt=0  # 試行回数カウント

until docker-compose $COMPOSE_FILE_PATH exec -T db mysql --protocol=tcp -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS -e "SELECT 1" > /dev/null; do
  attempt=$((attempt + 1))  # 試行回数をインクリメント
  if [ $attempt -ge $max_attempts ]; then
    handle_error "Waiting for MySQL readiness" "MySQL did not become ready after $max_attempts attempts"
  fi
  echo "Waiting for MySQL to be ready (attempt $attempt of $max_attempts)..."
  sleep 10
done
echo "MySQL is ready!"

# insert_test.sqlを実行してデータを挿入
echo "Inserting test data from insert_test.sql..."
error_message=$(docker-compose $COMPOSE_FILE_PATH exec -T db mysql -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS $DB_NAME < ./test/insert_test.sql 2>&1)
if [ $? -ne 0 ];then handle_error "Inserting data from insert_test.sql" "$error_message"; fi
echo "Test data inserted successfully."

# select_test.sqlを実行してデータを確認
echo "Selecting test data from select_test.sql..."
error_message=$(docker-compose $COMPOSE_FILE_PATH exec -T db mysql -h$DB_HOST -P$DB_PORT -u$DB_USER -p$DB_PASS $DB_NAME < ./test/select_test.sql 2>&1)
if [ $? -ne 0 ]; then handle_error "Selecting data from select_test.sql" "$error_message"; fi
echo "Test data selected successfully."

# すべての処理が成功
echo "All tests completed successfully!!"
