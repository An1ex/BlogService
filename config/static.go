package config

type MySQL struct {
	Address  string `toml:"address"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}
