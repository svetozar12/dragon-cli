package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() {
	var err error
	DB, err = gorm.Open(postgres.Open("postgresql://postgres:GApjZrHqsbqicm2M7lCe@containers-us-west-172.railway.app:7894/railway"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		panic(err)
	}
}
