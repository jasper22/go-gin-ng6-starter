package config

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    "fmt"
    "strings"
    "github.com/app8izer/go-gin-ng6-starter/backend/models"
)

var db *gorm.DB // database

func init() {
    cfg := GetConfig() // load global viper config
    dbUrl := cfg.GetString("DATABASE_URL")
    username, password, dbName, dbHost := splitDbUrl(dbUrl)
    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) // Build connection string

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

func splitDbUrl(dbUrl string) (username string, password string, dbName string, dbHost string) {
    username = strings.Split(strings.Split(dbUrl, "//")[1], ":")[0]
    password = strings.Split(strings.Split(dbUrl, ":")[2], "@")[0]
    dbName = strings.Split(strings.Split(dbUrl, "@")[1], "/")[1]
    dbHost = strings.Split(strings.Split(dbUrl, "@")[1], ":")[0]
    return username, password, dbName, dbHost
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
    return db
}
