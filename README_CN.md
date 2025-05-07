# Morelogin MCP Server

Morelogin MCP 服务器是一个用于 Morelogin 的模型上下文协议（Model Context Protocol，MCP）服务器实现。它提供了一系列与 Morelogin API 交互的工具，使 AI 助手能够管理仓库、问题、拉取请求等。

## 功能特点

- 与 Morelogin 环境、代理、分组、标签API进行交互
- 可配置的 API 基础 URL
- 命令行标志，便于配置
- 支持个人、组织和企业操作
>

## 安装

### 前提条件

- Go 1.23.0 或更高版本
- Morelogin 账号支持API权限

### 从源代码构建

1. 克隆仓库：
   ```bash
   git clone project
   cd mcp-morelogin
   ```

2. 构建项目：
   ```bash
   make build
   ```
   将 ./bin/mcp-morelogin
   
### 使用 go install 安装
   ```bash
   go install morelogin.com/mcp@latest
   ```

## 使用方法

检查 mcp-morelogin 版本：

```bash
mcp-morelogin --version
```

