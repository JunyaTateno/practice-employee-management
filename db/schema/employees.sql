-- employees テーブルの作成
CREATE TABLE IF NOT EXISTS employees (
    id          INT(6)        UNSIGNED ZEROFILL AUTO_INCREMENT PRIMARY KEY,           -- 社員ID
    family_name VARCHAR(30)   NOT NULL,                                               -- 名字
    first_name  VARCHAR(30)   NOT NULL,                                               -- 氏名
    position    VARCHAR(10)   ,                                                       -- 役職
    department  VARCHAR(10)   ,                                                       -- 部署
    updated_at  TIMESTAMP     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新日時
    created_at  TIMESTAMP     DEFAULT CURRENT_TIMESTAMP                               -- 作成日時
) CHARACTER SET utf8mb4;
