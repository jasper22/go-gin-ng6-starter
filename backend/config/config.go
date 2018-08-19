package config

import (
    "github.com/spf13/viper"
    "log"
    "os"
)

var config *viper.Viper

func init()  {
    var err error
    env := os.Args[1]
    v := viper.New()
    v.SetConfigType("yaml")
    v.SetConfigName(env)
    v.AddConfigPath("../config/")
    err = v.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }
    config = v
}

func GetConfig() *viper.Viper {
    return config
}
