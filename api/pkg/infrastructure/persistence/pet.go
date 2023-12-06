package persistence

import (
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
	dbmodel "api/pkg/infrastructure/model"
	"errors"
)

// infrastructure層は、DBアクセスなどの技術的関心を記述します。
// この層はdomain層に依存しています。純粋なレイヤードアーキテクチャの場合、
// 依存の向きがdomain → infrastructureですが、今回はDDDを取り込んだ設計になるので、依存の向きが逆転します。
// そのためinfrastructure層はdomain層のrepositoryで定義したインタフェースを実装します。
type petPersistence struct{}

// Get implements repository.PetRepository.
func (*petPersistence) Get(db *infrastructure.RDB, petId uint64) (*dbmodel.Pets, error) {
	var pets dbmodel.Pets
	result := db.Find(&pets, petId)
	if result.RowsAffected == 0 {
		return nil, errors.New("検索できませんでした")
	}

	return &pets, nil
}

func NewPetPersistence() repository.PetRepository {
	return &petPersistence{}
}
