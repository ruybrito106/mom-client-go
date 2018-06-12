package distribution

type Period struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Config struct {
	Type   string  `json:"type"`
	Period *Period `json:"period"`
}
