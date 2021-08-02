package search

import (
	"externos.io/AppStore3/apps/extern"
	"externos.io/AppStore3/apps/flatpak"
	"externos.io/AppStore3/apps/snap"
	"externos.io/AppStore3/query"
	"github.com/eXtern-OS/common/app"
	"github.com/eXtern-OS/common/utils"
)

func Search(q query.Query) []app.ExportedApp {
	t := 1 + utils.SumBtoI(q.SnapEnabled, q.FlatpakEnabled)

	targets := q.Results / t

	res := make(chan []app.App, 3)

	go extern.Search(q, res, targets)

	if q.SnapEnabled {
		go snap.Search(q.Query, res, targets)
	}

	if q.FlatpakEnabled {
		go flatpak.Search(q.Query, res, targets)
	}

	var results []app.ExportedApp

	for i := 0; i < 3; i++ {
		results = append(results, app.ExportApps(<-res)...)
	}
	close(res)

	return results
}
