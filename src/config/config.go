package config

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:"database"`
	Log      Log      `mapstructure:"log"`
}

type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type Database struct {
	TableName string `mapstructure:"table_name"`
}

func InitConfig(configPath string) (config Config, err error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	/* default */
	v.SetDefault("log_level", "INFO")
	v.SetDefault("log_format", "console")

	defaultPath := `./config`

	if configPath == "" {
		configPath = defaultPath
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AddConfigPath(configPath)

	files, _ := ioutil.ReadDir(configPath)
	index := 0

	for _, file := range files {
		if filepath.Ext("./"+file.Name()) != ".yaml" && filepath.Ext("./"+file.Name()) != ".yml" {
			continue
		}

		v.SetConfigName(file.Name())
		var err error
		if index == 0 {
			err = v.ReadInConfig()
		} else {
			err = v.MergeInConfig()
		}
		if err == nil {
			index++
		}
	}

	if err = v.Unmarshal(&config); err != nil {
		return
	}

	return
}
