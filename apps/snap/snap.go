package snap

import (
	beatrix "github.com/eXtern-OS/Beatrix"
	"github.com/eXtern-OS/common/app"
	"sync"
)

func Search(q string, res chan []app.App, wg *sync.WaitGroup, limit int) {
	d, err := getData(q)

	var apps []app.App

	if err != nil {
		go beatrix.SendError("Failed to get data: "+err.Error(), "apps.snap.Search")
	} else {
		snapApps, err := parseData(d)

		if err != nil {
			go beatrix.SendError("Failed to parse data: "+err.Error(), "apps.snap.Search")
		} else {
			if len(snapApps) > limit {
				snapApps = snapApps[:(limit - 1)]
			}

			for _, x := range snapApps {
				apps = append(apps, &x)
			}
		}
	}

	res <- apps
	wg.Done()
	return
}
