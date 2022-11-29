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

	var departments = []Department{
		{
			Name: "XYZ",
		},
		{
			Name: "ABC",
		},
	}

	var employees = []Employee{
		{
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "mail1@mail.com",
			City:         "New York",
			DepartmentID: 1,
		},
		{
			FirstName:    "Jane",
			LastName:     "Doe",
			Email:        "mail2@mail.com",
			City:         "New York",
			DepartmentID: 2,
		},
		{
			FirstName:    "John",
			LastName:     "Smith",
			Email:        "mail3@mail.com",
			City:         "Jakarta",
			DepartmentID: 1,
		},
	}

	result := db.Create(&departments)
	result = db.Create(&employees)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("berhasil")
}
