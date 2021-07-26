package app

import (
	"extenos.io/AppStore3/payment"
	"extenos.io/AppStore3/publisher"
	"extenos.io/AppStore3/stats"
)

// Extern is an app which was developed for extern OS
type Extern struct {
	Name        string
	Description string

	Icon   string `json:"icon" bson:"icon"`
	Header string `json:"header" bson:"header"`

	Publisher publisher.Publisher `json:"publisher" bson:"publisher"`

	Latest   Package   `json:"latest" bson:"latest"`
	Packages []Package `json:"packages" bson:"packages"`

	Payment payment.Payment `json:"payment" bson:"payment"`

	StatsAvailable bool `json:"stats_available" bson:"stats_available"`

	Stats stats.Stats `json:"stats" bson:"stats"`
}

func (e *Extern) Export() ExportedApp {
	return ExportedApp{
		AppType:        ExternApp,
		Name:           e.Name,
		Description:    e.Description,
		Version:        e.Latest.Version,
		StatsAvailable: e.StatsAvailable,
		Stats:          e.Stats.Export(),
		IconURL:        e.Icon,
		HeaderURL:      e.Header,
		Publisher:      e.Publisher.Export(),
		Package:        e.Latest,
	}
}

func (e *Extern) IsPaid() bool {
	return e.Payment != payment.Free
}
