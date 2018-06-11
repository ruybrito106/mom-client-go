package distribution

type Packet struct {
	Type string `json:"type"`
	Message *Message `json:"message"`
}