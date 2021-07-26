package search

import "extenos.io/AppStore3/app"

type Query struct {
	Query          string `json:"query"`
	SnapEnabled    bool   `json:"snap_enabled"`
	FlatpakEnabled bool   `json:"flatpak_enabled"`
	Results        int    // Default 100
}

func bint(income bool) int {
	if income {
		return 1
	} else {
		return 0
	}
}

func Search(q Query) []app.ExportedApp {
	t := 1 + bint(q.SnapEnabled) + bint(q.FlatpakEnabled)

	targets := q.Results / t

}
