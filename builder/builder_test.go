package builder_test

import (
	"testing"

	"github.com/axtoneIO/design-patterns-golang/builder"
)

func TestSender_BuildMessage(t *testing.T) {
	//json := &builder.JSONMessageBuilder{}
	xmlm := &builder.XMLMessageBuilder{}
	sender := &builder.Sender{}
	sender.SetBuilder(xmlm)
	msg, err := sender.BuildMessage("Gopher","Hello World")
	if err != nil {
		t.Fatalf("BuildMessage(): Error has ocurred: %v", err)
	}

	t.Log(msg.Body)
}
