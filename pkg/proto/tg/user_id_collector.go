package tg

type UserIDCollector struct {
	ids  []int64
	seen map[int64]struct{}
}

func CollectUserIDsFromUpdates(updates *Updates) []int64 {
	c := newUserIDCollector()
	if updates == nil {
		return c.ids
	}
	if full, ok := updates.ToUpdates(); ok {
		c.collectUpdates(full.Updates)
	}
	if combined, ok := updates.ToUpdatesCombined(); ok {
		c.collectUpdates(combined.Updates)
	}
	return c.ids
}

func CollectUserIDsFromDifference(diff *UpdatesDifference) []int64 {
	c := newUserIDCollector()
	if diff == nil {
		return c.ids
	}
	if full, ok := diff.ToUpdatesDifference(); ok {
		c.collectMessages(full.NewMessages)
		c.collectUpdates(full.OtherUpdates)
	}
	if slice, ok := diff.ToUpdatesDifferenceSlice(); ok {
		c.collectMessages(slice.NewMessages)
		c.collectUpdates(slice.OtherUpdates)
	}
	return c.ids
}

func CollectChatIDsFromDifference(diff *UpdatesDifference) []int64 {
	c := newChatIDCollector()
	if diff == nil {
		return c.ids
	}
	if full, ok := diff.ToUpdatesDifference(); ok {
		c.collectMessages(full.NewMessages)
		c.collectUpdates(full.OtherUpdates)
	}
	if slice, ok := diff.ToUpdatesDifferenceSlice(); ok {
		c.collectMessages(slice.NewMessages)
		c.collectUpdates(slice.OtherUpdates)
	}
	return c.ids
}

func CollectUserIDsFromMessagesMessages(messages *MessagesMessages) []int64 {
	c := newUserIDCollector()
	if messages == nil {
		return c.ids
	}
	if full, ok := messages.ToMessagesMessages(); ok {
		c.collectMessages(full.Messages)
	}
	if slice, ok := messages.ToMessagesMessagesSlice(); ok {
		c.collectMessages(slice.Messages)
	}
	if channel, ok := messages.ToMessagesChannelMessages(); ok {
		c.collectMessages(channel.Messages)
	}
	return c.ids
}

func CollectUserIDsFromMessage(message MessageClazz) []int64 {
	c := newUserIDCollector()
	c.collectMessage(message)
	return c.ids
}

func CollectUserIDsFromUpdate(update UpdateClazz) []int64 {
	c := newUserIDCollector()
	c.collectUpdate(update)
	return c.ids
}

func newUserIDCollector() *UserIDCollector {
	return &UserIDCollector{seen: make(map[int64]struct{})}
}

func (c *UserIDCollector) add(id int64) {
	if id <= 0 {
		return
	}
	if _, ok := c.seen[id]; ok {
		return
	}
	c.seen[id] = struct{}{}
	c.ids = append(c.ids, id)
}

func (c *UserIDCollector) collectUpdates(updates []UpdateClazz) {
	for _, update := range updates {
		c.collectUpdate(update)
	}
}

func (c *UserIDCollector) collectMessages(messages []MessageClazz) {
	for _, message := range messages {
		c.collectMessage(message)
	}
}

func (c *UserIDCollector) collectPeer(peer PeerClazz) {
	if p, ok := peer.(*TLPeerUser); ok {
		c.add(p.UserId)
	}
}

func (c *UserIDCollector) collectMessage(message MessageClazz) {
	switch m := message.(type) {
	case *TLMessage:
		c.collectPeer(m.FromId)
		c.collectPeer(m.PeerId)
		c.collectPeer(m.SavedPeerId)
		c.collectMessageFwdHeader(m.FwdFrom)
		c.collectMessageMedia(m.Media)
		c.collectInt64Ptr(m.ViaBotId)
		c.collectInt64Ptr(m.ViaBusinessBotId)
		c.collectMessageReplyHeader(m.ReplyTo)
	case *TLMessageEmpty:
		c.collectPeer(m.PeerId)
	case *TLMessageService:
		c.collectPeer(m.FromId)
		c.collectPeer(m.PeerId)
		c.collectPeer(m.SavedPeerId)
		c.collectMessageReplyHeader(m.ReplyTo)
	}
}

