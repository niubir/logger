[English](https://github.com/niubir/logger/blob/main/helper/README-en.md) | 简体中文

## 用法

### 开始使用

1. 下载

```sh
go get -u github.com/niubir/logger
```

2. 构造logger对象

```go
import "github.com/niubir/logger"

logger := logger.NewLogger()
```

3. 使用
   
```go
logger.Debug("This debug")
logger.Info("This Info")
logger.Warn("This Warn")
logger.Error("This Error")
```

## 配置选项

| 选项 | 默认值 | 描述 |
| - | - | - |
| WithLevel | Info | log level(Debug Info Warn Error) |
| WithTimeFormat | - | log with time(Use golang time layout) |
| WithStdout | false | log with os.Stdout |
| WithPath | ./ | log file path |
| WithPrefix | - | log file prefix |
| WithDuration | 365 day | log file duration |
| WithMaxByte | 1024 | log file max byte |
| WithStack | false | log stack |
