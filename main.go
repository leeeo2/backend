package main

import (
	"flag"
	"github.com/leeeo2/backend/pkg/common/config"
	"github.com/leeeo2/backend/pkg/common/logger"
	"github.com/leeeo2/backend/pkg/model"
	"github.com/leeeo2/backend/pkg/router"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./etc/backend.yml", "config path")
	flag.Parse()

	// load config
	if err := config.Init(configPath); err != nil {
		panic(err)
	}

	// setup logger
	if err := logger.Setup(&config.Global().Log); err != nil {
		panic(err)
	}

	// setup database
	if err := model.Setup(&config.Global().Database, &config.Global().Log, &config.Global().GormLog); err != nil {
		panic(err)
	}

	// setup router
	if err := router.Setup(&config.Global().Server); err != nil {
		panic(err)
	}
}
