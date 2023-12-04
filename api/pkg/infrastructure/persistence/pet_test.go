package persistence_test

import (
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
	"api/pkg/infrastructure/persistence"
	mock_repository "api/pkg/mock/repository"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dbCfg := infrastructure.NewMySQLConfig(host, 3306, dbname, user, pass)
	db, err := infrastructure.ConnRDB(dbCfg)

	resp := &dbmodel.Pets{}

	// モックを呼び出すための Controller 生成

	//　モックの生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// テストに呼ばれるべきメソッドと引数・戻り値を指定
	pr := mock_repository.NewMockPetRepository(ctrl)
	pr.EXPECT().Get(db, 1).Return(resp, nil)

	output, err := persistence.NewPetPersistence().Get(db, 1)
	require.NoError(t, err)
	require.Equal(t, resp, output)
}
