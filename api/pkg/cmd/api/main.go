package main

import (
	"api/pkg/cmd/api/app"
	"api/pkg/infrastructure/repository"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// Go言語：文法基礎まとめ | https://qiita.com/HiromuMasuda0228/items/65b9a593275f769f6b69
func main() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")

	dbCfg := repository.NewMySQLConfig(host, 3306, dbname, user, pass)
	db, err := repository.ConnRDB(dbCfg)

	container := app.Inject(db)

	//
	// [参考]log.Fatalなどは、main.goなどプログラムエントリーでのみ使用する
	//
	// func Fatal(v ...any) {
	// 	std.Output(2, fmt.Sprint(v...))
	// 	os.Exit(1)
	// }
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	log.Printf("db :: %#v\n", db)

	e.Logger.Fatal(e.Start(":8080"))
}
