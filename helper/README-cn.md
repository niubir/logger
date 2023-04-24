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
logger.Debug("This Debug")
logger.Info("This Info")
logger.Warning("This Warning")
logger.Error("This Error")
```

## 配置选项

| 选项 | 默认值 | 描述 |
| - | - | - |
| WithTime | - | log with time(Use golang time layout) |
| SetLevel | Info | log level(Debug Info Warning Error Fatal) |
| WithStack | false | log stack |
| WithStdout | false | log with os.Stdout |
| WithFileout | false | log with file |
| WithFileout.WithFilePath | ./ | log file path |
| WithFileout.WithFilePrefix | - | log file prefix |
| WithFileout.WithFileDuration | 365 day | log file duration |
| WithFileout.WithFileMaxByte | 1024 | log file max byte |