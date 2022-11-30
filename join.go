package main

import (
	"encoding/json"
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

type Joined struct {
	Name           string
	DepartmentName string
}

func main() {
	dsn := "host=192.168.0.104 user=postgres password=password dbname=hrApp port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	rows := []Joined{}

	db.Table("departments").
		Select("employees.first_name as name, departments.name as department_name").
		Joins("inner join employees on departments.id = employees.departement_id").
		Scan(&rows)

	result, _ := json.Marshal(rows)

	fmt.Println(string(result))
}
