package orm

import (
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCompany(t *testing.T) {
	dsn := "postgresql://postgres:123456@0.0.0.0:5432/test?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// var company Company
	var companies []Company
	db.Find((&companies))
	fmt.Print(companies)
}
