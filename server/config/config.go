package config

import "github.com/spf13/viper"

// Конфигурация Базы Данных и Сервера
type Config struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbDriver   string `mapstructure:"DB_DRIVER"`
	Host       string `mapstructure:"HOST"`
	Port       string `mapstructure:"PORT"`
	StaticRoot string `mapstructure:"STATIC_ROOT"`
	SigningKey string `mapstructure:"SINGING_KEY"`
	LogFile    string `mapstructure:"LOGFILE"`
}

func LoadConfig(path string, file string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
