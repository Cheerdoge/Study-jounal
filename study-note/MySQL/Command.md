# 启动和关闭
## Windows
`net start mysql95`
`net stop mysql95`

## Linux
`sudo systemctl start mysql`
`sudo systemctl stop mysql`
重启:`sudo systemctl restart mysql`
检查:`sudo systemctl status mysql`

# 登录
`mysql -h 主机名 -u 用户名 -p`
- **-h** : 指定客户端所要登录的 MySQL 主机名, 登录本机(localhost 或 127.0.0.1)该参数可以省略;
- **-u** : 登录的用户名;
- **-p** : 告诉服务器将会使用一个密码来登录, 如果所要登录的用户名密码为空, 可以忽略此选项
	列出所有可用的数据库：
	
	`SHOW DATABASES`;
	
	选择要使用的数据库：
	
	`USE your_database`;
	
	列出所选数据库中的所有表：
	
	`SHOW TABLES`;


# 用户设置

1. 创建用户
	`CREAT USER 'username'@'host' IDENTIFIED BY 'password';`
		'host'指定可以从哪些主机连接，如`'localhost'`为本地连接，`%`允许任何主机连接
	
2. 授权权限
	`GRANT privileges ON database_name.* TO 'username'@'host';`
	1. `privileges`：所需的权限，如 `ALL PRIVILEGES`、`SELECT`、`INSERT`、`UPDATE`、`DELETE` 等
	
	2. `database_name.*`：表示对某个数据库或表授予权限。`database_name.*` 表示对整个数据库的所有表授予权限，`database_name.table_name` 表示对指定的表授予权限
	
3. 刷新权限
	`FLUSH PRIVILEGES;`
	
4. 查看用户权限
	`SHOW GRANTS FOR 'username'@'host';`
	
5. 撤销权限
	`REVOKE privileges ON database_name.* FROM 'username'@'host';`
	
6. 删除用户
	`DROP USER 'username'@'host';`
	
7. 修改用户密码
	`ALTER USER 'username'@'host' IDENTIFIED BY 'new_password';`
	
8. 修改用户主机
	删除原有用户并创建新用户，创建用户时可以同时授予权限

# 管理MySQL
1. **USE _数据库名_**
	选择要操作的Mysql数据库，使用该命令后所有Mysql命令都只针对该数据库
	
2. **SHOW DATABASES**
	列出 MySQL 数据库管理系统的数据库列表
	
3. **SHOW TABLES**
	显示指定数据库的所有表，使用该命令前需要使用 use 命令来选择要操作的数据库
4. **SHOW COLUMNS FROM _数据表名称_**
	显示数据表的属性，属性类型，主键信息 ，是否为 NULL，默认值等其他信息
	
5. **SHOW INDEX FROM 数据表名称**
	显示数据表的详细索引信息，包括PRIMARY KEY（主键）
	
6. **SHOW TABLE STATUS [FROM databasename]   [LIKE 'pattern'] \G:**
	该命令将输出Mysql数据库管理系统的性能及统计信息
	`SHOW TABLE STATUS  FROM RUNOOB;   # 显示数据库 RUNOOB 中所有表的信息`
	`SHOW TABLE STATUS from RUNOOB LIKE 'runoob%';     # 表名以runoob开头的表的信息`
	`SHOW TABLE STATUS from RUNOOB LIKE 'runoob%'\G;   # 加上 \G，查询结果按列打印`

# MySQL语法

## 创建数据库
1. **CREATE**
	`CREATE DATABASE IF NOT EXISTS 数据库名`
	`CHARACTER SET`和`COLLATE`指定字符集和排序规则
2. **mysqladmin**
	`mysqladmin -u username -p create your_database`
	- `-u` 参数用于指定 MySQL 用户名。
	- `-p` 参数表示需要输入密码。
	- `create` 是执行的操作，表示创建数据库。
	- `your_database` 是要创建的数据库的名称。
	- `-default-character-set`和 `-default-collation`指定字符集和排序规则
	普通用户需要特定的权限来创建或者删除 MySQL 数据库，root 用户拥有最高权限
	
	