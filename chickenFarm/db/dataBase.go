package db

import (
	"chickenFarm/model"
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm/clause"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&model.UpInfo{})
	// InsertData()
	GetAllInfo()
}
func InitDB() {
	db := connct()
	db.AutoMigrate(&model.UpInfo{})

}
func connct() *gorm.DB {
	//
	db, err := gorm.Open(sqlite.Open("/home/caidong/gitRepositories/chickenFarm/chickenFarm/db/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func GetAllInfo() string {
	db := connct()
	var datas []model.UpInfo
	result := db.Find(&datas)
	var list2 []model.UpInfo
	for _, res := range datas {

		res.Flag = "https://flagcdn.com/" + strings.ToLower(res.CCode) + ".svg"

		list2 = append(list2, res) // note the = instead of :=
	}

	// for _, res := range datas {
	// 	fmt.Println("r", res.IP)
	// 	res.Flag = "https://flagcdn.com/" + strings.ToLower(res.CCode) + ".svg"
	// }
	fmt.Println(result, datas)
	out, _ := json.Marshal(list2)
	return string(out)

}

func InsertData(data model.UpInfo) {
	db := connct()
	// up := model.UpInfo{Os: "s"}
	result := db.Create(&data) // 通过数据的指针来创建
	fmt.Println(result)
}

func UpInsert(data model.UpInfo) {
	db := connct()
	// Update columns to new value on `id` conflict
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "IP"}},                                                                   // key colume
		DoUpdates: clause.AssignmentColumns([]string{"mem_used", "cpu_used", "update_time", "uptime", "platform"}), // column needed to be updated
	}).Create(&data)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age=VALUES(age); MySQL
}