func (c *UserIDCollector) collectMessageMedia(media MessageMediaClazz) {
	if m, ok := media.(*TLMessageMediaContact); ok {
		c.add(m.UserId)
	}
}

func (c *UserIDCollector) collectMessageFwdHeader(header MessageFwdHeaderClazz) {
	if header != nil {
		c.collectPeer(header.FromId)
		c.collectPeer(header.SavedFromPeer)
		c.collectPeer(header.SavedFromId)
	}
}

func (c *UserIDCollector) collectMessageReplyHeader(header MessageReplyHeaderClazz) {
	switch h := header.(type) {
	case *TLMessageReplyHeader:
		c.collectPeer(h.ReplyToPeerId)
		c.collectMessageFwdHeader(h.ReplyFrom)
	case *TLMessageReplyStoryHeader:
		c.collectPeer(h.Peer)
	}
}

func (c *UserIDCollector) collectInt64Ptr(id *int64) {
	if id != nil {
		c.add(*id)
	}
}

func (c *UserIDCollector) collectUpdate(update UpdateClazz) {
	switch u := update.(type) {
	case *TLUpdateEditMessage:
		c.collectMessage(u.Message)
	case *TLUpdateNewMessage:
		c.collectMessage(u.Message)
	case *TLUpdateNewChannelMessage:
		c.collectMessage(u.Message)
	case *TLUpdateEditChannelMessage:
		c.collectMessage(u.Message)
	case *TLUpdateNewScheduledMessage:
		c.collectMessage(u.Message)
	case *TLUpdateReadHistoryInbox:
		c.collectPeer(u.Peer)
	case *TLUpdateReadHistoryOutbox:
		c.collectPeer(u.Peer)
	case *TLUpdateDraftMessage:
		c.collectPeer(u.Peer)
		c.collectPeer(u.SavedPeerId)
	case *TLUpdatePeerSettings:
		c.collectPeer(u.Peer)
	case *TLUpdatePeerHistoryTTL:
		c.collectPeer(u.Peer)
	case *TLUpdatePinnedMessages:
		c.collectPeer(u.Peer)
	case *TLUpdateDeleteScheduledMessages:
		c.collectPeer(u.Peer)
	case *TLUpdateChatDefaultBannedRights:
		c.collectPeer(u.Peer)
	case *TLUpdateChatUserTyping:
		c.collectPeer(u.FromId)
	case *TLUpdateUserTyping:
		c.add(u.UserId)
	case *TLUpdateUserStatus:
		c.add(u.UserId)
	case *TLUpdateUserName:
		c.add(u.UserId)
	case *TLUpdateUserPhone:
		c.add(u.UserId)
	case *TLUpdateUser:
		c.add(u.UserId)
	case *TLUpdateUserEmojiStatus:
		c.add(u.UserId)
	case *TLUpdateChatParticipantAdd:
		c.add(u.UserId)
		c.add(u.InviterId)
	case *TLUpdateChatParticipantDelete:
		c.add(u.UserId)
	case *TLUpdateChatParticipantAdmin:
		c.add(u.UserId)
	case *TLUpdateChatParticipant:
		c.add(u.UserId)
		c.collectChatParticipant(u.PrevParticipant)
		c.collectChatParticipant(u.NewParticipant)
	case *TLUpdateChatParticipantRank:
		c.add(u.UserId)
	case *TLUpdateChatParticipants:
		c.collectChatParticipants(u.Participants)
	}
}

func (c *UserIDCollector) collectChatParticipants(participants ChatParticipantsClazz) {
	if p, ok := participants.(*TLChatParticipants); ok {
		for _, participant := range p.Participants {
			c.collectChatParticipant(participant)
		}
	}
}

func (c *UserIDCollector) collectChatParticipant(participant ChatParticipantClazz) {
	switch p := participant.(type) {
	case *TLChatParticipant:
		c.add(p.UserId)
		c.add(p.InviterId)
	case *TLChatParticipantCreator:
		c.add(p.UserId)
	case *TLChatParticipantAdmin:
		c.add(p.UserId)
		c.add(p.InviterId)
	}
}

type ChatIDCollector struct {
	ids  []int64
	seen map[int64]struct{}
}

