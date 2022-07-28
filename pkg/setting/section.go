package setting

import (
	"time"
)

type Config struct {
	Server Server `toml:"Server"`
	App    App    `toml:"App"`
	BD     DB     `toml:"Database"`
	JWT    JWT    `toml:"JWT"`
}

type Server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type App struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LogFileName          string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
}

type DB struct {
	Address      string
	Username     string
	Password     string
	Database     string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWT struct {
	Secret string
	Issuer string
	Expire time.Duration
}
