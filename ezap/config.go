package ezap

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	LogLevel    zap.AtomicLevel //	动态修改日志级别
	ProjectName string
	JSONFormat  bool
	Console     bool
	LogFile     *Logfile
}

type Logfile struct {
	// Enable     bool
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type Logger struct {
	Config *Config
	// Logger *zap.Logger
	Logger *zap.SugaredLogger
}

func NewConfig() *Config {
	return &Config{
		LogLevel:    zap.NewAtomicLevel(),
		ProjectName: "",
		JSONFormat:  true,
		Console:     true,
		LogFile: &Logfile{
			// Enable:     false,
			FileName:   "",
			MaxSize:    100,
			MaxAge:     24,
			MaxBackups: 7,
			Compress:   true,
		},
	}
}

func (c *Config) SetLogLevel(lv string) {
	switch lv {
	case "info":
		c.LogLevel.SetLevel(zap.InfoLevel)
	case "debug":
		c.LogLevel.SetLevel(zap.DebugLevel)
	case "warn":
		c.LogLevel.SetLevel(zap.WarnLevel)
	case "error":
		c.LogLevel.SetLevel(zap.ErrorLevel)
	case "panic":
		c.LogLevel.SetLevel(zap.PanicLevel)
	case "fatal":
		c.LogLevel.SetLevel(zap.FatalLevel)
	default:
		c.LogLevel.SetLevel(zap.InfoLevel)
	}
}

func (c *Config) SetProjectName(projectname string) {
	c.ProjectName = projectname
}

func (c *Config) EnableJSONFormat() {
	c.JSONFormat = true
}

func (c *Config) DisableJSONFormat() {
	c.JSONFormat = false
}

func (c *Config) EnableConsole() {
	c.Console = true
}

func (c *Config) DisableConsole() {
	c.Console = false
}

func (c *Config) SetLogFile(name string) {
	// c.LogFile.Enable = true
	c.LogFile.FileName = name
}

func (c *Config) SetLogrotate(maxsize, maxage, maxbackups int, compress bool) {
	c.LogFile.MaxSize = maxsize
	c.LogFile.MaxAge = maxage
	c.LogFile.MaxBackups = maxbackups
	c.LogFile.Compress = compress
}

func (l *Logger) ApplyConfig() {
	conf := l.Config
	cores := []zapcore.Core{}
	var encoder zapcore.Encoder
	encoderConfig := zap.NewProductionConfig().EncoderConfig
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-1-2T15:04:05.000Z0700"))
	}

	conf.LogLevel.SetLevel(zap.InfoLevel)

	if conf.JSONFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	if conf.Console {
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(encoder, writer, conf.LogLevel)
		cores = append(cores, core)
	}

	if conf.LogFile.FileName != "" {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.LogFile.FileName,
			MaxSize:    conf.LogFile.MaxSize,
			MaxAge:     conf.LogFile.MaxAge,
			MaxBackups: conf.LogFile.MaxBackups,
			Compress:   conf.LogFile.Compress,
			LocalTime:  true,
		})
		writer := zapcore.AddSync(w)
		core := zapcore.NewCore(encoder, writer, conf.LogLevel)
		cores = append(cores, core)
	}

	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	if conf.ProjectName != "" {
		logger = logger.Named(conf.ProjectName)
	}

	defer logger.Sync()
	l.Logger = logger.Sugar()
}
