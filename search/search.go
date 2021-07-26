package search

import (
	"extenos.io/AppStore3/app"
	"extenos.io/AppStore3/apps/extern"
	"extenos.io/AppStore3/query"
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

	wg.Add(1)
	var exChan = make(chan []app.App, 1)

	res = append(res, exChan)

	go extern.Search(q, exChan, &wg)

	if q.SnapEnabled {
		wg.Add(1)
		var snapChan = make(chan []app.App, 1)
		res = append(res, snapChan)
		// add snap search
	}

	if q.FlatpakEnabled {
		wg.Add(1)
		var flatChan = make(chan []app.App, 1)
		res = append(res, flatChan)
		// add flat search
	}

	wg.Wait()

	var results []app.App

	for _, c := range res {
		results = append(results, <-c...)
	}

	return app.ExportApps(results)
}
