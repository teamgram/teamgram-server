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

	"github.com/teamgram/proto/v2/tg"
	"github.com/teamgram/teamgram-server/v2/app/bff/contacts/internal/core"
)

// AccountGetContactSignUpNotification
// account.getContactSignUpNotification#9f07c728 = Bool;
func (s *Service) AccountGetContactSignUpNotification(ctx context.Context, request *tg.TLAccountGetContactSignUpNotification) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getContactSignUpNotification - metadata: {}, request: {%v}", request)

	r, err := c.AccountGetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getContactSignUpNotification - reply: {%v}", r)
	return r, err
}

// AccountSetContactSignUpNotification
// account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
func (s *Service) AccountSetContactSignUpNotification(ctx context.Context, request *tg.TLAccountSetContactSignUpNotification) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setContactSignUpNotification - metadata: {}, request: {%v}", request)

	r, err := c.AccountSetContactSignUpNotification(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setContactSignUpNotification - reply: {%v}", r)
	return r, err
}

// ContactsGetContactIDs
// contacts.getContactIDs#7adc669d hash:long = Vector<int>;
func (s *Service) ContactsGetContactIDs(ctx context.Context, request *tg.TLContactsGetContactIDs) (*tg.VectorInt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getContactIDs - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetContactIDs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getContactIDs - reply: {%v}", r)
	return r, err
}

