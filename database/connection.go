package database

import (
	"backend/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}
	log.Println("Connected to database")
	return nil
}

func MigrateLog(table *entity.Log) {
	Connector.AutoMigrate(&table)
	Connector.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&table)
	log.Println("Migrated table Log")
}

// func MigratePemeriksaan(table *entity.Pemeriksaan) {
// 	Connector.AutoMigrate(&table)
// 	log.Println("Migrated table pemeriksaan")
// }
