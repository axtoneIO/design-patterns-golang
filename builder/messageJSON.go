package builder

import (
	"encoding/json"
)

// JSONMessageBuilder is concrete builder
type JSONMessageBuilder struct {
	message Message
}

// SetRecipient - assigns the message recipient
func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage - assigns the message
func (b *JSONMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// Build - builds the object as JSON rep
func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.message)
	if err != nil {
		return nil, err
	}

	return &MessageRepresented{
		Body:   string(data),
		Format: "JSON",
	}, nil
}
