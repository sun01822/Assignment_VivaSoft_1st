package connection

import (
	"Assignment_Vivasoft/package/config"
	"Assignment_Vivasoft/package/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// connect to database
func Connect(){
	dbConfig := config.LocalConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIp, dbConfig.DBName)
	// dsn := "root:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err)
	}
	fmt.Println("Database Connected")
	DB = db
}

// creating new table in bookstore database
func migrate(){
	DB.Migrator().AutoMigrate(&models.BookDetails{})
	DB.Migrator().AutoMigrate(&models.AuthorDetails{})
	DB.Migrator().AutoMigrate(&models.UserDetails{})
}

// calling to connect function to initialize connection
func GetDB() *gorm.DB{
	if DB == nil{
		Connect()
	}
    migrate()
	return DB
}
