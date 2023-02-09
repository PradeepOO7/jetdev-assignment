package config

import (
	"blogs/model"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	var err error
	username := viper.Get("USERNAME")
	password := viper.Get("PASSWORD")
	host := viper.Get("HOST")
	port := viper.Get("PORT")
	db := viper.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, db)
	model.SqlOrm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default})
	if err != nil {
		fmt.Printf("Error Connecting to Database %v \n", err)
	}
	fmt.Println("DB connected successfully")
	//model.SqlOrm.AutoMigrate(&model.Comment{},&model.Article{})
}
