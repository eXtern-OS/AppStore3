package stats

type ExportedStats struct {
	Downloads int
	Rating    float64
}

type Stats struct {
	Downloads    int
	Rating       float64
	RatingPeriod float64
	Regions      []string
}

func (s *Stats) Export() ExportedStats {
	return ExportedStats{
		Downloads: s.Downloads,
		Rating:    s.Rating,
	}
}
