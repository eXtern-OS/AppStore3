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
	"sync"
)

func Search(q query.Query, res chan *app.ExportedApp, limit int, wg *sync.WaitGroup) {
	dx := utils.SumBtoI(q.Params.EnableFree, q.Params.EnablePaid, q.Params.EnableSubscription)
	if dx > 0 {
		var filter = bson.M{}
		var f1 = bson.M{}
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
			f1["$or"] = arr
		} else {
			if q.Params.EnableFree {
				f1["payment"] = 0
			} else if q.Params.EnablePaid {
				f1["payment"] = 1
			} else if q.Params.EnableSubscription {
				f1["payment"] = 2
			}
		}

		filter["$and"] = bson.A{f1, bson.D{
			{"$or", bson.A{
				bson.D{{"name", primitive.Regex{Pattern: q.Query, Options: ""}}},
				bson.D{{"description", primitive.Regex{Pattern: q.Query, Options: ""}}},
			}}}}

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
					m := a.Export()
					res <- &m
				}
			}
		}
	}
	//log.Println("Extern finished")
	wg.Done()
}
