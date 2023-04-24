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
logger.Debug("This Debug")
logger.Info("This Info")
logger.Warning("This Warning")
logger.Error("This Error")
```

## Configuration option

| Option | Default | Description |
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