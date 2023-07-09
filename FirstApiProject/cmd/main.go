package main

import (
	"github.com/soheilkhaledabdi/FirstApi/api"
	"github.com/soheilkhaledabdi/FirstApi/config"
	"github.com/soheilkhaledabdi/FirstApi/data/db"
	"github.com/soheilkhaledabdi/FirstApi/pkg/logging"
)

func main() {

	cfg := config.GetConfig()

	logger := logging.NewLogger(cfg)

	err := db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres , logging.Startup , err.Error() , nil)
	}

	api.InitServer()
}