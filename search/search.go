package search

import (
	"externos.io/AppStore3/apps/extern"
	"externos.io/AppStore3/apps/snap"
	"externos.io/AppStore3/query"
	"github.com/eXtern-OS/common/app"
	"sync"
)

func bint(income bool) int {
	if income {
		return 1
	} else {
		return 0
	}
}

func Search(q query.Query) []app.ExportedApp {
	t := 1 + bint(q.SnapEnabled) + bint(q.FlatpakEnabled)

	targets := q.Results / t

	var wg sync.WaitGroup

	var res []chan []app.App

	wg.Add(t)
	var exChan = make(chan []app.App, 1)

	res = append(res, exChan)

	go extern.Search(q, exChan, &wg, targets)

	if q.SnapEnabled {
		var snapChan = make(chan []app.App, 1)
		res = append(res, snapChan)
		go snap.Search(q.Query, snapChan, &wg, targets)
	}

	if q.FlatpakEnabled {
		var flatChan = make(chan []app.App, 1)
		res = append(res, flatChan)
		// add flat search
	}

	wg.Wait()

	var results []app.ExportedApp

	for _, c := range res {
		results = append(results, app.ExportApps(<-c)...)
	}

	return results
}
