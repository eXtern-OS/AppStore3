package app

import (
	"extenos.io/AppStore3/publisher"
	"extenos.io/AppStore3/stats"
)

/* Sample {
	"flatpakAppId":"com.play0ad.zeroad",
	"name":"0 A.D.",
	"summary": "Real-Time Strategy Game of Ancient Warfare",
	"currentReleaseVersion":"0.0.24",
	"currentReleaseDate":"2021-07-21",
	"iconDesktopUrl":"https://dl.flathub.org/repo/appstream/x86_64/icons/128x128/com.play0ad.zeroad.png",
	"iconMobileUrl":"https://dl.flathub.org/repo/appstream/x86_64/icons/128x128/com.play0ad.zeroad.png",
	"inStoreSinceDate":"2017-04-18T04:14:01Z"
}
*/

type Flatpak struct {
	FlatpakAppId          string `json:"flatpak_app_id" bson:"flatpak_app_id"`
	Name                  string `json:"name" bson:"name"`
	Summary               string `json:"summary" bson:"summary"`
	CurrentReleaseVersion string `json:"current_release_version" bson:"current_release_version"`
	CurrentReleaseDate    string `json:"current_release_date" bson:"current_release_date"`
	IconDesktopUrl        string `json:"icon_desktop_url" bson:"icon_desktop_url"`
	IconMobileUrl         string `json:"icon_mobile_url" bson:"icon_mobile_url"`
	InStoreSinceDate      string `json:"in_store_since_date" bson:"in_store_since_date"`
}

func (f *Flatpak) Export() ExportedApp {
	return ExportedApp{
		AppType:        FlatpakApp,
		Name:           f.Name,
		Description:    f.Summary,
		Version:        f.CurrentReleaseVersion,
		StatsAvailable: false,
		Stats:          stats.ExportedStats{},
		IconURL:        f.IconDesktopUrl,
		HeaderURL:      "",
		Publisher:      publisher.NewPublisher(f.Name),
		PackageName:    f.FlatpakAppId,
	}
}

func (f *Flatpak) IsPaid() bool {
	return false
}
