package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	DB struct {
		Port     string
		Name     string
		User     string
		Password string
		Host     string
	}
	Server struct {
		Host string
		Port int
	}
}

func LoadAndStore() (*Configuration, error) {

	config := &Configuration{}
	fmt.Println(config)
	viper.AddConfigPath("internals/cfg/config.yaml")
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	return nil, fmt.Errorf("could not read the config file: %v", err)
	// }

	// if err = viper.Unmarshal(&config); err != nil {
	// 	return nil, fmt.Errorf("could not unmarshal the config file: %v", err)
	// }
	return nil, nil
}

func main() {
	LoadAndStore()
	// if err != nil {
	// log.Fatalf("could not load: %v", err)
	// // }
	// fmt.Printf("%+v\n", *config)
}
