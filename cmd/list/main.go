package main

import (
	"log"

	"github.com/remotejob/namesilo"
	"github.com/spf13/viper"
)

type Constants struct {
	ApiKey string
}

type Config struct {
	Constants
	Client namesilo.Client
}

func initViper() (Constants, error) {

	viper.SetConfigFile("config.toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {

		return Constants{}, err
	}

	var constants Constants

	err = viper.Unmarshal(&constants)
	return constants, err

}
func New() (*Config, error) {
	log.Println("New")

	config := Config{}
	constants, err := initViper()
	config.Constants = constants
	if err != nil {
		return &config, err
	}

	log.Println(config.Constants.ApiKey)

	config.Client = namesilo.NewClient(config.Constants.ApiKey)

	return &config, err
}
func main() {

	config, err := New()
	if err != nil {

		log.Println(err)
	}

	// log.Println(config.ApiKey)

	mydomains, err := config.Client.ListDomains("")
	if err != nil {

		panic(err)
	}

	log.Println(mydomains)

	dominf, err := config.Client.GetDomainInfo("poika.top")
	if err != nil {

		panic(err)
	}
	log.Println(dominf)

}
