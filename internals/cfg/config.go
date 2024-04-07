package cfg

import (
	"fmt"
	// "log"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB     DB
	Server Server
}
type DB struct {
	Port     string
	Name     string
	User     string
	Password string
	Host     string
}

type Server struct {
	Host string
	Port int
}

func LoadAndStore(addr string) (*Configuration, error) {
	config := &Configuration{}
	viper.AddConfigPath(addr)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("could not unmarshal the config file: %v", err)
	}
	return config, nil
}
