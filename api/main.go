package main

import (
	"github.com/labstack/echo/v4"
)

// Go言語：文法基礎まとめ | https://qiita.com/HiromuMasuda0228/items/65b9a593275f769f6b69
func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":8080"))
}
