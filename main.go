package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func initConfig() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}
	fmt.Printf("Current Directory: %s\n", wd)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/data/project/campwiz-bot")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}
func initDB() {
	username := viper.GetString("user")
	password := viper.GetString("password")
	database := "s55789__campwiz_test"
	host := "tools.db.svc.wikimedia.cloud"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
	conn := mysql.Open(dsn)
	if conn == nil {
		log.Fatalf("Failed to create connection: %v", conn)
	}
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)

	}
	co, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}
	if err := co.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	co.Close()

}

func main() {
	initConfig()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		initDB()
		c.String(200, "Database connection successful")
	})

	r.Run("0.0.0.0:8000")
}
