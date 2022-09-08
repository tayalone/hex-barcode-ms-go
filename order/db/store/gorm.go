package store

import (
	"fmt"
	"log"
	"os"

	"github.com/tayalone/hex-barcode-ms-go/barcode/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*RDB is Definition of Value */
type RDB struct {
	db     *gorm.DB
	errMsg string
}

var rdb = RDB{}

func migrate(db *gorm.DB) {
	// db.Set("gorm:table_options", "ENGINE=InnoDB")

	// /  about 'barcode_condition'
	if (db.Migrator().HasTable(&core.BarcodeCondition{})) {
		log.Println("Table Existing, Drop IT")

		db.Migrator().DropTable(&core.BarcodeCondition{})
	}
	db.AutoMigrate(&core.BarcodeCondition{})
	log.Println("Create 'barcode_conditions'")

	// / Add Initail Data
	initBCC := []core.BarcodeCondition{
		{
			CourierCode:   "DHL",
			IsCod:         true,
			StartBarcode:  "DCA00000001XTH",
			BatchSize:     100,
			PrevCondLogID: 1,
			CondLogID:     101,
		}, {
			CourierCode:   "DHL",
			IsCod:         false,
			StartBarcode:  "DNA00000001XTH",
			BatchSize:     300,
			PrevCondLogID: 1,
			CondLogID:     301,
		},
	}
	db.Create(initBCC)

	log.Println("Create Initial 'BarcodeCondition' Data")
}

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

	if os.Getenv("RDM_MIGRATION") == "true" {
		migrate(db)
	}

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
