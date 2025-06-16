/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package contactsservice

import (
	"context"

	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AccountGetContactSignUpNotification(ctx context.Context, req *tg.TLAccountGetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error)
	AccountSetContactSignUpNotification(ctx context.Context, req *tg.TLAccountSetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsGetContactIDs(ctx context.Context, req *tg.TLContactsGetContactIDs, callOptions ...callopt.Option) (r *tg.VectorInt, err error)
	ContactsGetStatuses(ctx context.Context, req *tg.TLContactsGetStatuses, callOptions ...callopt.Option) (r *tg.VectorContactStatus, err error)
	ContactsGetContacts(ctx context.Context, req *tg.TLContactsGetContacts, callOptions ...callopt.Option) (r *tg.ContactsContacts, err error)
	ContactsImportContacts(ctx context.Context, req *tg.TLContactsImportContacts, callOptions ...callopt.Option) (r *tg.ContactsImportedContacts, err error)
	ContactsDeleteContacts(ctx context.Context, req *tg.TLContactsDeleteContacts, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ContactsDeleteByPhones(ctx context.Context, req *tg.TLContactsDeleteByPhones, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsBlock(ctx context.Context, req *tg.TLContactsBlock, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsUnblock(ctx context.Context, req *tg.TLContactsUnblock, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsGetBlocked(ctx context.Context, req *tg.TLContactsGetBlocked, callOptions ...callopt.Option) (r *tg.ContactsBlocked, err error)
	ContactsSearch(ctx context.Context, req *tg.TLContactsSearch, callOptions ...callopt.Option) (r *tg.ContactsFound, err error)
	ContactsGetTopPeers(ctx context.Context, req *tg.TLContactsGetTopPeers, callOptions ...callopt.Option) (r *tg.ContactsTopPeers, err error)
	ContactsResetTopPeerRating(ctx context.Context, req *tg.TLContactsResetTopPeerRating, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsResetSaved(ctx context.Context, req *tg.TLContactsResetSaved, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsGetSaved(ctx context.Context, req *tg.TLContactsGetSaved, callOptions ...callopt.Option) (r *tg.VectorSavedContact, err error)
	ContactsToggleTopPeers(ctx context.Context, req *tg.TLContactsToggleTopPeers, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsAddContact(ctx context.Context, req *tg.TLContactsAddContact, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ContactsAcceptContact(ctx context.Context, req *tg.TLContactsAcceptContact, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ContactsGetLocated(ctx context.Context, req *tg.TLContactsGetLocated, callOptions ...callopt.Option) (r *tg.Updates, err error)
	ContactsEditCloseFriends(ctx context.Context, req *tg.TLContactsEditCloseFriends, callOptions ...callopt.Option) (r *tg.Bool, err error)
	ContactsSetBlocked(ctx context.Context, req *tg.TLContactsSetBlocked, callOptions ...callopt.Option) (r *tg.Bool, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kContactsClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kContactsClient struct {
	*kClient
}

func NewRPCContactsClient(cli client.Client) Client {
	return &kContactsClient{
		kClient: newServiceClient(cli),
	}
}

func (p *kContactsClient) AccountGetContactSignUpNotification(ctx context.Context, req *tg.TLAccountGetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountGetContactSignUpNotification(ctx, req)
}

func (p *kContactsClient) AccountSetContactSignUpNotification(ctx context.Context, req *tg.TLAccountSetContactSignUpNotification, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AccountSetContactSignUpNotification(ctx, req)
}

func (p *kContactsClient) ContactsGetContactIDs(ctx context.Context, req *tg.TLContactsGetContactIDs, callOptions ...callopt.Option) (r *tg.VectorInt, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetContactIDs(ctx, req)
}

func (p *kContactsClient) ContactsGetStatuses(ctx context.Context, req *tg.TLContactsGetStatuses, callOptions ...callopt.Option) (r *tg.VectorContactStatus, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetStatuses(ctx, req)
}

func (p *kContactsClient) ContactsGetContacts(ctx context.Context, req *tg.TLContactsGetContacts, callOptions ...callopt.Option) (r *tg.ContactsContacts, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetContacts(ctx, req)
}

func (p *kContactsClient) ContactsImportContacts(ctx context.Context, req *tg.TLContactsImportContacts, callOptions ...callopt.Option) (r *tg.ContactsImportedContacts, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsImportContacts(ctx, req)
}

func (p *kContactsClient) ContactsDeleteContacts(ctx context.Context, req *tg.TLContactsDeleteContacts, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsDeleteContacts(ctx, req)
}

func (p *kContactsClient) ContactsDeleteByPhones(ctx context.Context, req *tg.TLContactsDeleteByPhones, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsDeleteByPhones(ctx, req)
}

func (p *kContactsClient) ContactsBlock(ctx context.Context, req *tg.TLContactsBlock, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsBlock(ctx, req)
}

func (p *kContactsClient) ContactsUnblock(ctx context.Context, req *tg.TLContactsUnblock, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsUnblock(ctx, req)
}

func (p *kContactsClient) ContactsGetBlocked(ctx context.Context, req *tg.TLContactsGetBlocked, callOptions ...callopt.Option) (r *tg.ContactsBlocked, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetBlocked(ctx, req)
}

func (p *kContactsClient) ContactsSearch(ctx context.Context, req *tg.TLContactsSearch, callOptions ...callopt.Option) (r *tg.ContactsFound, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsSearch(ctx, req)
}

func (p *kContactsClient) ContactsGetTopPeers(ctx context.Context, req *tg.TLContactsGetTopPeers, callOptions ...callopt.Option) (r *tg.ContactsTopPeers, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetTopPeers(ctx, req)
}

func (p *kContactsClient) ContactsResetTopPeerRating(ctx context.Context, req *tg.TLContactsResetTopPeerRating, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsResetTopPeerRating(ctx, req)
}

func (p *kContactsClient) ContactsResetSaved(ctx context.Context, req *tg.TLContactsResetSaved, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsResetSaved(ctx, req)
}

func (p *kContactsClient) ContactsGetSaved(ctx context.Context, req *tg.TLContactsGetSaved, callOptions ...callopt.Option) (r *tg.VectorSavedContact, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetSaved(ctx, req)
}

func (p *kContactsClient) ContactsToggleTopPeers(ctx context.Context, req *tg.TLContactsToggleTopPeers, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsToggleTopPeers(ctx, req)
}

func (p *kContactsClient) ContactsAddContact(ctx context.Context, req *tg.TLContactsAddContact, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsAddContact(ctx, req)
}

func (p *kContactsClient) ContactsAcceptContact(ctx context.Context, req *tg.TLContactsAcceptContact, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsAcceptContact(ctx, req)
}

func (p *kContactsClient) ContactsGetLocated(ctx context.Context, req *tg.TLContactsGetLocated, callOptions ...callopt.Option) (r *tg.Updates, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsGetLocated(ctx, req)
}

func (p *kContactsClient) ContactsEditCloseFriends(ctx context.Context, req *tg.TLContactsEditCloseFriends, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsEditCloseFriends(ctx, req)
}

func (p *kContactsClient) ContactsSetBlocked(ctx context.Context, req *tg.TLContactsSetBlocked, callOptions ...callopt.Option) (r *tg.Bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContactsSetBlocked(ctx, req)
}
