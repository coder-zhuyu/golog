package log

import "testing"

func TestLogger(t *testing.T) {

    logger, err := NewLoggerHandler("1.log", nil)
    if err != nil {
        t.Fatal(err)
    }
    logger.SetLevel(InfoLevel)

    logger.Debug("1111", "name", "zy", "age", 10)
    logger.Info("2222", "k", "v")
    logger.Error("4444")
}