geth/
├── cmd/
│   └── main.go   #主函数
├── internal/
│   ├── api/
│   │   ├── handler.go
│   │   └── router.go
│   ├── blockchain/
│   │   ├── client.go          #Client获取blockchain信息
│   │   ├── contract.go      # 由 abigen 生成
│   │   └── event_listener.go  #事件监听
│   ├── model/
│   │   └── models.go   #数据结构
│   ├── repository/
│   │   └── auction_repo.go   #数据访问
│   └── service/
│       └── auction_service.go   #业务逻辑
├── config/
│   └── config.yaml   #数据库&blockchain配置
├── migrations/
│   └── init.sql    #mysql表结构
├── .env
├── go.mod
└── go.sum