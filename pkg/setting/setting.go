package setting

import (
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
