package repositorie

import (
	"fmt"
	"log"

	util "github.com/Faizal-Asep/test-tech/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	host := util.Config.DbHost
	user := util.Config.DbUser
	password := util.Config.DbPassword
	dbName := util.Config.DbDatabase

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&HandPhone{})
}
