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

func Search(q string, res chan []app.App, wg *sync.WaitGroup, limit int) {
	var results []app.App

	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"name", primitive.Regex{Pattern: q, Options: ""}}},
			bson.D{{"summary", primitive.Regex{Pattern: q, Options: ""}}},
		}},
	}

	status.Mutex.Lock()
	cur, err := dbc.FindMany(filter, DatabaseName, CollectionName)
	if err != nil && err != mongo.ErrNoDocuments {
		go beatrix.SendError("Failed to find documents: "+err.Error(), "flatpak.Search")
		log.Println("Failed to find documents: " + err.Error())
	} else if err == mongo.ErrNoDocuments {
		log.Println("No entries for the query:", q)
	} else {
		for i := 0; cur != nil && cur.Next(context.TODO()) && i < limit; i++ {
			var f app.Flatpak
			if err = cur.Decode(&f); err != nil {
				go beatrix.SendError("Failed to decode value: "+err.Error(), "flatpak.Search")
				log.Println("Failed to decode value: " + err.Error())
			} else {
				results = append(results, &f)
			}
		}
	}
	status.Mutex.Unlock()

	res <- results
	wg.Done()
}
