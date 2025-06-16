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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/contacts/contacts/contactsservice"

	"github.com/cloudwego/kitex/client"
)

type ContactsClient interface {
	AccountGetContactSignUpNotification(ctx context.Context, in *tg.TLAccountGetContactSignUpNotification) (*tg.Bool, error)
	AccountSetContactSignUpNotification(ctx context.Context, in *tg.TLAccountSetContactSignUpNotification) (*tg.Bool, error)
	ContactsGetContactIDs(ctx context.Context, in *tg.TLContactsGetContactIDs) (*tg.VectorInt, error)
	ContactsGetStatuses(ctx context.Context, in *tg.TLContactsGetStatuses) (*tg.VectorContactStatus, error)
	ContactsGetContacts(ctx context.Context, in *tg.TLContactsGetContacts) (*tg.ContactsContacts, error)
	ContactsImportContacts(ctx context.Context, in *tg.TLContactsImportContacts) (*tg.ContactsImportedContacts, error)
	ContactsDeleteContacts(ctx context.Context, in *tg.TLContactsDeleteContacts) (*tg.Updates, error)
	ContactsDeleteByPhones(ctx context.Context, in *tg.TLContactsDeleteByPhones) (*tg.Bool, error)
	ContactsBlock(ctx context.Context, in *tg.TLContactsBlock) (*tg.Bool, error)
	ContactsUnblock(ctx context.Context, in *tg.TLContactsUnblock) (*tg.Bool, error)
	ContactsGetBlocked(ctx context.Context, in *tg.TLContactsGetBlocked) (*tg.ContactsBlocked, error)
	ContactsSearch(ctx context.Context, in *tg.TLContactsSearch) (*tg.ContactsFound, error)
	ContactsGetTopPeers(ctx context.Context, in *tg.TLContactsGetTopPeers) (*tg.ContactsTopPeers, error)
	ContactsResetTopPeerRating(ctx context.Context, in *tg.TLContactsResetTopPeerRating) (*tg.Bool, error)
	ContactsResetSaved(ctx context.Context, in *tg.TLContactsResetSaved) (*tg.Bool, error)
	ContactsGetSaved(ctx context.Context, in *tg.TLContactsGetSaved) (*tg.VectorSavedContact, error)
	ContactsToggleTopPeers(ctx context.Context, in *tg.TLContactsToggleTopPeers) (*tg.Bool, error)
	ContactsAddContact(ctx context.Context, in *tg.TLContactsAddContact) (*tg.Updates, error)
	ContactsAcceptContact(ctx context.Context, in *tg.TLContactsAcceptContact) (*tg.Updates, error)
	ContactsGetLocated(ctx context.Context, in *tg.TLContactsGetLocated) (*tg.Updates, error)
	ContactsEditCloseFriends(ctx context.Context, in *tg.TLContactsEditCloseFriends) (*tg.Bool, error)
	ContactsSetBlocked(ctx context.Context, in *tg.TLContactsSetBlocked) (*tg.Bool, error)
}

type defaultContactsClient struct {
	cli client.Client
}

func NewContactsClient(cli client.Client) ContactsClient {
	return &defaultContactsClient{
		cli: cli,
	}
}

// AccountGetContactSignUpNotification
// account.getContactSignUpNotification#9f07c728 = Bool;
func (m *defaultContactsClient) AccountGetContactSignUpNotification(ctx context.Context, in *tg.TLAccountGetContactSignUpNotification) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.AccountGetContactSignUpNotification(ctx, in)
}

// AccountSetContactSignUpNotification
// account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
func (m *defaultContactsClient) AccountSetContactSignUpNotification(ctx context.Context, in *tg.TLAccountSetContactSignUpNotification) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.AccountSetContactSignUpNotification(ctx, in)
}

// ContactsGetContactIDs
// contacts.getContactIDs#7adc669d hash:long = Vector<int>;
func (m *defaultContactsClient) ContactsGetContactIDs(ctx context.Context, in *tg.TLContactsGetContactIDs) (*tg.VectorInt, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetContactIDs(ctx, in)
}

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (m *defaultContactsClient) ContactsGetStatuses(ctx context.Context, in *tg.TLContactsGetStatuses) (*tg.VectorContactStatus, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetStatuses(ctx, in)
}

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (m *defaultContactsClient) ContactsGetContacts(ctx context.Context, in *tg.TLContactsGetContacts) (*tg.ContactsContacts, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetContacts(ctx, in)
}

