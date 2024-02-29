package logger

import (
	"github.com/Snorkin/auth_service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger interface {
	Init()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type ServerLogger struct {
	cfg   *config.Config
	sugar *zap.SugaredLogger
}

func CreateLogger(cfg *config.Config) *ServerLogger {
	return &ServerLogger{cfg: cfg}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (s *ServerLogger) getLevel(cfg *config.Config) zapcore.Level {
	level, ok := loggerLevelMap[cfg.Logger.Level]
	if !ok {
		return zapcore.DebugLevel
	}

	return level
}

func (s *ServerLogger) Init() {
	level := s.getLevel(s.cfg)

	writer := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if s.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.LevelKey = "Level"
	encoderCfg.CallerKey = "Caller"
	encoderCfg.TimeKey = "Time"
	encoderCfg.NameKey = "Name"
	encoderCfg.MessageKey = "Message"

	var encoder zapcore.Encoder
	if s.cfg.Logger.Encoding == "Console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, writer, zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	s.sugar = logger.Sugar()
	//if err := s.sugar.Sync(); err != nil {
	//	s.sugar.Error(err)
	//}
}

func (s *ServerLogger) Debug(args ...interface{}) {
	s.sugar.Debug(args...)
}

func (s *ServerLogger) Debugf(template string, args ...interface{}) {
	s.sugar.Debugf(template, args...)
}

func (s *ServerLogger) Info(args ...interface{}) {
	s.sugar.Info(args...)
}

func (s *ServerLogger) Infof(template string, args ...interface{}) {
	s.sugar.Infof(template, args...)
}

func (s *ServerLogger) Warn(args ...interface{}) {
	s.sugar.Warn(args...)
}

func (s *ServerLogger) Warnf(template string, args ...interface{}) {
	s.sugar.Warnf(template, args...)
}

func (s *ServerLogger) Error(args ...interface{}) {
	s.sugar.Error(args...)
}

func (s *ServerLogger) Errorf(template string, args ...interface{}) {
	s.sugar.Errorf(template, args...)
}

func (s *ServerLogger) DPanic(args ...interface{}) {
	s.sugar.DPanic(args...)
}

func (s *ServerLogger) DPanicf(template string, args ...interface{}) {
	s.sugar.DPanicf(template, args...)
}

func (s *ServerLogger) Panic(args ...interface{}) {
	s.sugar.Panic(args...)
}

func (s *ServerLogger) Panicf(template string, args ...interface{}) {
	s.sugar.Panicf(template, args...)
}

func (s *ServerLogger) Fatal(args ...interface{}) {
	s.sugar.Fatal(args...)
}

func (s *ServerLogger) Fatalf(template string, args ...interface{}) {
	s.sugar.Fatalf(template, args...)
}
