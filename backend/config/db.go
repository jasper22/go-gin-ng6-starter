package config

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/jinzhu/gorm"
    "fmt"
    "strings"
    "github.com/app8izer/go-gin-ng6-starter/backend/models"
    _ "github.com/golang-migrate/migrate/source/file"
    "database/sql"
    "github.com/golang-migrate/migrate/database/postgres"
    "log"
    "github.com/golang-migrate/migrate"
)

var db *gorm.DB // database

func init() {
    cfg := GetConfig() // load global viper config
    dbUrl := cfg.GetString("DATABASE_URL")
    username, password, dbName, dbHost, dbPort := splitDbUrl(dbUrl)
    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort) // Build connection string

    conn, err := gorm.Open("postgres", dbUri)
    if err != nil {
        fmt.Print(err)
    }

    db = conn
    // Database migration
    if cfg.GetString("APP_ENV") == "prod" {
        performMigrations(db.DB(), dbName)
    } else {
        db.Debug().AutoMigrate(&models.User{})
    }

}

func splitDbUrl(dbUrl string) (username string, password string, dbName string, dbHost string, dbPort string) {
    username = strings.Split(strings.Split(dbUrl, "//")[1], ":")[0]
    password = strings.Split(strings.Split(dbUrl, ":")[2], "@")[0]
    dbName = strings.Split(strings.Split(dbUrl, "@")[1], "/")[1]
    dbHost = strings.Split(strings.Split(dbUrl, "@")[1], ":")[0]
    dbPort = strings.Split(strings.Split(dbUrl, ":")[3], "/")[0]
    return username, password, dbName, dbHost, dbPort
}

func performMigrations(db *sql.DB, dbName string) {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        log.Fatal(err)
    }
    m, err := migrate.NewWithDatabaseInstance("file://backend/migrations", dbName, driver)
    if err != nil {
        log.Fatal(err)
    }
    // Migrate all the way up ...
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
    return db
}
