package snap

import (
	"fmt"
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"log"
	"sync"
)

func Search(q string, res chan []app.App, wg *sync.WaitGroup, limit int) {
	d, err := getData(q)

	log.Println("Got data")

	var apps []app.App

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

			log.Println("Trimmed len")

			fmt.Println("apps", snapApps)

			for _, x := range snapApps {
				apps = append(apps, &x)
			}
		}
	}

	res <- apps
	wg.Done()
	return
}
