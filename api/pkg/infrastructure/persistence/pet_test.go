package persistence_test

import (
	"api/pkg/infrastructure"
	"api/pkg/infrastructure/persistence"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func OpenDb() (*infrastructure.RDB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dbCfg := infrastructure.NewMySQLConfig(host, 3306, dbname, user, pass)
	db, err := infrastructure.ConnRDB(dbCfg)

	return db, err
}

// Goでモックを作成してテストをする | https://qiita.com/S-Masakatsu/items/2bc751df9583657181e9
func TestGet_PetIdで検索できない場合(t *testing.T) {
	db, err := OpenDb()
	//　モックの生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	output, err := persistence.NewPetPersistence().Get(db, 999)
	log.Print(output)
	require.Equal(t, errors.New("検索できませんでした"), err)
}

func TestGet_Petの属性が正しく検索できている(t *testing.T) {
	db, err := OpenDb()
	output, err := persistence.NewPetPersistence().Get(db, 1)
	require.NoError(t, err)
	require.Equal(t, uint64(1), output.ID)
	require.Equal(t, "natsu", output.Name)
	require.Equal(t, uint64(4000), output.NowWeight)
}
