package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	*Database
	*HTTP
	*Logger
}

func NewConfig() *Config {
	return &Config{
		Database: NewDatabase(),
		HTTP:     NewHTTP(),
		Logger:   NewLogger(),
	}
}

func GetEnvStr(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		v = def
	}
	return v
}

func GetEnvInt(key string, def int) int {
	vString := os.Getenv(key)
	if vString == "" {
		return def
	}
	v, err := strconv.Atoi(vString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot convert %s to int, use default %d: %v", vString, def, err)
		v = def
	}
	return v
}
