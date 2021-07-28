package main

import (
	"externos.io/AppStore3/apps/snap"
	"fmt"
	"github.com/eXtern-OS/common/app"
	"sync"
)

type Config struct {
	Mongo string `json:"mongo"`
}

func main() {
	//var c Config
	//config.ReadConfig(&c)

	//db.Init(c.Mongo)

	c := make(chan []app.App, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go snap.Search("spotify", c, &wg, 5)

	wg.Wait()

	fmt.Println(app.ExportApps(<-c))
}
