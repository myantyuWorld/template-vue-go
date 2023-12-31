package main

// Goでモックを作成してテストをする | https://qiita.com/S-Masakatsu/items/2bc751df9583657181e9
// go generate でモックを生成する | https://christina04.hatenablog.com/entry/use-go-generate-when-generating-mock
//　以下をリポジトリなどに書く。| go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
// プロジェクト直下で「go generate ./...」

import (
	"api/pkg/infrastructure"
	"api/pkg/infrastructure/persistence"
	"api/pkg/interfaces/handler"
	"api/pkg/usecase"
	"log"

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
	db, err := infrastructure.ConnRDB()
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

	petUseCase := usecase.NewPetUseCase(
		persistence.NewPetPersistence(),
		persistence.NewSchedulePersistence(),
		persistence.NewConditionPersistence(),
	)
	petHandler := handler.NewPetHandler(petUseCase, *db)

	e := echo.New()
	e.GET("/pet/:id", petHandler.HandlePetGet())

	e.Logger.Fatal(e.Start(":8080"))
}
