package main

import (
	"chickenFarm/db"
)

func main() {
	// runMetrics()
	// testDB()
	TestDB()
}

func TestDB() {
	// db.GetAllInfo()
	db.InitDB()
}
