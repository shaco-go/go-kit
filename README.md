# go 工具

使用

```go
go get -u github.com/shaco-go/go -kit
```

# log

根据 [zap](https://github.com/uber-go/zap) 日志库实现，支持多种日志输出渠道(文件分割,可读的控制台输出,飞书,server3酱)

## 使用方法

### 默认

默认的通道为`console` `file`

```go
logger := log.Default().Zap()
```

### 根据yaml配置文件实例化

yaml配置文件

```yaml
# 日志配置
log:
  name: "app"
  debug: false
  level: "info"  # 日志级别["debug","info","warn","error","dpanic","panic","fatal"]
  channel: # 日志通道["file","console","lark","server3"]
    - "file"
    - "console"
  console: # 控制台日志配置
    level: "debug"  # 控制台日志级别,不填写默认log的级别
  file: # 文件日志配置
    level: "warn"  # Lark 日志级别,不填写默认log的级别
    filename: "app.log"  # 日志文件名
    maxsize: 100         # 日志文件最大大小（MB）
    maxage: 30           # 日志文件保留最大天数
    maxbackups: 5        # 最大备份日志文件数
    localtime: true      # 是否使用本地时间
    compress: false      # 是否压缩日志文件
  lark: # Lark 通知配置
    level: "error"  # Lark 日志级别,不填写默认log的级别
    webhook: ""  # Lark Webhook 地址
    detailed: true   # 是否发送详细日志
  server3: # Server3酱 通知配置
    level: "error"  # Server3酱 日志级别
    sendkey: ""  # Server3酱 发送密钥
    detailed: false  # 是否发送详细日志
```