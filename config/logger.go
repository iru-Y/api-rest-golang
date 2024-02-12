package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer *io.Writer
}

func NewLogger(p string) *Logger {
	w:= io.Writer(os.Stdout)
	l:= log.New(w, p, log.Ldate|log.Ltime)
	return &Logger{
		debug: log.New(w, "DEBUG", l.Flags()),
		info: log.New(w, "INFO", l.Flags()),
		warn: log.New(w, "WARN", l.Flags()),
		err: log.New(w, "ERR", l.Flags()),
		writer: &w,
	}}

	
func (l *Logger) DEBUG(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}


func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
