package logger

func New() *Logger {
	logger := &Logger{
		Config: NewConfig(),
	}
	logger.SyncConfig()
	return logger
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.Logger.Debugf(template, args...)
}
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.Logger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}
func (l *Logger) Infof(template string, args ...interface{}) {
	l.Logger.Infof(template, args...)
}
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.Logger.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.Logger.Warnf(template, args...)
}
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.Logger.Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.Logger.Errorf(template, args...)
}
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.Logger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.Logger.Fatalf(template, args...)
}
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.Logger.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Logger.Panic(args...)
}
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.Logger.Panicf(template, args...)
}
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.Logger.Panicw(msg, keysAndValues...)
}
