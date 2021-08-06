package snap

import (
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"github.com/eXtern-OS/common/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var dbc *db.DBClient

func Init() {
	dbc = db.NewClient()
}

const (
	DatabaseName   = "Apps"
	CollectionName = "SnapCache"
)

type CacheFilter struct {
	Query  string     `bson:"query"`
	Result []app.Snap `bson:"result"`
}

func AddToCache(q string, apps []app.Snap) {
	if err := dbc.InsertData(CacheFilter{
		Query:  q,
		Result: apps,
	}, DatabaseName, CollectionName); err != nil {
		go beatrix.SendError("Failed to insert cache data: "+err.Error(), "snap.AddToCache")
		log.Println("Failed to insert cache data: "+err.Error(), "snap.AddToCache")
	} else {
		cmap.set(q)
	}
}

func LoadFromCache(q string) ([]app.Snap, error) {
	var res CacheFilter
	return res.Result, dbc.FindData(bson.M{"query": q}, &res, DatabaseName, CollectionName)
}
