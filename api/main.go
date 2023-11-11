package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openConnection() *gorm.DB {
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	} else {
		log.Print("connection successed!")
	}
	return db
}

// Go言語：文法基礎まとめ | https://qiita.com/HiromuMasuda0228/items/65b9a593275f769f6b69
func main() {
	e := echo.New()

	db := openConnection()
	log.Printf("db :: %#v\n", db)

	e.Logger.Fatal(e.Start(":8080"))
}
