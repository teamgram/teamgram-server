package dependencies

import (
	"reflect"
	"sort"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type DependencySet struct {
	UserIDs    []int64
	ChatIDs    []int64
	ChannelIDs []int64
}

func CollectUpdates(updates []tg.UpdateClazz) DependencySet {
	c := newCollector()
	c.collectUpdates(updates)
	return DependencySet{
		UserIDs:    sortedIDs(c.users),
		ChatIDs:    sortedIDs(c.chats),
		ChannelIDs: sortedIDs(c.channels),
	}
}

type collector struct {
	users    map[int64]struct{}
	chats    map[int64]struct{}
	channels map[int64]struct{}
}

func newCollector() *collector {
	return &collector{
		users:    make(map[int64]struct{}),
		chats:    make(map[int64]struct{}),
		channels: make(map[int64]struct{}),
	}
}

func sortedIDs(ids map[int64]struct{}) []int64 {
	if len(ids) == 0 {
		return nil
	}
	out := make([]int64, 0, len(ids))
	for id := range ids {
		out = append(out, id)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

func addID(ids map[int64]struct{}, id int64) {
	if id <= 0 {
		return
	}
	ids[id] = struct{}{}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func (c *collector) addUser(id int64) {
	addID(c.users, id)
}

func (c *collector) addChat(id int64) {
	addID(c.chats, id)
}

func (c *collector) addChannel(id int64) {
	addID(c.channels, id)
}

func (c *collector) collectUpdates(updates []tg.UpdateClazz) {
	for _, update := range updates {
		c.collectUpdate(update)
	}
}

func (c *collector) collectUpdate(update tg.UpdateClazz) {
	if isNil(update) {
		return
	}

	switch u := update.(type) {
	case *tg.TLUpdateNewMessage:
		c.collectMessage(u.Message)
	case *tg.TLUpdateEditMessage:
		c.collectMessage(u.Message)
	case *tg.TLUpdateNewChannelMessage:
		c.collectMessage(u.Message)
	case *tg.TLUpdateEditChannelMessage:
		c.collectMessage(u.Message)
	case *tg.TLUpdateNewScheduledMessage:
		c.collectMessage(u.Message)
	case *tg.TLUpdateReadHistoryInbox:
		c.collectPeer(u.Peer)
	case *tg.TLUpdateReadHistoryOutbox:
		c.collectPeer(u.Peer)
	case *tg.TLUpdateUserTyping:
		c.addUser(u.UserId)
	case *tg.TLUpdateChatUserTyping:
		c.addChat(u.ChatId)
		c.collectPeer(u.FromId)
	case *tg.TLUpdateChannelUserTyping:
		c.addChannel(u.ChannelId)
		c.collectPeer(u.FromId)
	case *tg.TLUpdateChatParticipants:
		c.collectChatParticipants(u.Participants)
	case *tg.TLUpdateChat:
		c.addChat(u.ChatId)
	case *tg.TLUpdateChatParticipantAdd:
		c.addChat(u.ChatId)
		c.addUser(u.UserId)
		c.addUser(u.InviterId)
	case *tg.TLUpdateChatParticipantDelete:
		c.addChat(u.ChatId)
		c.addUser(u.UserId)
	case *tg.TLUpdateChatParticipantAdmin:
		c.addChat(u.ChatId)
		c.addUser(u.UserId)
	case *tg.TLUpdateChatParticipant:
		c.addChat(u.ChatId)
		c.addUser(u.ActorId)
		c.addUser(u.UserId)
		c.collectChatParticipant(u.PrevParticipant)
		c.collectChatParticipant(u.NewParticipant)
	case *tg.TLUpdatePeerSettings:
		c.collectPeer(u.Peer)
	case *tg.TLUpdateNotifySettings:
		c.collectNotifyPeer(u.Peer)
	case *tg.TLUpdateDraftMessage:
		c.collectPeer(u.Peer)
		c.collectPeer(u.SavedPeerId)
	case *tg.TLUpdateDialogPinned:
		c.collectDialogPeer(u.Peer)
	case *tg.TLUpdatePinnedDialogs:
		c.collectDialogPeers(u.Order)
	case *tg.TLUpdateFolderPeers:
		c.collectFolderPeers(u.FolderPeers)
	case *tg.TLUpdatePeerHistoryTTL:
		c.collectPeer(u.Peer)
	case *tg.TLUpdatePeerWallpaper:
		c.collectPeer(u.Peer)
	case *tg.TLUpdateSavedDialogPinned:
		c.collectDialogPeer(u.Peer)
	case *tg.TLUpdatePinnedSavedDialogs:
		c.collectDialogPeers(u.Order)
	case *tg.TLUpdateDialogUnreadMark:
		c.collectDialogPeer(u.Peer)
		c.collectPeer(u.SavedPeerId)
	case *tg.TLUpdatePeerBlocked:
		c.collectPeer(u.PeerId)
	case *tg.TLUpdateChatDefaultBannedRights:
		c.collectPeer(u.Peer)
	case *tg.TLUpdateChannelTooLong:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateChannel:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateReadChannelInbox:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateReadChannelOutbox:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateDeleteChannelMessages:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateChannelMessageViews:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateChannelWebPage:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateChannelMessageForwards:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateReadChannelDiscussionInbox:
		c.addChannel(u.ChannelId)
		c.collectInt64PtrAsChannel(u.BroadcastId)
	case *tg.TLUpdateReadChannelDiscussionOutbox:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdatePinnedChannelMessages:
		c.addChannel(u.ChannelId)
	case *tg.TLUpdateChannelParticipant:
		c.addChannel(u.ChannelId)
		c.addUser(u.ActorId)
		c.addUser(u.UserId)
	}
}

func (c *collector) collectMessage(message tg.MessageClazz) {
	if isNil(message) {
		return
	}

	switch m := message.(type) {
	case *tg.TLMessage:
		c.collectPeer(m.FromId)
		c.collectPeer(m.PeerId)
		c.collectPeer(m.SavedPeerId)
		c.collectMessageFwdHeader(m.FwdFrom)
		c.collectInt64PtrAsUser(m.ViaBotId)
		c.collectInt64PtrAsUser(m.ViaBusinessBotId)
		c.collectMessageReplyHeader(m.ReplyTo)
		c.collectMessageMedia(m.Media)
	case *tg.TLMessageEmpty:
		c.collectPeer(m.PeerId)
	case *tg.TLMessageService:
		c.collectPeer(m.FromId)
		c.collectPeer(m.PeerId)
		c.collectPeer(m.SavedPeerId)
		c.collectMessageReplyHeader(m.ReplyTo)
		c.collectMessageAction(m.Action)
	}
}

func (c *collector) collectPeer(peer tg.PeerClazz) {
	if isNil(peer) {
		return
	}

	switch p := peer.(type) {
	case *tg.TLPeerUser:
		c.addUser(p.UserId)
	case *tg.TLPeerChat:
		c.addChat(p.ChatId)
	case *tg.TLPeerChannel:
		c.addChannel(p.ChannelId)
	}
}

func (c *collector) collectNotifyPeer(peer tg.NotifyPeerClazz) {
	if isNil(peer) {
		return
	}

	if p, ok := peer.(*tg.TLNotifyPeer); ok {
		c.collectPeer(p.Peer)
	}
}

func (c *collector) collectDialogPeer(peer tg.DialogPeerClazz) {
	if isNil(peer) {
		return
	}

	if p, ok := peer.(*tg.TLDialogPeer); ok {
		c.collectPeer(p.Peer)
	}
}

func (c *collector) collectDialogPeers(peers []tg.DialogPeerClazz) {
	for _, peer := range peers {
		c.collectDialogPeer(peer)
	}
}

func (c *collector) collectFolderPeers(peers []tg.FolderPeerClazz) {
	for _, peer := range peers {
		if peer != nil {
			c.collectPeer(peer.Peer)
		}
	}
}

func (c *collector) collectMessageAction(action tg.MessageActionClazz) {
	if isNil(action) {
		return
	}

	switch a := action.(type) {
	case *tg.TLMessageActionChatCreate:
		c.collectUserIDs(a.Users)
	case *tg.TLMessageActionChatAddUser:
		c.collectUserIDs(a.Users)
	case *tg.TLMessageActionChatDeleteUser:
		c.addUser(a.UserId)
	case *tg.TLMessageActionChatJoinedByLink:
		c.addUser(a.InviterId)
	case *tg.TLMessageActionChatMigrateTo:
		c.addChannel(a.ChannelId)
	case *tg.TLMessageActionChannelMigrateFrom:
		c.addChat(a.ChatId)
	}
}

func (c *collector) collectChatParticipants(participants tg.ChatParticipantsClazz) {
	if isNil(participants) {
		return
	}

	switch p := participants.(type) {
	case *tg.TLChatParticipants:
		c.addChat(p.ChatId)
		for _, participant := range p.Participants {
			c.collectChatParticipant(participant)
		}
	case *tg.TLChatParticipantsForbidden:
		c.addChat(p.ChatId)
		c.collectChatParticipant(p.SelfParticipant)
	}
}

func (c *collector) collectChatParticipant(participant tg.ChatParticipantClazz) {
	if isNil(participant) {
		return
	}

	switch p := participant.(type) {
	case *tg.TLChatParticipant:
		c.addUser(p.UserId)
		c.addUser(p.InviterId)
	case *tg.TLChatParticipantCreator:
		c.addUser(p.UserId)
	case *tg.TLChatParticipantAdmin:
		c.addUser(p.UserId)
		c.addUser(p.InviterId)
	}
}

func (c *collector) collectMessageFwdHeader(header tg.MessageFwdHeaderClazz) {
	if header == nil {
		return
	}
	c.collectPeer(header.FromId)
	c.collectPeer(header.SavedFromPeer)
	c.collectPeer(header.SavedFromId)
}

func (c *collector) collectMessageReplyHeader(header tg.MessageReplyHeaderClazz) {
	if isNil(header) {
		return
	}

	switch h := header.(type) {
	case *tg.TLMessageReplyHeader:
		c.collectPeer(h.ReplyToPeerId)
		c.collectMessageFwdHeader(h.ReplyFrom)
	case *tg.TLMessageReplyStoryHeader:
		c.collectPeer(h.Peer)
	}
}

func (c *collector) collectMessageMedia(media tg.MessageMediaClazz) {
	if isNil(media) {
		return
	}

	if m, ok := media.(*tg.TLMessageMediaContact); ok {
		c.addUser(m.UserId)
	}
}

func (c *collector) collectUserIDs(ids []int64) {
	for _, id := range ids {
		c.addUser(id)
	}
}

func (c *collector) collectInt64PtrAsUser(id *int64) {
	if id != nil {
		c.addUser(*id)
	}
}

func (c *collector) collectInt64PtrAsChannel(id *int64) {
	if id != nil {
		c.addChannel(*id)
	}
}
