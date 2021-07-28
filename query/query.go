package query

type Params struct {
	EnableFree         bool `json:"enable_free"`
	EnablePaid         bool `json:"enable_paid"`
	EnableSubscription bool `json:"enable_subscription"`
}

type Query struct {
	Query          string `json:"query"`
	SnapEnabled    bool   `json:"snap_enabled"`
	FlatpakEnabled bool   `json:"flatpak_enabled"`
	Results        int    `json:"results"` // Default 100
	NoCache        bool   `json:"no_cache"`
	Params         Params `json:"params"`
}
