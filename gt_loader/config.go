package gt_loader

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// LoadConfigCallback 加载配置文件回调函数
// @param v *viper.Viper
type LoadConfigCallback func(ctx context.Context, v *viper.Viper)

// LoadConfig 使用viper加载配置文件
// @param configFile string 配置文件路径
// @param callback LoadConfigCallback 回调函数
func LoadConfig(ctx context.Context, configFile string, callback LoadConfigCallback) (err error) {
	cfgLogger.CInfo(ctx, "加载配置文件开始: %s", configFile)
	v := viper.New()
	v.SetConfigFile(configFile)
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		callback(ctx, v)
	})

	if err = v.ReadInConfig(); err != nil {
		return
	}

	callback(ctx, v)

	cfgLogger.CInfo(ctx, "加载配置文件完成: %s", configFile)
	return
}
