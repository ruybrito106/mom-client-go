package distribution

type Packet struct {
	Type    string   `json:"type"`
	Config  *Config  `json:"config"`
	Message *Message `json:"message"`
}
