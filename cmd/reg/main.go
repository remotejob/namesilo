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

	res0, res1, err := config.Client.RegisterDomain("saunassa.xyz", 1, "payment_id=172297", "auto_renew=0", "ns1=ns1.digitalocean.con", "ns2=ns2.digitalocean.con", "ns3=ns3.digitalocean.con")
	if err != nil {

		log.Panicln("err",err.Error())
	}

	log.Println("res0", res0,"res1",res1)

}