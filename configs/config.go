package configs

import "time"

type Config struct {
	Server  Server             `toml:"Server"`
	App     App                `toml:"App"`
	BD      DB                 `toml:"Database"`
	JWT     JWT                `toml:"JWT"`
	Email   Email              `toml:"Email"`
	Limiter map[string]Limiter `toml:"Limiter"`
}

type Server struct {
	RunMode        string
	HttpPort       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ContextTimeout time.Duration
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

type Email struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
	To       []string
}

type Limiter struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
