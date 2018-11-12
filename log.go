package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "strings"
    "runtime"
    "errors"
)

const (
    DebugLevel = iota
    InfoLevel
    WarnLevel
    ErrorLevel
)

type Level int

type Logger struct {
    logger      *zap.SugaredLogger
    atomLevel   zap.AtomicLevel
}

type RotateConf struct {
    MaxSize     int     //MB
    MaxBackups  int
    MaxAge      int     //Day
}

var (
    levelMap = map[int]zapcore.Level {
        DebugLevel: zapcore.DebugLevel,
        InfoLevel:  zapcore.InfoLevel,
        WarnLevel:  zapcore.WarnLevel,
        ErrorLevel: zapcore.ErrorLevel,
    }

    defaultRotateConf = &RotateConf{
        MaxSize:    100,
        MaxBackups: 20,
        MaxAge:     7,
    }
)


func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

func NewLoggerHandler(fileName string, rotateConf *RotateConf) (*Logger, error) {
    if fileName == "" {
        return nil, errors.New("fileName is empty")
    }

    if rotateConf == nil {
        rotateConf = defaultRotateConf
    }

    if rotateConf.MaxSize <= 0 {
        rotateConf.MaxSize = 100
    }
    if rotateConf.MaxBackups <= 0 {
        rotateConf.MaxBackups = 20
    }
    if rotateConf.MaxAge <= 0 {
        rotateConf.MaxAge = 7
    }

    cfg := zap.NewProductionEncoderConfig()
    cfg.EncodeTime = zapcore.ISO8601TimeEncoder
    cfg.EncodeCaller = zapcore.ShortCallerEncoder

    atomLevel := zap.NewAtomicLevel()

    w := zapcore.AddSync(&lumberjack.Logger{
        Filename:   fileName,
        MaxSize:    rotateConf.MaxSize,
        MaxBackups: rotateConf.MaxBackups,
        MaxAge:     rotateConf.MaxAge,
        LocalTime:  true,
        Compress:   true,
    })
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(cfg),
        w,
        atomLevel,
    )
    sugarLogger := zap.New(core, zap.AddCaller()).Sugar()
    atomLevel.SetLevel(zapcore.DebugLevel)

    logger := &Logger{
        logger:     sugarLogger,
        atomLevel:  atomLevel,
    }

    return logger, nil
}

func (logger *Logger) SetLevel(l Level) {
    level, ok := levelMap[int(l)]
    if !ok {
        level = DebugLevel
    }

    logger.atomLevel.SetLevel(level)
}

func (logger *Logger) Debug(msg string, keysAndValues ...interface{}) {
    logger.logger.Debugw(msg, keysAndValues...)
}

func (logger *Logger) Info(msg string, keysAndValues ...interface{}) {
    logger.logger.Infow(msg, keysAndValues...)
}

func (logger *Logger) Warn(msg string, keysAndValues ...interface{}) {
    logger.logger.Warnw(msg, keysAndValues...)
}

func (logger *Logger) Error(msg string, keysAndValues ...interface{}) {
    logger.logger.Errorw(msg, keysAndValues...)
}

func (logger *Logger) Debugf(template string, args ...interface{}) {
    logger.logger.Debugf(template, args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
    logger.logger.Infof(template, args...)
}

func (logger *Logger) Warnf(template string, args ...interface{}) {
    logger.logger.Warnf(template, args...)
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
    logger.logger.Errorf(template, args...)
}