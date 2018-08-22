package config

import (
    "github.com/spf13/viper"
    "log"
        )

var config *viper.Viper

func init()  {
    var env string
    v := viper.New()
    v.AutomaticEnv() // auto scan for ENV
    if v.GetString("APP_ENV") == "" {
        log.Printf("No env specified. Using dev as default.")
        env = "dev"
    } else {
        env = v.GetString("APP_ENV")
    }
    log.Printf("Using environment settings for %s \n", env)

    // config parsing
    var err error
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
