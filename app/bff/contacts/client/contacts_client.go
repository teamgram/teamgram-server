/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package contactsclient

import (
	"context"

	"github.com/teamgram/proto/mtproto"

	"github.com/zeromicro/go-zero/zrpc"
)

var _ *mtproto.Bool

type ContactsClient interface {
	AccountGetContactSignUpNotification(ctx context.Context, in *mtproto.TLAccountGetContactSignUpNotification) (*mtproto.Bool, error)
	AccountSetContactSignUpNotification(ctx context.Context, in *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error)
	ContactsGetContactIDs(ctx context.Context, in *mtproto.TLContactsGetContactIDs) (*mtproto.Vector_Int, error)
	ContactsGetStatuses(ctx context.Context, in *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error)
	ContactsGetContacts(ctx context.Context, in *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error)
	ContactsImportContacts(ctx context.Context, in *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error)
	ContactsDeleteContacts(ctx context.Context, in *mtproto.TLContactsDeleteContacts) (*mtproto.Updates, error)
	ContactsDeleteByPhones(ctx context.Context, in *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error)
	ContactsBlock(ctx context.Context, in *mtproto.TLContactsBlock) (*mtproto.Bool, error)
	ContactsUnblock(ctx context.Context, in *mtproto.TLContactsUnblock) (*mtproto.Bool, error)
	ContactsGetBlocked(ctx context.Context, in *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error)
	ContactsSearch(ctx context.Context, in *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error)
	ContactsGetTopPeers(ctx context.Context, in *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error)
	ContactsResetTopPeerRating(ctx context.Context, in *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error)
	ContactsResetSaved(ctx context.Context, in *mtproto.TLContactsResetSaved) (*mtproto.Bool, error)
	ContactsGetSaved(ctx context.Context, in *mtproto.TLContactsGetSaved) (*mtproto.Vector_SavedContact, error)
	ContactsToggleTopPeers(ctx context.Context, in *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error)
	ContactsAddContact(ctx context.Context, in *mtproto.TLContactsAddContact) (*mtproto.Updates, error)
	ContactsAcceptContact(ctx context.Context, in *mtproto.TLContactsAcceptContact) (*mtproto.Updates, error)
	ContactsGetLocated(ctx context.Context, in *mtproto.TLContactsGetLocated) (*mtproto.Updates, error)
	ContactsEditCloseFriends(ctx context.Context, in *mtproto.TLContactsEditCloseFriends) (*mtproto.Bool, error)
	ContactsSetBlocked(ctx context.Context, in *mtproto.TLContactsSetBlocked) (*mtproto.Bool, error)
}

type defaultContactsClient struct {
	cli zrpc.Client
}

func NewContactsClient(cli zrpc.Client) ContactsClient {
	return &defaultContactsClient{
		cli: cli,
	}
}

// AccountGetContactSignUpNotification
// account.getContactSignUpNotification#9f07c728 = Bool;
func (m *defaultContactsClient) AccountGetContactSignUpNotification(ctx context.Context, in *mtproto.TLAccountGetContactSignUpNotification) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.AccountGetContactSignUpNotification(ctx, in)
}

// AccountSetContactSignUpNotification
// account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
func (m *defaultContactsClient) AccountSetContactSignUpNotification(ctx context.Context, in *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.AccountSetContactSignUpNotification(ctx, in)
}

// ContactsGetContactIDs
// contacts.getContactIDs#7adc669d hash:long = Vector<int>;
func (m *defaultContactsClient) ContactsGetContactIDs(ctx context.Context, in *mtproto.TLContactsGetContactIDs) (*mtproto.Vector_Int, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetContactIDs(ctx, in)
}

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (m *defaultContactsClient) ContactsGetStatuses(ctx context.Context, in *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetStatuses(ctx, in)
}

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (m *defaultContactsClient) ContactsGetContacts(ctx context.Context, in *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetContacts(ctx, in)
}

