# ErgoGo

一个API开发框架，体型小巧功能强大，有完善的基础处理功能，简单易懂，快速上手。

## 运行

```shell
go mod tidy

# 运行
go run main.go serve
# 或者
go run .
```

## 程序结构

```shell
├── bootstrap
├── config
├── go.mod
├── internal
├── main.go
├── pkg
├── README.md
└── router

# bootstrap 程序安装组件目录
# config 配置信息目录
# router 路由
# pkg 公共辅助包
# internal 程序具体逻辑代码
# main.go 程序入口

# tmp air 的工作目录
# .env 环境变量文件
```

## 功能日志

- 2024-05-07：项目启动
