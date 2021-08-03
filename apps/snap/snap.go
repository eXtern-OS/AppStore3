package snap

import (
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"log"
	"sync"
)

func Search(q string, res chan *app.ExportedApp, limit int, wg *sync.WaitGroup) {
	d, err := getData(q)

	if err != nil {
		go beatrix.SendError("Failed to get data: "+err.Error(), "apps.snap.Search")
		go log.Println("Failed to get data: " + err.Error())
	} else {
		snapApps, err := parseData(d)

		if err != nil {
			go beatrix.SendError("Failed to parse data: "+err.Error(), "apps.snap.Search")
			go log.Println(err)
		} else {
			if len(snapApps) > limit {
				snapApps = snapApps[:(limit - 1)]
			}

			for _, x := range snapApps {
				m := x.Export()
				res <- &m
			}
		}
	}

	wg.Done()

	return
}
