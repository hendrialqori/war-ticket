package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlConnection(config *DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset,
		config.ParseTime,
		config.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect mysql", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}

	connection.SetMaxIdleConns(config.Idle)
	connection.SetMaxOpenConns(config.Max)
	connection.SetConnMaxLifetime(60 * time.Second)

	log.Println("Database connected!")

	return db
}
