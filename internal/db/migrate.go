package db

import (
	"github.com/cesarbrancalhao/WorterLab/internal/config"
	"github.com/cesarbrancalhao/WorterLab/internal/model"
)

func init() {
	config.LoadEnv()
	ConnectPG()
}

func main() {
	DB.DisableForeignKeyConstraintWhenMigrating = true
	Migrate()
	DB.DisableForeignKeyConstraintWhenMigrating = false
	Migrate()
}

func Migrate() {
	DB.AutoMigrate(
		&model.User{},
		&model.Word{},
		&model.UserWord{},
		&model.Log{},
	)
}
