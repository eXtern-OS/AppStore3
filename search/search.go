package search

import (
	"externos.io/AppStore3/apps/extern"
	"externos.io/AppStore3/apps/flatpak"
	"externos.io/AppStore3/apps/snap"
	"externos.io/AppStore3/query"
	"github.com/eXtern-OS/common/app"
	"github.com/eXtern-OS/common/utils"
	"sync"
)


// Search consumes query and produces apps 
func Search(q query.Query) []app.ExportedApp {
	// t is the number of target groups
	t := 1 + utils.SumBtoI(q.SnapEnabled, q.FlatpakEnabled)

	// targets is the number that represents number of apps in each group
	targets := q.Results / t

	// wge for eliminator, wg for search
	var wg, wge sync.WaitGroup
	// wg has to wait until all daemons finish, wge is allways one
	wg.Add(t)
	wge.Add(1)

	e := newEliminator()

	res := make(chan *app.ExportedApp)

	go extern.Search(q, res, targets, &wg)

	// starting chosen daemons
	if q.SnapEnabled {
		go snap.Search(q.Query, res, targets, &wg)
	}

	if q.FlatpakEnabled {
		go flatpak.Search(q.Query, res, targets, &wg)
	}

	// eXtern apps are ALWAYS included
	go e.start(res, &wge)

	wg.Wait()

	// closing channel so that eliminator can start working
	close(res)

	// waiting for results
	wge.Wait()

	// getting results from eliminator
	return e.get()
}