// ContactsImportContacts
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (m *defaultContactsClient) ContactsImportContacts(ctx context.Context, in *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsImportContacts(ctx, in)
}

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (m *defaultContactsClient) ContactsDeleteContacts(ctx context.Context, in *mtproto.TLContactsDeleteContacts) (*mtproto.Updates, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsDeleteContacts(ctx, in)
}

// ContactsDeleteByPhones
// contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
func (m *defaultContactsClient) ContactsDeleteByPhones(ctx context.Context, in *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsDeleteByPhones(ctx, in)
}

// ContactsBlock
// contacts.block#2e2e8734 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (m *defaultContactsClient) ContactsBlock(ctx context.Context, in *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsBlock(ctx, in)
}

// ContactsUnblock
// contacts.unblock#b550d328 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (m *defaultContactsClient) ContactsUnblock(ctx context.Context, in *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsUnblock(ctx, in)
}

// ContactsGetBlocked
// contacts.getBlocked#9a868f80 flags:# my_stories_from:flags.0?true offset:int limit:int = contacts.Blocked;
func (m *defaultContactsClient) ContactsGetBlocked(ctx context.Context, in *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetBlocked(ctx, in)
}

// ContactsSearch
// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (m *defaultContactsClient) ContactsSearch(ctx context.Context, in *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsSearch(ctx, in)
}

// ContactsGetTopPeers
// contacts.getTopPeers#973478b6 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true bots_app:flags.16?true offset:int limit:int hash:long = contacts.TopPeers;
func (m *defaultContactsClient) ContactsGetTopPeers(ctx context.Context, in *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetTopPeers(ctx, in)
}

// ContactsResetTopPeerRating
// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (m *defaultContactsClient) ContactsResetTopPeerRating(ctx context.Context, in *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsResetTopPeerRating(ctx, in)
}

// ContactsResetSaved
// contacts.resetSaved#879537f1 = Bool;
func (m *defaultContactsClient) ContactsResetSaved(ctx context.Context, in *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsResetSaved(ctx, in)
}

// ContactsGetSaved
// contacts.getSaved#82f1e39f = Vector<SavedContact>;
func (m *defaultContactsClient) ContactsGetSaved(ctx context.Context, in *mtproto.TLContactsGetSaved) (*mtproto.Vector_SavedContact, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetSaved(ctx, in)
}

// ContactsToggleTopPeers
// contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
func (m *defaultContactsClient) ContactsToggleTopPeers(ctx context.Context, in *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsToggleTopPeers(ctx, in)
}

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (m *defaultContactsClient) ContactsAddContact(ctx context.Context, in *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsAddContact(ctx, in)
}

// ContactsAcceptContact
// contacts.acceptContact#f831a20f id:InputUser = Updates;
func (m *defaultContactsClient) ContactsAcceptContact(ctx context.Context, in *mtproto.TLContactsAcceptContact) (*mtproto.Updates, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsAcceptContact(ctx, in)
}

// ContactsGetLocated
// contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
func (m *defaultContactsClient) ContactsGetLocated(ctx context.Context, in *mtproto.TLContactsGetLocated) (*mtproto.Updates, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsGetLocated(ctx, in)
}

// ContactsEditCloseFriends
// contacts.editCloseFriends#ba6705f0 id:Vector<long> = Bool;
func (m *defaultContactsClient) ContactsEditCloseFriends(ctx context.Context, in *mtproto.TLContactsEditCloseFriends) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsEditCloseFriends(ctx, in)
}

// ContactsSetBlocked
// contacts.setBlocked#94c65c76 flags:# my_stories_from:flags.0?true id:Vector<InputPeer> limit:int = Bool;
func (m *defaultContactsClient) ContactsSetBlocked(ctx context.Context, in *mtproto.TLContactsSetBlocked) (*mtproto.Bool, error) {
	client := mtproto.NewRPCContactsClient(m.cli.Conn())
	return client.ContactsSetBlocked(ctx, in)
}
