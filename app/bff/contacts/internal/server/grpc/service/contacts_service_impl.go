/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
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
	c.Infof("account.getContactSignUpNotification - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountGetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.getContactSignUpNotification - reply: %s", r.DebugString())
	return r, err
}

// AccountSetContactSignUpNotification
// account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
func (s *Service) AccountSetContactSignUpNotification(ctx context.Context, request *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("account.setContactSignUpNotification - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.AccountSetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Infof("account.setContactSignUpNotification - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetContactIDs
// contacts.getContactIDs#7adc669d hash:long = Vector<int>;
func (s *Service) ContactsGetContactIDs(ctx context.Context, request *mtproto.TLContactsGetContactIDs) (*mtproto.Vector_Int, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getContactIDs - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetContactIDs(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getContactIDs - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (s *Service) ContactsGetStatuses(ctx context.Context, request *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getStatuses - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetStatuses(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getStatuses - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (s *Service) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getContacts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetContacts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getContacts - reply: %s", r.DebugString())
	return r, err
}

// ContactsImportContacts
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (s *Service) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.importContacts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsImportContacts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.importContacts - reply: %s", r.DebugString())
	return r, err
}

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (s *Service) ContactsDeleteContacts(ctx context.Context, request *mtproto.TLContactsDeleteContacts) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.deleteContacts - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsDeleteContacts(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.deleteContacts - reply: %s", r.DebugString())
	return r, err
}

// ContactsDeleteByPhones
// contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
func (s *Service) ContactsDeleteByPhones(ctx context.Context, request *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.deleteByPhones - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsDeleteByPhones(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.deleteByPhones - reply: %s", r.DebugString())
	return r, err
}

// ContactsBlock
// contacts.block#68cc1411 id:InputPeer = Bool;
func (s *Service) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.block - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsBlock(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.block - reply: %s", r.DebugString())
	return r, err
}

// ContactsUnblock
// contacts.unblock#bea65d50 id:InputPeer = Bool;
func (s *Service) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.unblock - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsUnblock(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.unblock - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetBlocked
// contacts.getBlocked#f57c350f offset:int limit:int = contacts.Blocked;
func (s *Service) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getBlocked - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetBlocked(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getBlocked - reply: %s", r.DebugString())
	return r, err
}

// ContactsSearch
// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (s *Service) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.search - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsSearch(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.search - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetTopPeers
// contacts.getTopPeers#973478b6 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true offset:int limit:int hash:long = contacts.TopPeers;
func (s *Service) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getTopPeers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetTopPeers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getTopPeers - reply: %s", r.DebugString())
	return r, err
}

// ContactsResetTopPeerRating
// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (s *Service) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.resetTopPeerRating - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsResetTopPeerRating(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.resetTopPeerRating - reply: %s", r.DebugString())
	return r, err
}

// ContactsResetSaved
// contacts.resetSaved#879537f1 = Bool;
func (s *Service) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.resetSaved - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsResetSaved(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.resetSaved - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetSaved
// contacts.getSaved#82f1e39f = Vector<SavedContact>;
func (s *Service) ContactsGetSaved(ctx context.Context, request *mtproto.TLContactsGetSaved) (*mtproto.Vector_SavedContact, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getSaved - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetSaved(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getSaved - reply: %s", r.DebugString())
	return r, err
}

// ContactsToggleTopPeers
// contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
func (s *Service) ContactsToggleTopPeers(ctx context.Context, request *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.toggleTopPeers - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsToggleTopPeers(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.toggleTopPeers - reply: %s", r.DebugString())
	return r, err
}

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (s *Service) ContactsAddContact(ctx context.Context, request *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.addContact - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsAddContact(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.addContact - reply: %s", r.DebugString())
	return r, err
}

// ContactsAcceptContact
// contacts.acceptContact#f831a20f id:InputUser = Updates;
func (s *Service) ContactsAcceptContact(ctx context.Context, request *mtproto.TLContactsAcceptContact) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.acceptContact - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsAcceptContact(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.acceptContact - reply: %s", r.DebugString())
	return r, err
}

// ContactsGetLocated
// contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
func (s *Service) ContactsGetLocated(ctx context.Context, request *mtproto.TLContactsGetLocated) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Infof("contacts.getLocated - metadata: %s, request: %s", c.MD.DebugString(), request.DebugString())

	r, err := c.ContactsGetLocated(request)
	if err != nil {
		return nil, err
	}

	c.Infof("contacts.getLocated - reply: %s", r.DebugString())
	return r, err
}
