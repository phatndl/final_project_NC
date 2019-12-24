package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ConfigSchema struct {
	MongoDB struct {
		Uri      string `mapstructure"Uri"`
		Hostname string `mapstructure"Hostname"`
		Username string `mapstructure"Username"`
		Password string `mapstructure"Password"`
	} `mapstructure"MongoDB"`
	JWTSecret struct {
		JWTEncodedKey string `mapstructure"JWTEncodedKey"`
	}`mapstructure"JWTSecret"`
}

var Config ConfigSchema

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("error in decode into struct: %v \n", err)
	}
	fmt.Println("config done", Config)
}
