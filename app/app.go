package app

import (
	"extenos.io/AppStore3/publisher"
	"extenos.io/AppStore3/stats"
)

type App interface {
	Export() ExportedApp
	IsPaid() bool
}

// Package provides version
type Package struct {
	Published int64 `json:"published" bson:"published"`

	Version string `json:"version" bson:"version"`
	URL     string `json:"url" bson:"url"`

	Requirements []string `json:"requirements" bson:"requirements"`
}

// ExportedApp provides structure for unified app
type ExportedApp struct {
	AppType Type `json:"app_type"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Version string `json:"version"`

	StatsAvailable bool `json:"stats_available"`

	Stats stats.ExportedStats `json:"stats"`

	IconURL   string                      `json:"icon_url"`
	HeaderURL string                      `json:"header_url"`
	Publisher publisher.ExportedPublisher `json:"publisher"`

	PackageName string  `json:"package_name" bson:"package_name"`
	Package     Package `json:"package" bson:"package"` // only  available for extern apps
}

func ExportApps(income []App) []ExportedApp {
	var result []ExportedApp

	for _, x := range income {
		result = append(result, x.Export())
	}

	return result
}
