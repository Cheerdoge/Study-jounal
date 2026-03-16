# MySQL 常用命令速查表

## 1. 连接与帮助
```sql
-- 连接数据库
mysql -u [用户名] -p
mysql -u [用户名] -p [数据库名]

-- 退出
exit;

-- 查看版本
SELECT VERSION();

-- 查看帮助
help contents;
```

## 2. 用户管理
```sql
-- 创建用户
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';

-- 授权
GRANT ALL PRIVILEGES ON *.* TO 'username'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;

-- 查看用户权限
SHOW GRANTS FOR 'username'@'localhost';

-- 修改密码 (MySQL 8.0+)
ALTER USER 'username'@'localhost' IDENTIFIED BY 'new_password';

-- 删除用户
DROP USER 'username'@'localhost';
```

## 3. 数据库操作
```sql
-- 查看所有数据库
SHOW DATABASES;

-- 创建数据库
CREATE DATABASE [IF NOT EXISTS] db_name CHARACTER SET utf8mb4;

-- 使用数据库
USE db_name;

-- 查看当前使用的数据库
SELECT DATABASE();

-- 删除数据库
DROP DATABASE [IF EXISTS] db_name;
```

## 4. 表操作
```sql
-- 查看所有表
SHOW TABLES;

-- 创建表
CREATE TABLE table_name (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 查看表结构
DESC table_name;
-- 或
SHOW CREATE TABLE table_name;

-- 修改表名
ALTER TABLE old_name RENAME TO new_name;

-- 添加字段
ALTER TABLE table_name ADD column_name datatype;

-- 修改字段
ALTER TABLE table_name MODIFY column_name new_datatype;

-- 删除字段
ALTER TABLE table_name DROP COLUMN column_name;

-- 删除表
DROP TABLE [IF EXISTS] table_name;
```

## 5. 数据操作 (CURD)
```sql
-- 插入数据
INSERT INTO table_name (col1, col2) VALUES (val1, val2);

-- 查询数据
SELECT * FROM table_name;
SELECT col1, col2 FROM table_name WHERE condition;

-- 更新数据
UPDATE table_name SET col1 = val1 WHERE condition;

-- 删除数据
DELETE FROM table_name WHERE condition;

-- 清空表数据 (重置自增 ID)
TRUNCATE TABLE table_name;
```

## 6. 查询进阶
```sql
-- 排序
SELECT * FROM table_name ORDER BY col_name ASC|DESC;

-- 限制结果
SELECT * FROM table_name LIMIT offset, count;

-- 聚合函数
SELECT COUNT(*), AVG(col), SUM(col), MAX(col), MIN(col) FROM table_name;

-- 分组查询
SELECT col, COUNT(*) FROM table_name GROUP BY col HAVING COUNT(*) > 1;

-- 连接查询 (JOIN)
SELECT * FROM t1 INNER JOIN t2 ON t1.id = t2.t1_id;
SELECT * FROM t1 LEFT JOIN t2 ON t1.id = t2.t1_id;
```

## 7. 索引与约束
```sql
-- 创建索引
CREATE INDEX idx_name ON table_name (column_name);

-- 创建唯一索引
CREATE UNIQUE INDEX idx_name ON table_name (column_name);

-- 查看索引
SHOW INDEX FROM table_name;

-- 删除索引
DROP INDEX idx_name ON table_name;
-- 或
ALTER TABLE table_name DROP INDEX idx_name;
```

## 8. 备份与恢复 (命令行执行)
```bash
# 备份数据库
mysqldump -u [username] -p [db_name] > backup.sql

# 恢复数据库
mysql -u [username] -p [db_name] < backup.sql
```

---
*注：`[]` 内的内容请根据实际情况替换。*