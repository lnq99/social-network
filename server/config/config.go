package config

// Конфигурация Базы Данных и Сервера
type Config struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbDriver   string
	Host       string
	Port       string
}

func LoadConfig(path string, file string) (config Config, err error) {
	return
}
