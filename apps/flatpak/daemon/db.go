// again pretty common file
package daemon

import (
	"github.com/eXtern-OS/common/db"
	"go.mongodb.org/mongo-driver/bson"
)

var dbc *db.DBClient

func Init() {
	dbc = db.NewClient()
	go StartDaemon()
}

const (
	DatabaseName   = "Apps"
	CollectionName = "Flathub"
)

func deleteEverything() error {
	return dbc.DeleteItems(bson.M{}, DatabaseName, CollectionName)
}
