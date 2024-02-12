package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
