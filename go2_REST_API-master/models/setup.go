package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	// To get the value from the config file using key
	// viper package read .env
	viper_user := "postgres"     //viper.Get("POSTGRES_USER")
	viper_password := "2709197q" // viper.Get("POSTGRES_PASSWORD")
	viper_db := "postgres"       // viper.Get("POSTGRES_DB")
	viper_host := "127.0.0.1"    //viper.Get("POSTGRES_HOST")
	viper_port := "5432"         //viper.Get("POSTGRES_PORT")
	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)
	fmt.Println("conname is\t\t", prosgret_conname)
	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&Book{})
	// Initialise value
	m := Book{Author: "author1", Title: "title1"}
	db.Create(&m)
	return db
}
