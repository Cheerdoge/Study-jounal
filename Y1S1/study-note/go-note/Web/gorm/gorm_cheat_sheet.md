# GORM 速查表 (Cheat Sheet)

## 1. 安装与连接

### 安装
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql # 对应数据库驱动
```

### 连接 MySQL
```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

---

## 2. 模型定义 (Models)

### 基础模型
```go
type User struct {
  gorm.Model           // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
  Name     string
  Age      int         `gorm:"default:18"`
  Email    *string     `gorm:"uniqueIndex"` // 指针可处理零值
  Birthday *time.Time
}
```

### 常用标签 (Tags)
- `primaryKey`: 主键
- `column:name`: 指定列名
- `type:varchar(100)`: 指定列类型
- `not null`: 不为空
- `index`: 创建索引
- `unique`: 唯一索引
- `-`: 忽略此字段

---

## 3. CRUD 操作

### 创建 (Create)
```go
user := User{Name: "Jinzhu", Age: 18}
result := db.Create(&user) // result.Error, result.RowsAffected

// 批量创建
var users = []User{{Name: "u1"}, {Name: "u2"}}
db.Create(&users)
```

### 查询 (Read)
```go
// 获取第一条记录（按主键排序）
db.First(&user)
// 获取一条记录（没有指定排序）
db.Take(&user)
// 主键查询
db.First(&user, 10)
db.First(&user, "10")
// 获取全部记录
db.Find(&users)

// 条件查询 (Where)
db.Where("name = ?", "jinzhu").First(&user)
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
db.Where("name LIKE ?", "%jin%").Find(&users)
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user) // 结构体查询（不查零值）
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
```

### 更新 (Update)
```go
// 保存所有字段
db.Save(&user)

// 更新单个列
db.Model(&user).Update("name", "hello")

// 更新多个列 (Struct 仅更新非零值)
db.Model(&user).Updates(User{Name: "hello", Age: 18})

// 更新多个列 (Map 更新所有指定值)
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "active": false})
```

### 删除 (Delete)
```go
// 软删除（如果模型包含 DeletedAt）
db.Delete(&user)
// 带条件的批量删除
db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
// 永久删除 (Unscoped)
db.Unscoped().Delete(&user)
```

---

## 4. 高级查询

### 链式操作
```go
db.Select("name", "age").Where("age > ?", 20).Limit(10).Offset(5).Order("age desc").Find(&users)
```

### 聚合函数
```go
var count int64
db.Model(&User{}).Count(&count)
db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
```

### 原生 SQL
```go
db.Raw("SELECT name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)
db.Exec("DROP TABLE users")
```

---

## 5. 关联 (Associations)

### 预加载 (Preload / Eager Loading)
```go
// 查询用户时同时加载其 Orders
db.Preload("Orders").Find(&users)
// 带条件的预加载
db.Preload("Orders", "state = ?", "paid").Find(&users)
```

---

## 6. 事务 (Transactions)

### 自动事务
```go
db.Transaction(func(tx *gorm.DB) error {
  if err := tx.Create(&User{Name: "Giraffe"}).Error; err != nil {
    return err // 返回 err 则回滚
  }
  return nil // 返回 nil 则提交
})
```

### 手动事务
```go
tx := db.Begin()
// ... 各种操作
tx.Rollback()
tx.Commit()
```

---

## 7. 钩子 (Hooks)
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()
  return
}
```

---

## 8. 配置与调试

### 打印日志 (Logger)
```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Info),
})
```

### 连接池设置
```go
sqlDB, _ := db.DB()
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```