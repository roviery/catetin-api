package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/roviery/catetin-api/models"
)

func init() {
	dbInstance = &DBInstance{initializer: dbInit}
}

var dbInstance *DBInstance

type DBInstance struct {
	initializer func() interface{}
	instance    interface{}
	once        sync.Once
}

func (i *DBInstance) Instance() interface{} {
	i.once.Do(func() {
		i.instance = i.initializer()
	})
	return i.instance
}

func dbInit() interface{} {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "admin"
	dbPass := "admin"
	dbName := "catetin"

	conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open("postgres", conn)

	if err != nil {
		log.Default().Println(err)
		os.Exit(1)
		db.Close()
	}

	db.AutoMigrate(&models.User{})

	return db
}

func DB() *gorm.DB {
	return dbInstance.Instance().(*gorm.DB)
}
