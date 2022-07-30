package setting

import (
	"strings"
	"time"

	"BlogService/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperSetting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*ViperSetting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("toml")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//通过协程实现热更新
	vs := &ViperSetting{vp: vp}
	vs.WatchSettingChange()
	return vs, nil
}

func (vs *ViperSetting) WatchSettingChange() {
	go func() {
		vs.vp.WatchConfig()
		vs.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = vs.ReloadAllSection()
		})
	}()
}

func InitSetting(port, runMode, configPath string) error {
	s, err := NewSetting(strings.Split(configPath, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.Config.Server)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.Config.App)
	if err != nil {
		return err
	}
	err = s.ReadSection("DataBase", &global.Config.BD)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.Config.JWT)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.Config.Email)
	if err != nil {
		return err
	}
	err = s.ReadSection("Limiter", &global.Config.Limiter)
	if err != nil {
		return err
	}

	if port != "" {
		global.Config.Server.HttpPort = port
	}
	if runMode != "" {
		global.Config.Server.RunMode = runMode
	}
	global.Config.Server.ReadTimeout *= time.Second
	global.Config.Server.WriteTimeout *= time.Second
	global.Config.JWT.Expire *= time.Second
	return nil
}
