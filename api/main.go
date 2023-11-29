package main

import (
	"api/pkg/infrastructure"
	"api/pkg/infrastructure/persistence"
	"api/pkg/interfaces/handler"
	"api/pkg/usecase"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

// Go言語：文法基礎まとめ | https://qiita.com/HiromuMasuda0228/items/65b9a593275f769f6b69
// --- 今すぐ「レイヤードアーキテクチャ+DDD」を理解しよう。（golang）| https://qiita.com/tono-maron/items/345c433b86f74d314c8d
// --- 若干理解が浅いため、実例が欲しいので、上記記事を参考に実装してみた。実装後、理解が深まるはずなので、サンプルを改めて目を通してみる
// |
// LayeredArchitecture
// ├── cmd
// │   └── api
// │       └── main.go
// ├── domain
// │   └── repository
// |   |   └── user_repository.go
// |   └── user.go
// ├── config //DBの起動など
// │   └── database.go
// ├── interfaces
// │   └── handler
// │   |   └── user.go
// |   └── response
// │       └── response.go
// ├── infrastructure
// │   └── persistence
// │       └── user.go
// └── usecase
//	   └── user.go

func main() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")

	dbCfg := infrastructure.NewMySQLConfig(host, 3306, dbname, user, pass)
	db, err := infrastructure.ConnRDB(dbCfg)
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
	log.Printf("db :: %#v\n", db)

	petPersistence := persistence.NewPetPersistence()
	petUseCase := usecase.NewPetUseCase(petPersistence)
	petHandler := handler.NewPetHandler(petUseCase, *db)

	e := echo.New()
	e.GET("/pet", petHandler.HandlePetGet())

	e.Logger.Fatal(e.Start(":8080"))
}
