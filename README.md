English | [简体中文](https://github.com/niubir/logger/blob/main/helper/README-cn.md)

## Usage

### Start using it

1. Download logger for Go by using

```sh
go get -u github.com/niubir/logger
```

2. Constructing a logger object

```go
import "github.com/niubir/logger"

logger := logger.NewLogger()
```

3. Use
   
```go
logger.Debug("This debug")
logger.Info("This Info")
logger.Warn("This Warn")
logger.Error("This Error")
```

## Configuration option

| Option | Default | Description |
| - | - | - |
| WithLevel | Info | log level(Debug Info Warn Error) |
| WithTimeFormat | - | log with time(Use golang time layout) |
| WithStdout | false | log with os.Stdout |
| WithPath | ./ | log file path |
| WithPrefix | - | log file prefix |
| WithDuration | 365 day | log file duration |
| WithMaxByte | 1024 | log file max byte |
| WithStack | false | log stack |
