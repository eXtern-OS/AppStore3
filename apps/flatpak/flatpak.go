package flatpak

import (
	"context"
	"externos.io/AppStore3/apps/flatpak/status"
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)


// Search searches in our DB, where we store data from flathub
func Search(q string, res chan *app.ExportedApp, limit int, wg *sync.WaitGroup) {
	// We create a filter which matches either name or description
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"name", primitive.Regex{Pattern: q, Options: ""}}},
			bson.D{{"summary", primitive.Regex{Pattern: q, Options: ""}}},
		}},
	}

	// Locking mutex to avoid that unfortunate situation when we are searching while updating db
	// Consider using traditional for ;; lock
	status.Mutex.Lock()
	cur, err := dbc.FindMany(filter, DatabaseName, CollectionName)
	if err != nil && err != mongo.ErrNoDocuments {
		go beatrix.SendError("Failed to find documents: "+err.Error(), "flatpak.Search")
		log.Println("Failed to find documents: " + err.Error())
	} else if err == mongo.ErrNoDocuments {
		log.Println("No entries for the query:", q)
	} else {
		// See? He did the trick again! (see extern and snap search)
		for i := 0; cur != nil && cur.Next(context.TODO()) && i < limit && i < status.ReasonableLimit; i++ {
			var f app.Flatpak
			if err = cur.Decode(&f); err != nil {
				go beatrix.SendError("Failed to decode value: "+err.Error(), "flatpak.Search")
				log.Println("Failed to decode value: " + err.Error())
				i--
			} else {
				m := f.Export()
				res <- &m
			}
		}
	}
	// unlocking mutex
	status.Mutex.Unlock()
	wg.Done()
}
