# go 工具

使用
```go
go get -u github.com/shaco-go/go-kit
```

## logc
```yaml
# 日志配置
log:
  # 环境类型，支持 dev 和 prod
  env: dev  # 可选值: dev, prod

  # 日志级别，支持 debug, info, warn, error, dpanic, panic, fatal
  level: info  # 可选值: debug, info, warn, error, dpanic, panic, fatal

  # 日志输出渠道，可以选择多个
  channel:
    - console  # 控制台输出
    - file     # 文件输出
    - lark     # 飞书
    # - dingtalk  # 钉钉（暂未实现）
    # - wecom     # 企业微信（暂未实现）

  # 控制台日志配置
  console:
    level: debug  # 控制台输出的日志级别

  # 飞书日志配置
  lark:
    webhook: "https://example.com/webhook"  # 飞书的 webhook 地址
    level: info  # 飞书输出的日志级别

  # 文件日志配置
  file:
    filename: "logs/app.log"  # 日志文件的路径
    maxsize: 100               # 日志文件的最大大小（MB）
    maxage: 30                 # 保留旧日志文件的最大天数
    maxbackups: 5              # 保留的旧日志文件的最大数量
    localtime: true            # 是否使用本地时间格式化备份文件中的时间戳
    compress: false            # 是否压缩轮换的日志文件
    level: error               # 文件输出的日志级别

```
> **注意：** 上方配置文件中`env` `channel` `level` 需要使用`Conv**`转换一下

### 使用
```go
    // 根据配置实例化,多个通道就在channel添加,目前支持 控制台,文件,飞书,其他通道可参数具体代码
    logger := logc.New(&logc.Config{
        Env:     logc.Dev,
        Level:   zapcore.DebugLevel,
        Channel: []logc.Channel{logc.ConsoleChannel},
    })
    logger.Error("error msg")
```

## logx 日志增强
> 主要是自带上下文,zap是没有上下文的,方便打印rid,排查全局日志
### 使用
```go
    logger := logc.New(&logc.Config{
        Env:     logc.Dev,
        Level:   zapcore.DebugLevel,
        Channel: []logc.Channel{logc.ConsoleChannel},
    })
    logx := New(logger, func(ctx context.Context, logger *zap.Logger) *zap.Logger {
        // 全局上下文参数的逻辑可以在这里实现
        rid := ctx.Value("rid").(string)
        return logger.With(zap.String("rid", rid))
    })
```

