package logger

import (
	"log"
	"os"
	"time"

	"github.com/mpuzanov/otus-go/calendar/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Logger глобальная переменная для работы
type Logger *zap.Logger

func getZapLevel(level string) zap.AtomicLevel {
	switch level {
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02.01.2006 03:04:05 PM"))
}

//NewLogger Возвращаем инициализированный логер
func NewLogger(config config.LogConf) *zap.Logger {
	EncodingFormat := "json"
	if !config.LogFormatJSON {
		EncodingFormat = "console"
	}
	OutputPath, ErrorOutputPath := "stderr", "stderr"

	if config.LogFile != "" {
		_, err := os.Create(config.LogFile)
		if err != nil {
			log.Printf("ошибка создания файла логов %s %v", config.LogFile, err)
		} else {
			OutputPath, ErrorOutputPath = config.LogFile, config.LogFile
		}
	}
	cfg := zap.Config{
		Encoding:         EncodingFormat,
		Level:            getZapLevel(config.LogLevel),
		OutputPaths:      []string{OutputPath},
		ErrorOutputPaths: []string{ErrorOutputPath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey: "time",
			//EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeTime: syslogTimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := cfg.Build()
	return logger
}

//InitLogger Вариант инициализации логера
func InitLogger(config config.LogConf) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = syslogTimeEncoder //zapcore.ISO8601TimeEncoder
	cfg.CallerKey = "caller"
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewJSONEncoder(cfg)
	if !config.LogFormatJSON {
		encoder = zapcore.NewConsoleEncoder(cfg)
	}
	writerSyncer := zapcore.Lock(os.Stderr) //os.Stdout
	if config.LogFile != "" {
		file, err := os.Create(config.LogFile)
		if err != nil {
			log.Printf("ошибка создания файла логов %s %v", config.LogFile, err)
		} else {
			writerSyncer = zapcore.Lock(file)
		}
	}
	level := getZapLevel(config.LogLevel)
	logger := zap.New(zapcore.NewCore(encoder, writerSyncer, level))

	//zap.ReplaceGlobals(logger)
	return logger
}
