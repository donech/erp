# erp

一个为中小企业设计的精简版 ERP 系统.

## 使用指南

    1. 快速启动 `make run`
    1. 新增自命令 `cobra add SUBCOMMAND`, 安装 cobra `go get -u github.com/spf13/cobra/cobra`

## layout 描述

    1. bin 系统命令输出目录
    1. cmd 系统各命令目录
    1. frontend erp 系统前端组建库
    1. hack 系统脚本目录, bash 或者 jenkinsfile 等等
    1. internal 系统内部目录, 主逻辑目录
        1. common 公共目录, 如 code 码, 及一些内部使用工具
        1. domain 领域层(service 层) 主业务逻辑
        1. entry  入口, grpc-api | gin-api
        1. proto 系统 grpc 定义目录
        1. service grpc service 层
        1. tool 系统内部工具
    1. main.go 仅仅是 go 运行的外部入口, 主要命令实现在 cmd 中
    1. app.yaml 配置文件的样本

## 项目依赖
    1. go 1.15.2+
    1. protoc
    1. github.com/donech/tool 基础工具包
    1. github.com/spf13/cobra
    1. github.com/spf13/viper
