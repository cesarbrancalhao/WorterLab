package main

import (
	"github.com/cesarbrancalhao/WorterLab/api/router"
	"github.com/cesarbrancalhao/WorterLab/internal/config"
	"github.com/cesarbrancalhao/WorterLab/internal/db"
)

func init() {
	config.LoadEnv()
	db.ConnectDB()
}

func main() {
	router.Start()
}
