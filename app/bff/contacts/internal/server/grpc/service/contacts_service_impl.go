/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2024 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/contacts/internal/core"
)

// AccountGetContactSignUpNotification
// account.getContactSignUpNotification#9f07c728 = Bool;
func (s *Service) AccountGetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountGetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getContactSignUpNotification - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getContactSignUpNotification - reply: {%s}", r)
	return r, err
}

// AccountSetContactSignUpNotification
// account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
func (s *Service) AccountSetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setContactSignUpNotification - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setContactSignUpNotification - reply: {%s}", r)
	return r, err
}

// ContactsGetContactIDs
// contacts.getContactIDs#7adc669d hash:long = Vector<int>;
func (s *Service) ContactsGetContactIDs(ctx context.Context, request *mtproto.TLContactsGetContactIDs) (*mtproto.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getContactIDs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetContactIDs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getContactIDs - reply: {%s}", r)
	return r, err
}

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (s *Service) ContactsGetStatuses(ctx context.Context, request *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getStatuses - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetStatuses(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getStatuses - reply: {%s}", r)
	return r, err
}

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (s *Service) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getContacts - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getContacts - reply: {%s}", r)
	return r, err
}

// ContactsImportContacts
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (s *Service) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.importContacts - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsImportContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.importContacts - reply: {%s}", r)
	return r, err
}

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (s *Service) ContactsDeleteContacts(ctx context.Context, request *mtproto.TLContactsDeleteContacts) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.deleteContacts - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsDeleteContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.deleteContacts - reply: {%s}", r)
	return r, err
}

// ContactsDeleteByPhones
// contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
func (s *Service) ContactsDeleteByPhones(ctx context.Context, request *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.deleteByPhones - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsDeleteByPhones(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.deleteByPhones - reply: {%s}", r)
	return r, err
}

// ContactsBlock
// contacts.block#2e2e8734 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (s *Service) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.block - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsBlock(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.block - reply: {%s}", r)
	return r, err
}

// ContactsUnblock
// contacts.unblock#b550d328 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (s *Service) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.unblock - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsUnblock(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.unblock - reply: {%s}", r)
	return r, err
}

// ContactsGetBlocked
// contacts.getBlocked#9a868f80 flags:# my_stories_from:flags.0?true offset:int limit:int = contacts.Blocked;
func (s *Service) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getBlocked - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetBlocked(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getBlocked - reply: {%s}", r)
	return r, err
}

// ContactsSearch
// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (s *Service) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.search - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.search - reply: {%s}", r)
	return r, err
}

// ContactsGetTopPeers
// contacts.getTopPeers#973478b6 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true bots_app:flags.16?true offset:int limit:int hash:long = contacts.TopPeers;
func (s *Service) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getTopPeers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetTopPeers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getTopPeers - reply: {%s}", r)
	return r, err
}

// ContactsResetTopPeerRating
// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (s *Service) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resetTopPeerRating - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsResetTopPeerRating(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resetTopPeerRating - reply: {%s}", r)
	return r, err
}

// ContactsResetSaved
// contacts.resetSaved#879537f1 = Bool;
func (s *Service) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resetSaved - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsResetSaved(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resetSaved - reply: {%s}", r)
	return r, err
}

// ContactsGetSaved
// contacts.getSaved#82f1e39f = Vector<SavedContact>;
func (s *Service) ContactsGetSaved(ctx context.Context, request *mtproto.TLContactsGetSaved) (*mtproto.Vector_SavedContact, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getSaved - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetSaved(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getSaved - reply: {%s}", r)
	return r, err
}

// ContactsToggleTopPeers
// contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
func (s *Service) ContactsToggleTopPeers(ctx context.Context, request *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.toggleTopPeers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsToggleTopPeers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.toggleTopPeers - reply: {%s}", r)
	return r, err
}

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (s *Service) ContactsAddContact(ctx context.Context, request *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.addContact - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsAddContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.addContact - reply: {%s}", r)
	return r, err
}

// ContactsAcceptContact
// contacts.acceptContact#f831a20f id:InputUser = Updates;
func (s *Service) ContactsAcceptContact(ctx context.Context, request *mtproto.TLContactsAcceptContact) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.acceptContact - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsAcceptContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.acceptContact - reply: {%s}", r)
	return r, err
}

// ContactsGetLocated
// contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
func (s *Service) ContactsGetLocated(ctx context.Context, request *mtproto.TLContactsGetLocated) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getLocated - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsGetLocated(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getLocated - reply: {%s}", r)
	return r, err
}

// ContactsEditCloseFriends
// contacts.editCloseFriends#ba6705f0 id:Vector<long> = Bool;
func (s *Service) ContactsEditCloseFriends(ctx context.Context, request *mtproto.TLContactsEditCloseFriends) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.editCloseFriends - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsEditCloseFriends(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.editCloseFriends - reply: {%s}", r)
	return r, err
}

// ContactsSetBlocked
// contacts.setBlocked#94c65c76 flags:# my_stories_from:flags.0?true id:Vector<InputPeer> limit:int = Bool;
func (s *Service) ContactsSetBlocked(ctx context.Context, request *mtproto.TLContactsSetBlocked) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.setBlocked - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsSetBlocked(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.setBlocked - reply: {%s}", r)
	return r, err
}
