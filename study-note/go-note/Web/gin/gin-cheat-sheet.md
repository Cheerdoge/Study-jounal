# Gin Web Framework 速查表

## 1. 安装与初始化

```bash
# 安装 gin
go get -u github.com/gin-gonic/gin
```

```go
import "github.com/gin-gonic/gin"

func main() {
    // 创建一个默认的路由引擎，自带 Logger 和 Recovery 中间件
    r := gin.Default()
    
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    // 启动服务，默认监听 8080 端口
    r.Run(":8080")
}
```

## 2. 路由 (Routing)

### 基本路由
```go
r.GET("/someGet", getting)
r.POST("/somePost", posting)
r.PUT("/somePut", putting)
r.DELETE("/someDelete", deleting)
r.PATCH("/somePatch", patching)
r.HEAD("/someHead", head)
r.OPTIONS("/someOptions", options)
```

### 路径参数 (Path Parameters)
```go
// 匹配 /user/john
r.GET("/user/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(200, "Hello %s", name)
})

// 匹配 /user/john/send (星号通配)
r.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    c.String(200, "%s is %s", name, action)
})
```

### 路由组 (Route Groups)
```go
v1 := r.Group("/v1")
{
    v1.POST("/login", loginEndpoint)
    v1.POST("/submit", submitEndpoint)
}

v2 := r.Group("/v2")
{
    v2.POST("/login", loginEndpoint)
}
```

## 3. 获取请求参数

### Query 参数 (`?id=123&name=manu`)
```go
r.GET("/welcome", func(c *gin.Context) {
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的简写
    c.String(200, "Hello %s %s", firstname, lastname)
})
```

### 表单参数 (Form Post)
```go
r.POST("/form_post", func(c *gin.Context) {
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")
    c.JSON(200, gin.H{
        "status":  "posted",
        "message": message,
        "nick":    nick,
    })
})
```

## 4. 模型绑定与验证 (Binding)

Gin 使用 `binding` 标签进行自动校验（底层基于 `validator` 库）。

```go
type Login struct {
    User     string `form:"user" json:"user" xml:"user"  binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

r.POST("/loginJSON", func(c *gin.Context) {
    var json Login
    // ShouldBindJSON 会根据 Content-Type 选择解析器
    if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    if json.User != "manu" || json.Password != "123" {
        c.JSON(401, gin.H{"status": "unauthorized"})
        return
    }
    
    c.JSON(200, gin.H{"status": "you are logged in"})
})
```

## 5. 响应渲染 (Rendering)

```go
// JSON
c.JSON(200, gin.H{"message": "hey", "status": 200})

// XML
c.XML(200, gin.H{"message": "hey", "status": 200})

// YAML
c.YAML(200, gin.H{"message": "hey", "status": 200})

// String
c.String(200, "Hello World")

// 重定向
c.Redirect(301, "https://www.google.com/")
```

## 6. 中间件 (Middleware)

### 使用中间件
```go
r := gin.New() // 不带默认中间件

// 全局中间件
r.Use(gin.Logger())
r.Use(gin.Recovery())

// 给特定路由组使用中间件
authorized := r.Group("/", AuthRequired())
{
    authorized.POST("/read", readHandler)
}
```

### 自定义中间件
```go
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        
        // 设置变量到上下文
        c.Set("example", "12345")
        
        // 请求前逻辑
        c.Next() // 执行后续的处理程序
        
        // 请求后逻辑
        latency := time.Since(t)
        fmt.Print(latency)
    }
}
```

## 7. 文件上传

### 单文件
```go
r.POST("/upload", func(c *gin.Context) {
    file, _ := c.FormFile("file")
    log.Println(file.Filename)
    
    // 上传文件至指定目录
    c.SaveUploadedFile(file, dst)
    c.String(200, fmt.Sprintf("'%s' uploaded!", file.Filename))
})
```

## 8. 静态资源服务
```go
r.Static("/assets", "./assets")
r.StaticFS("/more_static", http.Dir("my_file_system"))
r.StaticFile("/favicon.ico", "./resources/favicon.ico")
```