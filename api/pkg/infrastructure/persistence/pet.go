package persistence

import (
	"api/pkg/domain/model"
	"api/pkg/domain/repository"
	"api/pkg/infrastructure"
)

// infrastructure層は、DBアクセスなどの技術的関心を記述します。
// この層はdomain層に依存しています。純粋なレイヤードアーキテクチャの場合、
// 依存の向きがdomain → infrastructureですが、今回はDDDを取り込んだ設計になるので、依存の向きが逆転します。
// そのためinfrastructure層はdomain層のrepositoryで定義したインタフェースを実装します。
type petPersistence struct{}

// Get implements repository.PetRepository.
func (*petPersistence) Get(db *infrastructure.RDB, petId uint64) (*model.PetSummary, error) {
	panic("unimplemented")
}

func NewPetPersistence() repository.PetRepository {
	return &petPersistence{}
}
