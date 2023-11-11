package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	UserId   string
	UserName string
}

func seeds(db *gorm.DB) error {

	users := Users{UserId: "user1", UserName: "hogehoge"}
	if err := db.Create(&users).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}

func openConnection() *gorm.DB {
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	dsn := "root:password@tcp(db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func main() {
	db := openConnection()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	// defer db.Close() // seed/sample_users_seed.go:44:14: db.Close undefined (type *gorm.DB has no field or method Close)
}
