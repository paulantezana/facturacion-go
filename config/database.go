package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// GetConnection get connection database
func GetConnection() *gorm.DB {
	c := GetConfig()

	// Connection database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Database.User, c.Database.Pass, c.Database.Server, c.Database.Port, c.Database.Database)
	}

	//dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Pass, c.Server, c.Port, c.Database)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
