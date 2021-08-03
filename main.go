package main

import (
	"externos.io/AppStore3/apps/extern"
	"externos.io/AppStore3/apps/flatpak"
	"externos.io/AppStore3/apps/flatpak/daemon"
	"externos.io/AppStore3/query"
	"externos.io/AppStore3/search"
	"github.com/davecgh/go-spew/spew"
	"github.com/eXtern-OS/common/config"
	"github.com/eXtern-OS/common/db"
	"time"
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

	extern.Init()

	time.Sleep(1 * time.Minute)

	res := search.Search(query.Query{
		Query:          "Spotify",
		SnapEnabled:    true,
		FlatpakEnabled: true,
		Results:        130,
		NoCache:        false,
		Params: query.Params{
			EnableFree:         true,
			EnablePaid:         true,
			EnableSubscription: true,
		},
	})

	spew.Dump(res)
}
