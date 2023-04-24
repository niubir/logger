package logger

type GRPCLogger struct {
	*Logger
}

// level info:0,warning:1,error:2,fatal:3
func (gl *GRPCLogger) V(level int) bool {
	grpcLogLevel := -1
	switch gl.option.logLevel {
	case DebugLevel:
		grpcLogLevel = -1
	case InfoLevel:
		grpcLogLevel = 0
	case WarningLevel:
		grpcLogLevel = 1
	case ErrorLevel:
		grpcLogLevel = 2
	case FatalLevel:
		grpcLogLevel = 3
	}
	return level <= grpcLogLevel
}
