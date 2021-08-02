package extern

import (
	"context"
	"externos.io/AppStore3/query"
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"github.com/eXtern-OS/common/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Search(q query.Query, res chan []app.App, limit int) {
	var resApps []app.App
	dx := utils.SumBtoI(q.Params.EnableFree, q.Params.EnablePaid, q.Params.EnableSubscription)
	if dx > 0 {
		var filter bson.A
		if dx > 2 {
			var arr bson.A
			if q.Params.EnableFree {
				arr = append(arr, bson.M{"payment": 0})
			}
			if q.Params.EnablePaid {
				arr = append(arr, bson.M{"payment": 1})
			}
			if q.Params.EnableSubscription {
				arr = append(arr, bson.M{"payment": 2})
			}
			filter = append(filter, bson.D{{"$or", arr}})
		} else {
			if q.Params.EnableFree {
				filter = append(filter, bson.M{"payment": 0})
			} else if q.Params.EnablePaid {
				filter = append(filter, bson.M{"payment": 1})
			} else if q.Params.EnableSubscription {
				filter = append(filter, bson.M{"payment": 2})
			}
		}

		filter = append(filter, bson.D{
			{"$or", bson.A{
				bson.D{{"name", primitive.Regex{Pattern: q.Query, Options: ""}}},
				bson.D{{"description", primitive.Regex{Pattern: q.Query, Options: ""}}},
			}}})

		cur, err := dbc.FindMany(filter, DatabaseName, CollectionName)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Println("No results found for query:", q)
			} else {
				log.Println("Failed to find data in db", err, "extern.Search")
				go beatrix.SendError("Failed to find data in db: "+err.Error(), "extern.Search")
			}
		} else {
			for i := 0; i < limit && cur.Next(context.TODO()); i++ {
				var a app.Extern

				if err = cur.Decode(&a); err != nil {
					log.Println("Failed to decode document", err, "extern.Search")
					go beatrix.SendError("Failed to decode document: "+err.Error(), "extern.Search")
					i--
				} else {
					resApps = append(resApps, &a)
				}
			}
		}
	}
	res <- resApps
}
