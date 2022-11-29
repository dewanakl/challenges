package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID           int `gorm:"primaryKey"`
	FirstName    string
	LastName     string
	Email        string
	City         string
	DepartmentID int
	gorm.Model
}

type Department struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

func main() {
	dsn := "host=192.168.0.104 user=postgres password=password dbname=hrApp port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	departments := []Department{}
	employees := []Employee{}

	result := db.Find(&departments)
	result = db.Find(&employees)

	if result.Error != nil {
		panic(result.Error)
	}

	for _, v := range departments {
		fmt.Println(v)
	}

	for _, v := range employees {
		fmt.Println(v)
	}
}
