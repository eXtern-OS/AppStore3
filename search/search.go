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

func Search(q query.Query) []app.ExportedApp {
	t := 1 + utils.SumBtoI(q.SnapEnabled, q.FlatpakEnabled)

	targets := q.Results / t

	var wg, wge sync.WaitGroup
	wg.Add(t)
	wge.Add(1)

	e := newEliminator()

	res := make(chan *app.ExportedApp)

	go extern.Search(q, res, targets, &wg)

	if q.SnapEnabled {
		go snap.Search(q.Query, res, targets, &wg)
	}

	if q.FlatpakEnabled {
		go flatpak.Search(q.Query, res, targets, &wg)
	}

	go e.start(res, &wge)

	wg.Wait()

	close(res)

	wge.Wait()

	return e.get()
}
