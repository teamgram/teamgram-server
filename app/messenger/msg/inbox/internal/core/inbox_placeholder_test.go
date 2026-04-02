package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/messenger/msg/inbox/inbox"
)

func TestInboxPlaceholderVoids(t *testing.T) {
	c := New(context.Background(), nil)

	if result, err := c.InboxSendUserMessageToInboxV2(&inbox.TLInboxSendUserMessageToInboxV2{}); err != nil || result == nil {
		t.Fatalf("expected sendUserMessageToInboxV2 void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxEditMessageToInboxV2(&inbox.TLInboxEditMessageToInboxV2{}); err != nil || result == nil {
		t.Fatalf("expected editMessageToInboxV2 void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxDeleteMessagesToInbox(&inbox.TLInboxDeleteMessagesToInbox{}); err != nil || result == nil {
		t.Fatalf("expected deleteMessagesToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxReadInboxHistory(&inbox.TLInboxReadInboxHistory{}); err != nil || result == nil {
		t.Fatalf("expected readInboxHistory void placeholder, got result=%#v err=%v", result, err)
	}
}
