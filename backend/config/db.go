package config

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    "fmt"
    "github.com/app8izer/go-gin-ng6-starter/backend/models"
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
    username = cfg.GetString("postgres.db_usr")
    password = cfg.GetString("postgres.db_pass")
    dbName = cfg.GetString("postgres.db_name")
    dbHost = cfg.GetString("postgres.db_host")
    dbUri2 := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
    //fmt.Println(dbUri2)

    conn, err := gorm.Open("postgres", dbUri2) // changed here
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
