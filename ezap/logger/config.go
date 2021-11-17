package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	LogLevel    zap.AtomicLevel
	ProjectName string
	JSONFormat  bool
	Console     bool
	AddCaller   bool
	LogFile     *Logfile
}

type Logfile struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type Logger struct {
	Config *Config
	Logger *zap.SugaredLogger
}

func NewConfig() *Config {
	return &Config{
		LogLevel:    zap.NewAtomicLevel(),
		ProjectName: "",
		JSONFormat:  true,
		Console:     true,
		AddCaller:   false,
		LogFile: &Logfile{
			FileName:   "",
			MaxSize:    100,
			MaxAge:     24,
			MaxBackups: 7,
			Compress:   true,
		},
	}
}

// 应用修改后的配置
func (l *Logger) SyncConfig() {
	conf := l.Config
	cores := []zapcore.Core{}
	var logger *zap.Logger
	var encoder zapcore.Encoder
	// cfg := zap.NewProductionConfig()
	// encoderConfig := cfg.EncoderConfig
	encoderConfig := zap.NewProductionConfig().EncoderConfig
	encoderConfig.LevelKey = "lv"
	encoderConfig.NameKey = "pj"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-1-2T15:04:05.000Z0700"))
	}

	if conf.JSONFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		// 终端日志级别显示颜色
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
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

	if conf.AddCaller {
		logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	} else {
		logger = zap.New(zapcore.NewTee(cores...))
	}

	if conf.ProjectName != "" {
		logger = logger.Named(conf.ProjectName)
	}

	defer logger.Sync()
	l.Logger = logger.Sugar()
}

// 设置输出的日志等级
// Options: info debug warn error panic fatal
func (l *Logger) SetLevel(lv string) {
	setlevel := l.Config.LogLevel.SetLevel
	switch lv {
	case "debug":
		setlevel(zap.DebugLevel)
	case "info":
		setlevel(zap.InfoLevel)
	case "warn":
		setlevel(zap.WarnLevel)
	case "error":
		setlevel(zap.ErrorLevel)
	case "panic":
		setlevel(zap.PanicLevel)
	case "fatal":
		setlevel(zap.FatalLevel)
	default:
		setlevel(zap.InfoLevel)
	}
	l.SyncConfig()
}

// 设置工程名称
func (l *Logger) SetProjectName(projectname string) {
	l.Config.ProjectName = projectname
	l.SyncConfig()
}

// 启用 JSON 格式输出
func (l *Logger) EnableJSONFormat() {
	l.Config.JSONFormat = true
	l.SyncConfig()
}

// 关闭 JSON 格式输出
func (l *Logger) DisableJSONFormat() {
	l.Config.JSONFormat = false
	l.SyncConfig()
}

// 开启调用日志输出
func (l *Logger) EnableCaller() {
	l.Config.AddCaller = true
	l.SyncConfig()
}

// 关闭调用日志输出
func (l *Logger) DisableCaller() {
	l.Config.AddCaller = false
	l.SyncConfig()
}

// 开启控制台日志输出
func (l *Logger) EnableConsole() {
	l.Config.Console = true
	l.SyncConfig()
}

// 关闭控制台日志输出
func (l *Logger) DisableConsole() {
	l.Config.Console = false
	l.SyncConfig()
}

// 设置日志保存文件，例如 /var/log/myapp/myapp.log
// 如果配置为空，则不会保存到日志文件
func (l *Logger) SetLogFile(name string) {
	l.Config.LogFile.FileName = name
	l.SyncConfig()
}

// 配置日志滚动配置，默认 100M，24h，7天，启用压缩
func (l *Logger) SetLogrotate(maxsize, maxage, maxbackups int, compress bool) {
	f := l.Config.LogFile
	f.MaxSize = maxsize
	f.MaxAge = maxage
	f.MaxBackups = maxbackups
	f.Compress = compress
	l.SyncConfig()
}
