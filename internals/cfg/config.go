package cfg

import (
	"fmt"
	"log"
	"os"

	// "log"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB     DB
	Server Server
	Redis  Redis
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
	Port string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

func LoadAndStore(addr string) (*Configuration, error) {
	config := &Configuration{}
	workingdir, err := os.Getwd()
	viper.AddConfigPath(workingdir + "/" + addr)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()

	if err != nil {
		log.Println(workingdir)
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("could not unmarshal the config file: %v", err)
	}
	return config, nil
}
