package config

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    "fmt"
    "go-gin-ng6-starter/backend/models"
    "strings"
)

var db *gorm.DB //database

func init() {
    var (
        username string
        password string
        dbName   string
        dbHost   string
    )
    cfg := GetConfig() // load global viper config
    dbUrl := cfg.GetString("DATABASE_URL")
    username = strings.Split(strings.Split(dbUrl, "//")[1], ":")[0]
    password = strings.Split(strings.Split(dbUrl, ":")[2], "@")[0]
    dbName = strings.Split(strings.Split(dbUrl, "@")[1], "/")[1]
    dbHost = strings.Split(strings.Split(dbUrl, "@")[1], ":")[0]
    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

    conn, err := gorm.Open("postgres", dbUri)
    if err != nil {
        fmt.Print(err)
    }

    db = conn
    // Database migration
    db.Debug().AutoMigrate(&models.User{})
    db.Debug().AutoMigrate(&models.Address{})
    db.Debug().AutoMigrate(&models.CreditCard{})
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
    return db
}
