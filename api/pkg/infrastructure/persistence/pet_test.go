package persistence_test

import (
	"api/pkg/infrastructure"
	"api/pkg/infrastructure/persistence"
	"errors"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// Goでモックを作成してテストをする | https://qiita.com/S-Masakatsu/items/2bc751df9583657181e9
func TestGet_PetIdで検索できない場合(t *testing.T) {
	db, err := infrastructure.ConnRDB()
	//　モックの生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	output, err := persistence.NewPetPersistence().Get(db, 999)
	log.Print(output)
	require.Equal(t, errors.New("検索できませんでした"), err)
}

func TestGet_Petの属性が正しく検索できている(t *testing.T) {
	db, err := infrastructure.ConnRDB()
	output, err := persistence.NewPetPersistence().Get(db, 1)
	require.NoError(t, err)
	require.Equal(t, uint64(1), output.ID)
	require.Equal(t, "natsu", output.Name)
	require.Equal(t, uint64(4000), output.NowWeight)
}
