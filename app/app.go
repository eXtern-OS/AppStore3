package app

import (
	"extenos.io/AppStore3/payment"
	"extenos.io/AppStore3/stats"
)

// App is an app which was developed for extern OS
type App struct {
	Name        string
	Description string

	Latest   Package
	Packages []Package

	Payment payment.Payment
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
}
