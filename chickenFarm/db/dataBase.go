package db

import (
	"chickenFarm/global"
	"chickenFarm/model"
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"

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
	// db := connct()
	Connct()
	global.DB.AutoMigrate(&model.UpInfo{})

}
func Connct() {
	///www/xj/chickenFarm/db/sqlite.db
	global.DB, _ = gorm.Open(sqlite.Open("../db/sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
}
func GetAllInfo() string {
	// db := connct()
	var datas []model.UpInfo
	result := global.DB.Find(&datas)
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
	// db := connct()
	// up := model.UpInfo{Os: "s"}
	result := global.DB.Create(&data) // 通过数据的指针来创建
	fmt.Println(result)
}

func UpInsert(data model.UpInfo) {
	// Update columns to new value on `id` conflict
	global.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "IP"}},                                                                                       // key colume
		DoUpdates: clause.AssignmentColumns([]string{"mem_used", "cpu_used", "update_time", "uptime", "send_traffic", "recv_traffic"}), // column needed to be updated
	}).Create(&data)
	// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
	// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
	// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age=VALUES(age); MySQL
}

func DeleteOne(condation string, value string, model interface{}) {
	global.DB.Where(condation, value).Delete(&model)
}