// ContactsGetStatuses
// contacts.getStatuses#c4a353ee = Vector<ContactStatus>;
func (s *Service) ContactsGetStatuses(ctx context.Context, request *tg.TLContactsGetStatuses) (*tg.VectorContactStatus, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getStatuses - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetStatuses(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getStatuses - reply: {%v}", r)
	return r, err
}

// ContactsGetContacts
// contacts.getContacts#5dd69e12 hash:long = contacts.Contacts;
func (s *Service) ContactsGetContacts(ctx context.Context, request *tg.TLContactsGetContacts) (*tg.ContactsContacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getContacts - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getContacts - reply: {%v}", r)
	return r, err
}

// ContactsImportContacts
// contacts.importContacts#2c800be5 contacts:Vector<InputContact> = contacts.ImportedContacts;
func (s *Service) ContactsImportContacts(ctx context.Context, request *tg.TLContactsImportContacts) (*tg.ContactsImportedContacts, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.importContacts - metadata: {}, request: {%v}", request)

	r, err := c.ContactsImportContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.importContacts - reply: {%v}", r)
	return r, err
}

// ContactsDeleteContacts
// contacts.deleteContacts#96a0e00 id:Vector<InputUser> = Updates;
func (s *Service) ContactsDeleteContacts(ctx context.Context, request *tg.TLContactsDeleteContacts) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.deleteContacts - metadata: {}, request: {%v}", request)

	r, err := c.ContactsDeleteContacts(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.deleteContacts - reply: {%v}", r)
	return r, err
}

// ContactsDeleteByPhones
// contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
func (s *Service) ContactsDeleteByPhones(ctx context.Context, request *tg.TLContactsDeleteByPhones) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.deleteByPhones - metadata: {}, request: {%v}", request)

	r, err := c.ContactsDeleteByPhones(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.deleteByPhones - reply: {%v}", r)
	return r, err
}

// ContactsBlock
// contacts.block#2e2e8734 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (s *Service) ContactsBlock(ctx context.Context, request *tg.TLContactsBlock) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.block - metadata: {}, request: {%v}", request)

	r, err := c.ContactsBlock(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.block - reply: {%v}", r)
	return r, err
}

// ContactsUnblock
// contacts.unblock#b550d328 flags:# my_stories_from:flags.0?true id:InputPeer = Bool;
func (s *Service) ContactsUnblock(ctx context.Context, request *tg.TLContactsUnblock) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.unblock - metadata: {}, request: {%v}", request)

	r, err := c.ContactsUnblock(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.unblock - reply: {%v}", r)
	return r, err
}

// ContactsGetBlocked
// contacts.getBlocked#9a868f80 flags:# my_stories_from:flags.0?true offset:int limit:int = contacts.Blocked;
func (s *Service) ContactsGetBlocked(ctx context.Context, request *tg.TLContactsGetBlocked) (*tg.ContactsBlocked, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getBlocked - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetBlocked(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getBlocked - reply: {%v}", r)
	return r, err
}

// ContactsSearch
// contacts.search#11f812d8 q:string limit:int = contacts.Found;
func (s *Service) ContactsSearch(ctx context.Context, request *tg.TLContactsSearch) (*tg.ContactsFound, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.search - metadata: {}, request: {%v}", request)

	r, err := c.ContactsSearch(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.search - reply: {%v}", r)
	return r, err
}

// ContactsGetTopPeers
// contacts.getTopPeers#973478b6 flags:# correspondents:flags.0?true bots_pm:flags.1?true bots_inline:flags.2?true phone_calls:flags.3?true forward_users:flags.4?true forward_chats:flags.5?true groups:flags.10?true channels:flags.15?true bots_app:flags.16?true offset:int limit:int hash:long = contacts.TopPeers;
func (s *Service) ContactsGetTopPeers(ctx context.Context, request *tg.TLContactsGetTopPeers) (*tg.ContactsTopPeers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getTopPeers - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetTopPeers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getTopPeers - reply: {%v}", r)
	return r, err
}

// ContactsResetTopPeerRating
// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (s *Service) ContactsResetTopPeerRating(ctx context.Context, request *tg.TLContactsResetTopPeerRating) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resetTopPeerRating - metadata: {}, request: {%v}", request)

	r, err := c.ContactsResetTopPeerRating(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resetTopPeerRating - reply: {%v}", r)
	return r, err
}

// ContactsResetSaved
// contacts.resetSaved#879537f1 = Bool;
func (s *Service) ContactsResetSaved(ctx context.Context, request *tg.TLContactsResetSaved) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.resetSaved - metadata: {}, request: {%v}", request)

	r, err := c.ContactsResetSaved(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.resetSaved - reply: {%v}", r)
	return r, err
}

// ContactsGetSaved
// contacts.getSaved#82f1e39f = Vector<SavedContact>;
func (s *Service) ContactsGetSaved(ctx context.Context, request *tg.TLContactsGetSaved) (*tg.VectorSavedContact, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getSaved - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetSaved(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getSaved - reply: {%v}", r)
	return r, err
}

// ContactsToggleTopPeers
// contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
func (s *Service) ContactsToggleTopPeers(ctx context.Context, request *tg.TLContactsToggleTopPeers) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.toggleTopPeers - metadata: {}, request: {%v}", request)

	r, err := c.ContactsToggleTopPeers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.toggleTopPeers - reply: {%v}", r)
	return r, err
}

// ContactsAddContact
// contacts.addContact#e8f463d0 flags:# add_phone_privacy_exception:flags.0?true id:InputUser first_name:string last_name:string phone:string = Updates;
func (s *Service) ContactsAddContact(ctx context.Context, request *tg.TLContactsAddContact) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.addContact - metadata: {}, request: {%v}", request)

	r, err := c.ContactsAddContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.addContact - reply: {%v}", r)
	return r, err
}

// ContactsAcceptContact
// contacts.acceptContact#f831a20f id:InputUser = Updates;
func (s *Service) ContactsAcceptContact(ctx context.Context, request *tg.TLContactsAcceptContact) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.acceptContact - metadata: {}, request: {%v}", request)

	r, err := c.ContactsAcceptContact(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.acceptContact - reply: {%v}", r)
	return r, err
}

// ContactsGetLocated
// contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
func (s *Service) ContactsGetLocated(ctx context.Context, request *tg.TLContactsGetLocated) (*tg.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.getLocated - metadata: {}, request: {%v}", request)

	r, err := c.ContactsGetLocated(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.getLocated - reply: {%v}", r)
	return r, err
}

// ContactsEditCloseFriends
// contacts.editCloseFriends#ba6705f0 id:Vector<long> = Bool;
func (s *Service) ContactsEditCloseFriends(ctx context.Context, request *tg.TLContactsEditCloseFriends) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.editCloseFriends - metadata: {}, request: {%v}", request)

	r, err := c.ContactsEditCloseFriends(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.editCloseFriends - reply: {%v}", r)
	return r, err
}

// ContactsSetBlocked
// contacts.setBlocked#94c65c76 flags:# my_stories_from:flags.0?true id:Vector<InputPeer> limit:int = Bool;
func (s *Service) ContactsSetBlocked(ctx context.Context, request *tg.TLContactsSetBlocked) (*tg.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("contacts.setBlocked - metadata: {}, request: {%v}", request)

	r, err := c.ContactsSetBlocked(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("contacts.setBlocked - reply: {%v}", r)
	return r, err
}
