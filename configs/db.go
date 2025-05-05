package configs

import (
	"fmt"
	"makanan-app/models"

	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbPort := os.Getenv("DB_PORT")
	dbPortInt, _ := strconv.Atoi(dbPort)
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), dbPortInt, os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.News{}, &models.Comment{})
	return db
}