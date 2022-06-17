package builder

import (
	"encoding/xml"
)

// XMLMessageBuilder is concrete builder
type XMLMessageBuilder struct {
	message Message
}

// SetRecipient - assigns the message recipient
func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

// SetMessage - assigns the message 
func (b *XMLMessageBuilder) SetMessage(text string) MessageBuilder {
	b.message.Text = text
	return b
}

// Build - builds the object as XML rep
func (b *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}

	return &MessageRepresented{
		Body:   string(data),
		Format: "XML",
	},nil
}
