package main

import (
	"externos.io/AppStore3/apps/flatpak"
	"externos.io/AppStore3/apps/flatpak/daemon"
	"github.com/davecgh/go-spew/spew"
	"github.com/eXtern-OS/common/app"
	"github.com/eXtern-OS/common/config"
	"github.com/eXtern-OS/common/db"
	"sync"
)

type Config struct {
	Mongo string `json:"mongo"`
}

func main() {
	var c Config
	config.ReadConfig(&c)

	db.Init(c.Mongo)

	daemon.Init()

	flatpak.Init()

	res := make(chan []app.App, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go flatpak.Search("Spotify", res, &wg, 5)

	wg.Wait()

	spew.Dump(app.ExportApps(<-res))
}
