package database

import (
	"fmt"

	"github.com/rajenderK7/go-mysql/pkg/models"
	"github.com/rajenderK7/go-mysql/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBPORT     string
}

var (
	db *gorm.DB
)

func ConnectDB(config *DBConfig) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBPORT,
		config.DBName,
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.CheckDBErr(err)
	db = d
}

func RunMigrations() {
	models.InitUser(db)
	models.InitMovie(db)
	models.InitBooking(db)
}

func GetDB() *gorm.DB {
	return db
}
