package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=localhost user=postgres password=kamil dbname=fileexplorer port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Nie można połączyć z bazą ")
	}
}
