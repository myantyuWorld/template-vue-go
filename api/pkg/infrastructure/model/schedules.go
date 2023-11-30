package dbmodel

import (
	"api/pkg/infrastructure"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Schedules struct {
	PetId     uint64
	Title     string
	Date      time.Time
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (mdl *Schedules) TableName() string {
	return "schedules"
}

func seeds(db *infrastructure.RDB) error {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			schedule := Schedules{
				PetId:     uint64(i + 1),
				Title:     "サンプル" + strconv.Itoa(i),
				Date:      time.Now(),
				Location:  "サンプルペットショップ" + strconv.Itoa(j),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := db.Create(&schedule).Error; err != nil {
				fmt.Printf("%+v", err)
			}
		}
	}

	return nil
}

func main() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")

	dbCfg := infrastructure.NewMySQLConfig(host, 3306, dbname, user, pass)
	db, err := infrastructure.ConnRDB(dbCfg)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
}
