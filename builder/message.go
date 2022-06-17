package builder

// Message - concrete product
type Message struct {
	Recipient	string `json:"recipient" xml:"recipient"`
	Text 		string `json:"text" xml:"text"`
}

// MessageRepresented - Object representation
type MessageRepresented struct {
	Body 	string
	Format	string
}