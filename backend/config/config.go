package config

import (
    "github.com/spf13/viper"
    "log"
    "os"
    "fmt"
)

var config *viper.Viper

func init()  {
    // env parsing
    if len(os.Args) < 2{
        log.Fatal("Must provide env argument in run command")
    }
    env := os.Args[1]
    fmt.Printf("Using environment settings for %s \n", env)

    // TODO Implement config file for prod (and test)

    // config parsing
    var err error
    v := viper.New()
    v.AutomaticEnv() // auto scan for ENV
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
