package configs

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Logger LoggerConf
	DB     DBConf
	Server Grpc
	Kafka  Kafka
}

type Kafka struct {
	Topic  string
	Broker string
	Time   int64
}

type LoggerConf struct {
	Level string
	Path  string
}

type Grpc struct {
	Port string
}

type DBConf struct {
	Dsn string
}

func NewConfig(path string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		fmt.Println(err)
		return Config{}, err
	}
	return config, nil
}
