package config

import "github.com/ButterflyGate/logger/levels"

type Logger struct {
	lvl levels.LogLevel
}

func NewLogger() *Logger {
	lvl := GetEnvInt("LOG_LEVEL", d.log.logLevel)
	d := &Logger{
		lvl: levels.LogLevel(lvl),
	}
	return d
}

func (l *Logger) LogLevel() levels.LogLevel {
	return l.lvl
}
