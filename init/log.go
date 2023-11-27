package init

import (
	"fintechpractices/configs"
	"fintechpractices/global"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/v2pro/plz/countlog/output/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	level   zapcore.Level
	options []zap.Option
)

func InitZapLog() *zap.Logger {
	logCfg := global.AppCfg.LogCfg
	err := tryCreateDir(logCfg.Dir)
	if err != nil {
		fmt.Printf("tryCreateDir error: %s", err.Error())
		return nil
	}

	setLogLevel(logCfg.Level)

	if logCfg.LineNo {
		options = append(options, zap.AddCaller())
	}
	return zap.New(getZapCore(&logCfg), options...)
}

func getWriteSyncer(logCfg *configs.LogConfig) zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   strings.Join([]string{logCfg.Dir, logCfg.FileName}, "/"),
		MaxSize:    logCfg.MaxSize,
		MaxBackups: logCfg.MaxBackups,
		MaxAge:     logCfg.MaxAge,
	}
	return zapcore.AddSync(l)
}

func getZapCore(logCfg *configs.LogConfig) zapcore.Core {
	var encoder zapcore.Encoder

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderCfg.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(l.String())
	}

	if logCfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	return zapcore.NewCore(encoder, getWriteSyncer(logCfg), level)
}

func tryCreateDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return err
}

func setLogLevel(lvl string) {
	switch lvl {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}
