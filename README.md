# golog
A simple golang log lib using zap and lumberjack, support log to file and file rotate.

## Quick Start
```golang
logger, err := NewLoggerHandler("1.log", nil)
if err != nil {
    fmt.Println(err)
    return
}
logger.SetLevel(InfoLevel)

logger.Debug("1111", "name", "zy", "age", 10)
logger.Info("2222", "k", "v")
logger.Error("4444")
```