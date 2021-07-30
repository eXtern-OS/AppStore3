package flatpak

import (
	"github.com/eXtern-OS/common/db"
)

var dbc *db.DBClient

func Init() {
	dbc = db.NewClient()
}

const (
	DatabaseName   = "Apps"
	CollectionName = "Flathub"
)