func newChatIDCollector() *ChatIDCollector {
	return &ChatIDCollector{seen: make(map[int64]struct{})}
}

func (c *ChatIDCollector) add(id int64) {
	if id <= 0 {
		return
	}
	if _, ok := c.seen[id]; ok {
		return
	}
	c.seen[id] = struct{}{}
	c.ids = append(c.ids, id)
}

func (c *ChatIDCollector) collectUpdates(updates []UpdateClazz) {
	for _, update := range updates {
		c.collectUpdate(update)
	}
}

func (c *ChatIDCollector) collectMessages(messages []MessageClazz) {
	for _, message := range messages {
		c.collectMessage(message)
	}
}

func (c *ChatIDCollector) collectPeer(peer PeerClazz) {
	if p, ok := peer.(*TLPeerChat); ok {
		c.add(p.ChatId)
	}
}

func (c *ChatIDCollector) collectMessage(message MessageClazz) {
	switch m := message.(type) {
	case *TLMessage:
		c.collectPeer(m.FromId)
		c.collectPeer(m.PeerId)
		c.collectPeer(m.SavedPeerId)
		c.collectMessageFwdHeader(m.FwdFrom)
		c.collectMessageReplyHeader(m.ReplyTo)
	case *TLMessageEmpty:
		c.collectPeer(m.PeerId)
	case *TLMessageService:
		c.collectPeer(m.FromId)
		c.collectPeer(m.PeerId)
		c.collectPeer(m.SavedPeerId)
		c.collectMessageReplyHeader(m.ReplyTo)
	}
}

func (c *ChatIDCollector) collectMessageFwdHeader(header MessageFwdHeaderClazz) {
	if header != nil {
		c.collectPeer(header.FromId)
		c.collectPeer(header.SavedFromPeer)
		c.collectPeer(header.SavedFromId)
	}
}

func (c *ChatIDCollector) collectMessageReplyHeader(header MessageReplyHeaderClazz) {
	switch h := header.(type) {
	case *TLMessageReplyHeader:
		c.collectPeer(h.ReplyToPeerId)
		c.collectMessageFwdHeader(h.ReplyFrom)
	case *TLMessageReplyStoryHeader:
		c.collectPeer(h.Peer)
	}
}

func (c *ChatIDCollector) collectUpdate(update UpdateClazz) {
	switch u := update.(type) {
	case *TLUpdateEditMessage:
		c.collectMessage(u.Message)
	case *TLUpdateNewMessage:
		c.collectMessage(u.Message)
	case *TLUpdateNewChannelMessage:
		c.collectMessage(u.Message)
	case *TLUpdateEditChannelMessage:
		c.collectMessage(u.Message)
	case *TLUpdateNewScheduledMessage:
		c.collectMessage(u.Message)
	case *TLUpdateReadHistoryInbox:
		c.collectPeer(u.Peer)
	case *TLUpdateReadHistoryOutbox:
		c.collectPeer(u.Peer)
	case *TLUpdateDraftMessage:
		c.collectPeer(u.Peer)
		c.collectPeer(u.SavedPeerId)
	case *TLUpdatePeerSettings:
		c.collectPeer(u.Peer)
	case *TLUpdatePeerHistoryTTL:
		c.collectPeer(u.Peer)
	case *TLUpdatePinnedMessages:
		c.collectPeer(u.Peer)
	case *TLUpdateDeleteScheduledMessages:
		c.collectPeer(u.Peer)
	case *TLUpdateChatDefaultBannedRights:
		c.collectPeer(u.Peer)
	case *TLUpdateChatUserTyping:
		c.add(u.ChatId)
	case *TLUpdateChatParticipantAdd:
		c.add(u.ChatId)
	case *TLUpdateChatParticipantDelete:
		c.add(u.ChatId)
	case *TLUpdateChatParticipantAdmin:
		c.add(u.ChatId)
	case *TLUpdateChatParticipant:
		c.add(u.ChatId)
	case *TLUpdateChatParticipantRank:
		c.add(u.ChatId)
	case *TLUpdateChat:
		c.add(u.ChatId)
	case *TLUpdateChatParticipants:
		if p, ok := u.Participants.(*TLChatParticipants); ok {
			c.add(p.ChatId)
		}
	}
}