// ContactsImportContacts
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (m *defaultContactsClient) ContactsImportContacts(ctx context.Context, in *tg.TLContactsImportContacts) (*tg.ContactsImportedContacts, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsImportContacts(ctx, in)
}

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (m *defaultContactsClient) ContactsDeleteContacts(ctx context.Context, in *tg.TLContactsDeleteContacts) (*tg.Updates, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsDeleteContacts(ctx, in)
}

// ContactsDeleteByPhones
// contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
func (m *defaultContactsClient) ContactsDeleteByPhones(ctx context.Context, in *tg.TLContactsDeleteByPhones) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsDeleteByPhones(ctx, in)
}

// ContactsBlock
// contacts.block#2e2e8734 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (m *defaultContactsClient) ContactsBlock(ctx context.Context, in *tg.TLContactsBlock) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsBlock(ctx, in)
}

// ContactsUnblock
// contacts.unblock#b550d328 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (m *defaultContactsClient) ContactsUnblock(ctx context.Context, in *tg.TLContactsUnblock) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsUnblock(ctx, in)
}

// ContactsGetBlocked
// contacts.getBlocked#9a868f80 flags:# my_stories_from:flags.0?true offset:int limit:int = contacts.Blocked;
func (m *defaultContactsClient) ContactsGetBlocked(ctx context.Context, in *tg.TLContactsGetBlocked) (*tg.ContactsBlocked, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetBlocked(ctx, in)
}

// ContactsSearch
// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (m *defaultContactsClient) ContactsSearch(ctx context.Context, in *tg.TLContactsSearch) (*tg.ContactsFound, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsSearch(ctx, in)
}

// ContactsGetTopPeers
// contacts.getTopPeers#973478b6 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true bots_app:flags.16?true offset:int limit:int hash:long = contacts.TopPeers;
func (m *defaultContactsClient) ContactsGetTopPeers(ctx context.Context, in *tg.TLContactsGetTopPeers) (*tg.ContactsTopPeers, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetTopPeers(ctx, in)
}

// ContactsResetTopPeerRating
// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (m *defaultContactsClient) ContactsResetTopPeerRating(ctx context.Context, in *tg.TLContactsResetTopPeerRating) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsResetTopPeerRating(ctx, in)
}

// ContactsResetSaved
// contacts.resetSaved#879537f1 = Bool;
func (m *defaultContactsClient) ContactsResetSaved(ctx context.Context, in *tg.TLContactsResetSaved) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsResetSaved(ctx, in)
}

// ContactsGetSaved
// contacts.getSaved#82f1e39f = Vector<SavedContact>;
func (m *defaultContactsClient) ContactsGetSaved(ctx context.Context, in *tg.TLContactsGetSaved) (*tg.VectorSavedContact, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetSaved(ctx, in)
}

// ContactsToggleTopPeers
// contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
func (m *defaultContactsClient) ContactsToggleTopPeers(ctx context.Context, in *tg.TLContactsToggleTopPeers) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsToggleTopPeers(ctx, in)
}

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (m *defaultContactsClient) ContactsAddContact(ctx context.Context, in *tg.TLContactsAddContact) (*tg.Updates, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsAddContact(ctx, in)
}

// ContactsAcceptContact
// contacts.acceptContact#f831a20f id:InputUser = Updates;
func (m *defaultContactsClient) ContactsAcceptContact(ctx context.Context, in *tg.TLContactsAcceptContact) (*tg.Updates, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsAcceptContact(ctx, in)
}

// ContactsGetLocated
// contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
func (m *defaultContactsClient) ContactsGetLocated(ctx context.Context, in *tg.TLContactsGetLocated) (*tg.Updates, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsGetLocated(ctx, in)
}

// ContactsEditCloseFriends
// contacts.editCloseFriends#ba6705f0 id:Vector<long> = Bool;
func (m *defaultContactsClient) ContactsEditCloseFriends(ctx context.Context, in *tg.TLContactsEditCloseFriends) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsEditCloseFriends(ctx, in)
}

// ContactsSetBlocked
// contacts.setBlocked#94c65c76 flags:# my_stories_from:flags.0?true id:Vector<InputPeer> limit:int = Bool;
func (m *defaultContactsClient) ContactsSetBlocked(ctx context.Context, in *tg.TLContactsSetBlocked) (*tg.Bool, error) {
	cli := contactsservice.NewRPCContactsClient(m.cli)
	return cli.ContactsSetBlocked(ctx, in)
}
