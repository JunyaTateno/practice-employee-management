-- データベースの作成
CREATE DATABASE IF NOT EXISTS company;

-- 使用するデータベースを選択
USE company;

-- ユーザーの作成と権限の付与
-- TODO: youruserとyourpasswordは別途設定する
DROP USER IF EXISTS 'youruser'@'%';
CREATE USER 'youruser'@'%' IDENTIFIED BY 'yourpassword';         
GRANT ALL PRIVILEGES ON company.* TO 'youruser'@'%';
FLUSH PRIVILEGES;
