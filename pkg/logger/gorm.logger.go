package logger

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogger struct {
	LogLevel                  logger.LogLevel
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration
}

func (gl *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *gl
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (gl GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= logger.Info {
		Log.Infof(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (gl GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= logger.Warn {
		Log.Warnf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (gl GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= logger.Error {
		Log.Errorf(msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

var (
	traceStr     = "[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s\n[%.3fms] [rows:%v] %s"
)

// Trace print sql message
func (gl GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && gl.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !gl.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			Log.Errorf(traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Log.Errorf(traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > gl.SlowThreshold && gl.SlowThreshold != 0 && gl.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", gl.SlowThreshold)
		if rows == -1 {
			Log.Warnf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Log.Warnf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case gl.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			Log.Debugf(traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Log.Debugf(traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func CreateGormLogger(logLevel string) (logger.Interface, error) {

	var gormLogLevel logger.LogLevel
	switch strings.ToLower(logLevel) {
	case "silent":
		gormLogLevel = logger.Silent
	case "debug", "info":
		gormLogLevel = logger.Info
	case "error":
		gormLogLevel = logger.Error
	case "warn":
		gormLogLevel = logger.Warn
	default:
		return nil, fmt.Errorf("Invalid log level: %s", logLevel)
	}

	return &GormLogger{
		LogLevel:                  gormLogLevel,
		IgnoreRecordNotFoundError: false,
		SlowThreshold:             200 * time.Millisecond, // TODO take value from env
	}, nil

}
