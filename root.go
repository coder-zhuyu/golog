package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

var (
    root = &Logger{
        logger: zap.New(zapcore.NewCore(
            zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
            os.Stdout,
            zap.NewAtomicLevelAt(zapcore.DebugLevel),
        )).Sugar(),
    }
)

func Debug(msg string, keysAndValues ...interface{}) {
    root.logger.Debugw(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...interface{}) {
    root.logger.Infow(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
    root.logger.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
    root.logger.Errorw(msg, keysAndValues...)
}

func Debugf(template string, args ...interface{}) {
    root.logger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
    root.logger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
    root.logger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
    root.logger.Errorf(template, args...)
}
