package store

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*RDB is Definition of Value */
type RDB struct {
	db     *gorm.DB
	errMsg string
}

var rdb = RDB{}

/*New do Create Rdb Connection*/
func New() *RDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("RDM_HOST"),
		os.Getenv("RDM_USER"),
		os.Getenv("RDM_PASSWORD"),
		os.Getenv("RDM_DB"),
		os.Getenv("RDM_PORT"),
		os.Getenv("TIME_ZONE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Println("FAIL: Connect RDB Error", err.Error())

		rdb.db = nil
		rdb.errMsg = err.Error()
		return &rdb

	}
	log.Println("Connect RDB Success!!!")
	rdb.db = db
	rdb.errMsg = ""

	return &rdb
}

/*Status return Connect Status*/
func Status() error {
	if rdb.errMsg != "" {
		return fmt.Errorf("%s", rdb.errMsg)
	}
	return nil
}

/*GetInstant return db instant*/
func (r *RDB) GetInstant() *gorm.DB {
	if r.errMsg != "" {
		return nil
	}
	return r.db
}
