package builder

import "log"

// Sender - Director
type Sender struct {
	builder MessageBuilder
}

// SetBuilder - assigns the constructor
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

// BuildMessage a concrete message via MessageBuilder 
func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	s.builder.SetRecipient(recipient)
	s.builder.SetMessage(message)
	m, err := s.builder.Build()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return m,nil
}