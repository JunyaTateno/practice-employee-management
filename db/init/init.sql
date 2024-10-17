-- データベースの作成
CREATE DATABASE IF NOT EXISTS company;

-- 使用するデータベースを選択
USE company;

-- ユーザーの作成と権限の付与
-- TODO: userNameとpassは別途設定する
CREATE USER 'userName'@'%' IDENTIFIED BY 'pass';         
GRANT ALL PRIVILEGES ON company.* TO 'userName'@'%';
FLUSH PRIVILEGES;
