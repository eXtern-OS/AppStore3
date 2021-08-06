package main

import (
	"externos.io/AppStore3/apps/extern"
	"externos.io/AppStore3/apps/flatpak"
	"externos.io/AppStore3/apps/flatpak/daemon"
	"externos.io/AppStore3/apps/snap"
	"externos.io/AppStore3/server"
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/config"
	"github.com/eXtern-OS/common/db"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Config struct {
	Mongo            string `json:"mongo"`
	BeatrixToken     string `json:"beatrix_token"`
	BeatrixChannelID string `json:"beatrix_channel_id"`
}

func main() {
	var c Config
	config.ReadConfig(&c)

	db.Init(c.Mongo)
	daemon.Init()
	flatpak.Init()
	extern.Init()
	snap.Init()
	beatrix.Init("AppStore3", c.BeatrixToken, c.BeatrixChannelID)

	time.Sleep(10 * time.Second)

	r := gin.Default()
	server.Init(r)
	log.Panicln(r.Run())
}
