package main

import (
	"heyChat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/heyChat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("fail to connect to databse")
	}

	// 迁移schema
	db.AutoMigrate(&models.UserBasic{})

	// create
	// user := &models.UserBasic{
	// 	Name:          "王五",
	// 	LoginTime:     time.Now(),
	// 	HeartbeatTime: time.Now(),
	// 	LogOutTime:    time.Now(),
	// }
	// // user.Name = "王五"
	// db.Create(user)

	// read
	// fmt.Println(db.First(user, 1))

	// update
	// db.Model(user).Update("PassWord", "1234")
}
