my-project/
├── api/               # API 协议定义（如 Swagger, Protobuf, GraphQL 模式）
├── cmd/               # 项目入口（每个子目录对应一个可执行程序）
│   └── main-app/
│       └── main.go
├── internal/          # 私有代码（不允许外部项目引用，业务逻辑核心）
│   ├── logic/
│   ├── model/         # 数据模型定义
│   └── service/
├── pkg/               # 公共库（可被外部项目引用的工具类代码）
├── docs/              # 重点：文档与设计图存放处
│   ├── diagrams/      # ER 图、OML 图、架构图等
│   │   ├── er-diagram.drawio
│   │   └── domain-model.png
│   ├── architecture.md
│   └── api-spec.md
├── scripts/           # 脚本（如数据库迁移脚本、构建脚本）
├── test/              # 额外的外部测试和测试数据
├── go.mod             # 模块依赖
└── README.md