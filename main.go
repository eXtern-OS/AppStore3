package main

import (
	"externos.io/AppStore3/apps/flatpak/daemon"
	"github.com/eXtern-OS/common/config"
	"github.com/eXtern-OS/common/db"
)

type Config struct {
	Mongo string `json:"mongo"`
}

func main() {
	var c Config
	config.ReadConfig(&c)

	db.Init(c.Mongo)

	daemon.Init()

	daemon.StartDaemon()
}
