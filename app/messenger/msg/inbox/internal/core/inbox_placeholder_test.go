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
	if result, err := c.InboxReadOutboxHistory(&inbox.TLInboxReadOutboxHistory{}); err != nil || result == nil {
		t.Fatalf("expected readOutboxHistory void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxUpdateHistoryReaded(&inbox.TLInboxUpdateHistoryReaded{}); err != nil || result == nil {
		t.Fatalf("expected updateHistoryReaded void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxReadUserMediaUnreadToInbox(&inbox.TLInboxReadUserMediaUnreadToInbox{}); err != nil || result == nil {
		t.Fatalf("expected readUserMediaUnreadToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxReadChatMediaUnreadToInbox(&inbox.TLInboxReadChatMediaUnreadToInbox{}); err != nil || result == nil {
		t.Fatalf("expected readChatMediaUnreadToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxEditUserMessageToInbox(&inbox.TLInboxEditUserMessageToInbox{}); err != nil || result == nil {
		t.Fatalf("expected editUserMessageToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxEditChatMessageToInbox(&inbox.TLInboxEditChatMessageToInbox{}); err != nil || result == nil {
		t.Fatalf("expected editChatMessageToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxUpdatePinnedMessage(&inbox.TLInboxUpdatePinnedMessage{}); err != nil || result == nil {
		t.Fatalf("expected updatePinnedMessage void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxUpdatePinnedMessageV2(&inbox.TLInboxUpdatePinnedMessageV2{}); err != nil || result == nil {
		t.Fatalf("expected updatePinnedMessageV2 void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxUnpinAllMessages(&inbox.TLInboxUnpinAllMessages{}); err != nil || result == nil {
		t.Fatalf("expected unpinAllMessages void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxDeleteUserHistoryToInbox(&inbox.TLInboxDeleteUserHistoryToInbox{}); err != nil || result == nil {
		t.Fatalf("expected deleteUserHistoryToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxDeleteChatHistoryToInbox(&inbox.TLInboxDeleteChatHistoryToInbox{}); err != nil || result == nil {
		t.Fatalf("expected deleteChatHistoryToInbox void placeholder, got result=%#v err=%v", result, err)
	}
	if result, err := c.InboxReadMediaUnreadToInboxV2(&inbox.TLInboxReadMediaUnreadToInboxV2{}); err != nil || result == nil {
		t.Fatalf("expected readMediaUnreadToInboxV2 void placeholder, got result=%#v err=%v", result, err)
	}
}
