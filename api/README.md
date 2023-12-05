# 使用ライブラリ
- [echo](https://echo.labstack.com/docs/quick-start)
- [gorm](https://gorm.io/ja_JP/docs/connecting_to_the_database.html)
- [sql-migrate](https://github.com/rubenv/sql-migrate)

# 特記事項
パッケージ構成は自由にしていただいて構いません！

# go generate
## mockを生成したいファイルに以下追記
1. `//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE`
2. PJ直下で、`go generate ./...`

# go test
## -cover | カバレッジ取得
`go test -cover`

## -v | 詳細表示
`go test -v` 

## テストの絞り込み
`go test -run "Test名"