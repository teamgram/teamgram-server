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
	"encoding/json"
	"errors"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/tg"

	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"account.getContactSignUpNotification": kitex.NewMethodInfo(
		accountGetContactSignUpNotificationHandler,
		newAccountGetContactSignUpNotificationArgs,
		newAccountGetContactSignUpNotificationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"account.setContactSignUpNotification": kitex.NewMethodInfo(
		accountSetContactSignUpNotificationHandler,
		newAccountSetContactSignUpNotificationArgs,
		newAccountSetContactSignUpNotificationResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getContactIDs": kitex.NewMethodInfo(
		contactsGetContactIDsHandler,
		newContactsGetContactIDsArgs,
		newContactsGetContactIDsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getStatuses": kitex.NewMethodInfo(
		contactsGetStatusesHandler,
		newContactsGetStatusesArgs,
		newContactsGetStatusesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getContacts": kitex.NewMethodInfo(
		contactsGetContactsHandler,
		newContactsGetContactsArgs,
		newContactsGetContactsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.importContacts": kitex.NewMethodInfo(
		contactsImportContactsHandler,
		newContactsImportContactsArgs,
		newContactsImportContactsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.deleteContacts": kitex.NewMethodInfo(
		contactsDeleteContactsHandler,
		newContactsDeleteContactsArgs,
		newContactsDeleteContactsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.deleteByPhones": kitex.NewMethodInfo(
		contactsDeleteByPhonesHandler,
		newContactsDeleteByPhonesArgs,
		newContactsDeleteByPhonesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.block": kitex.NewMethodInfo(
		contactsBlockHandler,
		newContactsBlockArgs,
		newContactsBlockResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.unblock": kitex.NewMethodInfo(
		contactsUnblockHandler,
		newContactsUnblockArgs,
		newContactsUnblockResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getBlocked": kitex.NewMethodInfo(
		contactsGetBlockedHandler,
		newContactsGetBlockedArgs,
		newContactsGetBlockedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.search": kitex.NewMethodInfo(
		contactsSearchHandler,
		newContactsSearchArgs,
		newContactsSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getTopPeers": kitex.NewMethodInfo(
		contactsGetTopPeersHandler,
		newContactsGetTopPeersArgs,
		newContactsGetTopPeersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.resetTopPeerRating": kitex.NewMethodInfo(
		contactsResetTopPeerRatingHandler,
		newContactsResetTopPeerRatingArgs,
		newContactsResetTopPeerRatingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.resetSaved": kitex.NewMethodInfo(
		contactsResetSavedHandler,
		newContactsResetSavedArgs,
		newContactsResetSavedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getSaved": kitex.NewMethodInfo(
		contactsGetSavedHandler,
		newContactsGetSavedArgs,
		newContactsGetSavedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.toggleTopPeers": kitex.NewMethodInfo(
		contactsToggleTopPeersHandler,
		newContactsToggleTopPeersArgs,
		newContactsToggleTopPeersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.addContact": kitex.NewMethodInfo(
		contactsAddContactHandler,
		newContactsAddContactArgs,
		newContactsAddContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.acceptContact": kitex.NewMethodInfo(
		contactsAcceptContactHandler,
		newContactsAcceptContactArgs,
		newContactsAcceptContactResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.getLocated": kitex.NewMethodInfo(
		contactsGetLocatedHandler,
		newContactsGetLocatedArgs,
		newContactsGetLocatedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.editCloseFriends": kitex.NewMethodInfo(
		contactsEditCloseFriendsHandler,
		newContactsEditCloseFriendsArgs,
		newContactsEditCloseFriendsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"contacts.setBlocked": kitex.NewMethodInfo(
		contactsSetBlockedHandler,
		newContactsSetBlockedArgs,
		newContactsSetBlockedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	contactsServiceServiceInfo                = NewServiceInfo()
	contactsServiceServiceInfoForClient       = NewServiceInfoForClient()
	contactsServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return contactsServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return contactsServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return contactsServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfoForClient creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RPCContacts"
	handlerType := (*tg.RPCContacts)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "contacts",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		KiteXGenVersion: "0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func accountGetContactSignUpNotificationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountGetContactSignUpNotificationArgs)
	realResult := result.(*AccountGetContactSignUpNotificationResult)
	success, err := handler.(tg.RPCContacts).AccountGetContactSignUpNotification(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountGetContactSignUpNotificationArgs() interface{} {
	return &AccountGetContactSignUpNotificationArgs{}
}

func newAccountGetContactSignUpNotificationResult() interface{} {
	return &AccountGetContactSignUpNotificationResult{}
}

type AccountGetContactSignUpNotificationArgs struct {
	Req *tg.TLAccountGetContactSignUpNotification
}

func (p *AccountGetContactSignUpNotificationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountGetContactSignUpNotificationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountGetContactSignUpNotificationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountGetContactSignUpNotification)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountGetContactSignUpNotificationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountGetContactSignUpNotificationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountGetContactSignUpNotificationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountGetContactSignUpNotification)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountGetContactSignUpNotificationArgs_Req_DEFAULT *tg.TLAccountGetContactSignUpNotification

func (p *AccountGetContactSignUpNotificationArgs) GetReq() *tg.TLAccountGetContactSignUpNotification {
	if !p.IsSetReq() {
		return AccountGetContactSignUpNotificationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountGetContactSignUpNotificationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountGetContactSignUpNotificationResult struct {
	Success *tg.Bool
}

var AccountGetContactSignUpNotificationResult_Success_DEFAULT *tg.Bool

func (p *AccountGetContactSignUpNotificationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountGetContactSignUpNotificationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountGetContactSignUpNotificationResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetContactSignUpNotificationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountGetContactSignUpNotificationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountGetContactSignUpNotificationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountGetContactSignUpNotificationResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountGetContactSignUpNotificationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountGetContactSignUpNotificationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountGetContactSignUpNotificationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountGetContactSignUpNotificationResult) GetResult() interface{} {
	return p.Success
}

func accountSetContactSignUpNotificationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*AccountSetContactSignUpNotificationArgs)
	realResult := result.(*AccountSetContactSignUpNotificationResult)
	success, err := handler.(tg.RPCContacts).AccountSetContactSignUpNotification(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newAccountSetContactSignUpNotificationArgs() interface{} {
	return &AccountSetContactSignUpNotificationArgs{}
}

func newAccountSetContactSignUpNotificationResult() interface{} {
	return &AccountSetContactSignUpNotificationResult{}
}

type AccountSetContactSignUpNotificationArgs struct {
	Req *tg.TLAccountSetContactSignUpNotification
}

func (p *AccountSetContactSignUpNotificationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AccountSetContactSignUpNotificationArgs")
	}
	return json.Marshal(p.Req)
}

func (p *AccountSetContactSignUpNotificationArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLAccountSetContactSignUpNotification)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *AccountSetContactSignUpNotificationArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in AccountSetContactSignUpNotificationArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *AccountSetContactSignUpNotificationArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLAccountSetContactSignUpNotification)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var AccountSetContactSignUpNotificationArgs_Req_DEFAULT *tg.TLAccountSetContactSignUpNotification

func (p *AccountSetContactSignUpNotificationArgs) GetReq() *tg.TLAccountSetContactSignUpNotification {
	if !p.IsSetReq() {
		return AccountSetContactSignUpNotificationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AccountSetContactSignUpNotificationArgs) IsSetReq() bool {
	return p.Req != nil
}

type AccountSetContactSignUpNotificationResult struct {
	Success *tg.Bool
}

var AccountSetContactSignUpNotificationResult_Success_DEFAULT *tg.Bool

func (p *AccountSetContactSignUpNotificationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AccountSetContactSignUpNotificationResult")
	}
	return json.Marshal(p.Success)
}

func (p *AccountSetContactSignUpNotificationResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetContactSignUpNotificationResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in AccountSetContactSignUpNotificationResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *AccountSetContactSignUpNotificationResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AccountSetContactSignUpNotificationResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return AccountSetContactSignUpNotificationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AccountSetContactSignUpNotificationResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *AccountSetContactSignUpNotificationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountSetContactSignUpNotificationResult) GetResult() interface{} {
	return p.Success
}

func contactsGetContactIDsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetContactIDsArgs)
	realResult := result.(*ContactsGetContactIDsResult)
	success, err := handler.(tg.RPCContacts).ContactsGetContactIDs(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetContactIDsArgs() interface{} {
	return &ContactsGetContactIDsArgs{}
}

func newContactsGetContactIDsResult() interface{} {
	return &ContactsGetContactIDsResult{}
}

type ContactsGetContactIDsArgs struct {
	Req *tg.TLContactsGetContactIDs
}

func (p *ContactsGetContactIDsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetContactIDsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetContactIDsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetContactIDs)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetContactIDsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetContactIDsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetContactIDsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetContactIDs)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetContactIDsArgs_Req_DEFAULT *tg.TLContactsGetContactIDs

func (p *ContactsGetContactIDsArgs) GetReq() *tg.TLContactsGetContactIDs {
	if !p.IsSetReq() {
		return ContactsGetContactIDsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetContactIDsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetContactIDsResult struct {
	Success *tg.VectorInt
}

var ContactsGetContactIDsResult_Success_DEFAULT *tg.VectorInt

func (p *ContactsGetContactIDsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetContactIDsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetContactIDsResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorInt)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetContactIDsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetContactIDsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetContactIDsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorInt)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetContactIDsResult) GetSuccess() *tg.VectorInt {
	if !p.IsSetSuccess() {
		return ContactsGetContactIDsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetContactIDsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorInt)
}

func (p *ContactsGetContactIDsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetContactIDsResult) GetResult() interface{} {
	return p.Success
}

func contactsGetStatusesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetStatusesArgs)
	realResult := result.(*ContactsGetStatusesResult)
	success, err := handler.(tg.RPCContacts).ContactsGetStatuses(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetStatusesArgs() interface{} {
	return &ContactsGetStatusesArgs{}
}

func newContactsGetStatusesResult() interface{} {
	return &ContactsGetStatusesResult{}
}

type ContactsGetStatusesArgs struct {
	Req *tg.TLContactsGetStatuses
}

func (p *ContactsGetStatusesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetStatusesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetStatusesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetStatuses)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetStatusesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetStatusesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetStatusesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetStatuses)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetStatusesArgs_Req_DEFAULT *tg.TLContactsGetStatuses

func (p *ContactsGetStatusesArgs) GetReq() *tg.TLContactsGetStatuses {
	if !p.IsSetReq() {
		return ContactsGetStatusesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetStatusesArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetStatusesResult struct {
	Success *tg.VectorContactStatus
}

var ContactsGetStatusesResult_Success_DEFAULT *tg.VectorContactStatus

func (p *ContactsGetStatusesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetStatusesResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetStatusesResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorContactStatus)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetStatusesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetStatusesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetStatusesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorContactStatus)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetStatusesResult) GetSuccess() *tg.VectorContactStatus {
	if !p.IsSetSuccess() {
		return ContactsGetStatusesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetStatusesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorContactStatus)
}

func (p *ContactsGetStatusesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetStatusesResult) GetResult() interface{} {
	return p.Success
}

func contactsGetContactsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetContactsArgs)
	realResult := result.(*ContactsGetContactsResult)
	success, err := handler.(tg.RPCContacts).ContactsGetContacts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetContactsArgs() interface{} {
	return &ContactsGetContactsArgs{}
}

func newContactsGetContactsResult() interface{} {
	return &ContactsGetContactsResult{}
}

type ContactsGetContactsArgs struct {
	Req *tg.TLContactsGetContacts
}

func (p *ContactsGetContactsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetContactsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetContactsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetContactsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetContactsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetContactsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetContacts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetContactsArgs_Req_DEFAULT *tg.TLContactsGetContacts

func (p *ContactsGetContactsArgs) GetReq() *tg.TLContactsGetContacts {
	if !p.IsSetReq() {
		return ContactsGetContactsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetContactsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetContactsResult struct {
	Success *tg.ContactsContacts
}

var ContactsGetContactsResult_Success_DEFAULT *tg.ContactsContacts

func (p *ContactsGetContactsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetContactsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetContactsResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetContactsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetContactsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetContactsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsContacts)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetContactsResult) GetSuccess() *tg.ContactsContacts {
	if !p.IsSetSuccess() {
		return ContactsGetContactsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetContactsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsContacts)
}

func (p *ContactsGetContactsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetContactsResult) GetResult() interface{} {
	return p.Success
}

func contactsImportContactsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsImportContactsArgs)
	realResult := result.(*ContactsImportContactsResult)
	success, err := handler.(tg.RPCContacts).ContactsImportContacts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsImportContactsArgs() interface{} {
	return &ContactsImportContactsArgs{}
}

func newContactsImportContactsResult() interface{} {
	return &ContactsImportContactsResult{}
}

type ContactsImportContactsArgs struct {
	Req *tg.TLContactsImportContacts
}

func (p *ContactsImportContactsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsImportContactsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsImportContactsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsImportContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsImportContactsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsImportContactsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsImportContactsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsImportContacts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsImportContactsArgs_Req_DEFAULT *tg.TLContactsImportContacts

func (p *ContactsImportContactsArgs) GetReq() *tg.TLContactsImportContacts {
	if !p.IsSetReq() {
		return ContactsImportContactsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsImportContactsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsImportContactsResult struct {
	Success *tg.ContactsImportedContacts
}

var ContactsImportContactsResult_Success_DEFAULT *tg.ContactsImportedContacts

func (p *ContactsImportContactsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsImportContactsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsImportContactsResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsImportedContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsImportContactsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsImportContactsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsImportContactsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsImportedContacts)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsImportContactsResult) GetSuccess() *tg.ContactsImportedContacts {
	if !p.IsSetSuccess() {
		return ContactsImportContactsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsImportContactsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsImportedContacts)
}

func (p *ContactsImportContactsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsImportContactsResult) GetResult() interface{} {
	return p.Success
}

func contactsDeleteContactsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsDeleteContactsArgs)
	realResult := result.(*ContactsDeleteContactsResult)
	success, err := handler.(tg.RPCContacts).ContactsDeleteContacts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsDeleteContactsArgs() interface{} {
	return &ContactsDeleteContactsArgs{}
}

func newContactsDeleteContactsResult() interface{} {
	return &ContactsDeleteContactsResult{}
}

type ContactsDeleteContactsArgs struct {
	Req *tg.TLContactsDeleteContacts
}

func (p *ContactsDeleteContactsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsDeleteContactsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsDeleteContactsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsDeleteContacts)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsDeleteContactsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsDeleteContactsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsDeleteContactsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsDeleteContacts)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsDeleteContactsArgs_Req_DEFAULT *tg.TLContactsDeleteContacts

func (p *ContactsDeleteContactsArgs) GetReq() *tg.TLContactsDeleteContacts {
	if !p.IsSetReq() {
		return ContactsDeleteContactsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsDeleteContactsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsDeleteContactsResult struct {
	Success *tg.Updates
}

var ContactsDeleteContactsResult_Success_DEFAULT *tg.Updates

func (p *ContactsDeleteContactsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsDeleteContactsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsDeleteContactsResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsDeleteContactsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsDeleteContactsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsDeleteContactsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsDeleteContactsResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ContactsDeleteContactsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsDeleteContactsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ContactsDeleteContactsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsDeleteContactsResult) GetResult() interface{} {
	return p.Success
}

func contactsDeleteByPhonesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsDeleteByPhonesArgs)
	realResult := result.(*ContactsDeleteByPhonesResult)
	success, err := handler.(tg.RPCContacts).ContactsDeleteByPhones(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsDeleteByPhonesArgs() interface{} {
	return &ContactsDeleteByPhonesArgs{}
}

func newContactsDeleteByPhonesResult() interface{} {
	return &ContactsDeleteByPhonesResult{}
}

type ContactsDeleteByPhonesArgs struct {
	Req *tg.TLContactsDeleteByPhones
}

func (p *ContactsDeleteByPhonesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsDeleteByPhonesArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsDeleteByPhonesArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsDeleteByPhones)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsDeleteByPhonesArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsDeleteByPhonesArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsDeleteByPhonesArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsDeleteByPhones)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsDeleteByPhonesArgs_Req_DEFAULT *tg.TLContactsDeleteByPhones

func (p *ContactsDeleteByPhonesArgs) GetReq() *tg.TLContactsDeleteByPhones {
	if !p.IsSetReq() {
		return ContactsDeleteByPhonesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsDeleteByPhonesArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsDeleteByPhonesResult struct {
	Success *tg.Bool
}

var ContactsDeleteByPhonesResult_Success_DEFAULT *tg.Bool

func (p *ContactsDeleteByPhonesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsDeleteByPhonesResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsDeleteByPhonesResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsDeleteByPhonesResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsDeleteByPhonesResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsDeleteByPhonesResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsDeleteByPhonesResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsDeleteByPhonesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsDeleteByPhonesResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsDeleteByPhonesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsDeleteByPhonesResult) GetResult() interface{} {
	return p.Success
}

func contactsBlockHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsBlockArgs)
	realResult := result.(*ContactsBlockResult)
	success, err := handler.(tg.RPCContacts).ContactsBlock(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsBlockArgs() interface{} {
	return &ContactsBlockArgs{}
}

func newContactsBlockResult() interface{} {
	return &ContactsBlockResult{}
}

type ContactsBlockArgs struct {
	Req *tg.TLContactsBlock
}

func (p *ContactsBlockArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsBlockArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsBlockArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsBlock)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsBlockArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsBlockArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsBlockArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsBlock)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsBlockArgs_Req_DEFAULT *tg.TLContactsBlock

func (p *ContactsBlockArgs) GetReq() *tg.TLContactsBlock {
	if !p.IsSetReq() {
		return ContactsBlockArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsBlockArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsBlockResult struct {
	Success *tg.Bool
}

var ContactsBlockResult_Success_DEFAULT *tg.Bool

func (p *ContactsBlockResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsBlockResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsBlockResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsBlockResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsBlockResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsBlockResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsBlockResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsBlockResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsBlockResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsBlockResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsBlockResult) GetResult() interface{} {
	return p.Success
}

func contactsUnblockHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsUnblockArgs)
	realResult := result.(*ContactsUnblockResult)
	success, err := handler.(tg.RPCContacts).ContactsUnblock(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsUnblockArgs() interface{} {
	return &ContactsUnblockArgs{}
}

func newContactsUnblockResult() interface{} {
	return &ContactsUnblockResult{}
}

type ContactsUnblockArgs struct {
	Req *tg.TLContactsUnblock
}

func (p *ContactsUnblockArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsUnblockArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsUnblockArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsUnblock)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsUnblockArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsUnblockArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsUnblockArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsUnblock)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsUnblockArgs_Req_DEFAULT *tg.TLContactsUnblock

func (p *ContactsUnblockArgs) GetReq() *tg.TLContactsUnblock {
	if !p.IsSetReq() {
		return ContactsUnblockArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsUnblockArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsUnblockResult struct {
	Success *tg.Bool
}

var ContactsUnblockResult_Success_DEFAULT *tg.Bool

func (p *ContactsUnblockResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsUnblockResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsUnblockResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsUnblockResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsUnblockResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsUnblockResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsUnblockResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsUnblockResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsUnblockResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsUnblockResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsUnblockResult) GetResult() interface{} {
	return p.Success
}

func contactsGetBlockedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetBlockedArgs)
	realResult := result.(*ContactsGetBlockedResult)
	success, err := handler.(tg.RPCContacts).ContactsGetBlocked(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetBlockedArgs() interface{} {
	return &ContactsGetBlockedArgs{}
}

func newContactsGetBlockedResult() interface{} {
	return &ContactsGetBlockedResult{}
}

type ContactsGetBlockedArgs struct {
	Req *tg.TLContactsGetBlocked
}

func (p *ContactsGetBlockedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetBlockedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetBlockedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetBlocked)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetBlockedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetBlockedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetBlockedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetBlocked)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetBlockedArgs_Req_DEFAULT *tg.TLContactsGetBlocked

func (p *ContactsGetBlockedArgs) GetReq() *tg.TLContactsGetBlocked {
	if !p.IsSetReq() {
		return ContactsGetBlockedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetBlockedArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetBlockedResult struct {
	Success *tg.ContactsBlocked
}

var ContactsGetBlockedResult_Success_DEFAULT *tg.ContactsBlocked

func (p *ContactsGetBlockedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetBlockedResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetBlockedResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsBlocked)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetBlockedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetBlockedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetBlockedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsBlocked)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetBlockedResult) GetSuccess() *tg.ContactsBlocked {
	if !p.IsSetSuccess() {
		return ContactsGetBlockedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetBlockedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsBlocked)
}

func (p *ContactsGetBlockedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetBlockedResult) GetResult() interface{} {
	return p.Success
}

func contactsSearchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsSearchArgs)
	realResult := result.(*ContactsSearchResult)
	success, err := handler.(tg.RPCContacts).ContactsSearch(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsSearchArgs() interface{} {
	return &ContactsSearchArgs{}
}

func newContactsSearchResult() interface{} {
	return &ContactsSearchResult{}
}

type ContactsSearchArgs struct {
	Req *tg.TLContactsSearch
}

func (p *ContactsSearchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsSearchArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsSearchArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsSearch)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsSearchArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsSearchArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsSearchArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsSearch)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsSearchArgs_Req_DEFAULT *tg.TLContactsSearch

func (p *ContactsSearchArgs) GetReq() *tg.TLContactsSearch {
	if !p.IsSetReq() {
		return ContactsSearchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsSearchArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsSearchResult struct {
	Success *tg.ContactsFound
}

var ContactsSearchResult_Success_DEFAULT *tg.ContactsFound

func (p *ContactsSearchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsSearchResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsSearchResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsFound)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsSearchResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsSearchResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsSearchResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsFound)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsSearchResult) GetSuccess() *tg.ContactsFound {
	if !p.IsSetSuccess() {
		return ContactsSearchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsSearchResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsFound)
}

func (p *ContactsSearchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsSearchResult) GetResult() interface{} {
	return p.Success
}

func contactsGetTopPeersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetTopPeersArgs)
	realResult := result.(*ContactsGetTopPeersResult)
	success, err := handler.(tg.RPCContacts).ContactsGetTopPeers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetTopPeersArgs() interface{} {
	return &ContactsGetTopPeersArgs{}
}

func newContactsGetTopPeersResult() interface{} {
	return &ContactsGetTopPeersResult{}
}

type ContactsGetTopPeersArgs struct {
	Req *tg.TLContactsGetTopPeers
}

func (p *ContactsGetTopPeersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetTopPeersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetTopPeersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetTopPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetTopPeersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetTopPeersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetTopPeersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetTopPeers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetTopPeersArgs_Req_DEFAULT *tg.TLContactsGetTopPeers

func (p *ContactsGetTopPeersArgs) GetReq() *tg.TLContactsGetTopPeers {
	if !p.IsSetReq() {
		return ContactsGetTopPeersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetTopPeersArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetTopPeersResult struct {
	Success *tg.ContactsTopPeers
}

var ContactsGetTopPeersResult_Success_DEFAULT *tg.ContactsTopPeers

func (p *ContactsGetTopPeersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetTopPeersResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetTopPeersResult) Unmarshal(in []byte) error {
	msg := new(tg.ContactsTopPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetTopPeersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetTopPeersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetTopPeersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.ContactsTopPeers)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetTopPeersResult) GetSuccess() *tg.ContactsTopPeers {
	if !p.IsSetSuccess() {
		return ContactsGetTopPeersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetTopPeersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.ContactsTopPeers)
}

func (p *ContactsGetTopPeersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetTopPeersResult) GetResult() interface{} {
	return p.Success
}

func contactsResetTopPeerRatingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsResetTopPeerRatingArgs)
	realResult := result.(*ContactsResetTopPeerRatingResult)
	success, err := handler.(tg.RPCContacts).ContactsResetTopPeerRating(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsResetTopPeerRatingArgs() interface{} {
	return &ContactsResetTopPeerRatingArgs{}
}

func newContactsResetTopPeerRatingResult() interface{} {
	return &ContactsResetTopPeerRatingResult{}
}

type ContactsResetTopPeerRatingArgs struct {
	Req *tg.TLContactsResetTopPeerRating
}

func (p *ContactsResetTopPeerRatingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsResetTopPeerRatingArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsResetTopPeerRatingArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsResetTopPeerRating)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsResetTopPeerRatingArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsResetTopPeerRatingArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsResetTopPeerRatingArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsResetTopPeerRating)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsResetTopPeerRatingArgs_Req_DEFAULT *tg.TLContactsResetTopPeerRating

func (p *ContactsResetTopPeerRatingArgs) GetReq() *tg.TLContactsResetTopPeerRating {
	if !p.IsSetReq() {
		return ContactsResetTopPeerRatingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsResetTopPeerRatingArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsResetTopPeerRatingResult struct {
	Success *tg.Bool
}

var ContactsResetTopPeerRatingResult_Success_DEFAULT *tg.Bool

func (p *ContactsResetTopPeerRatingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsResetTopPeerRatingResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsResetTopPeerRatingResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResetTopPeerRatingResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsResetTopPeerRatingResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsResetTopPeerRatingResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResetTopPeerRatingResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsResetTopPeerRatingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsResetTopPeerRatingResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsResetTopPeerRatingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsResetTopPeerRatingResult) GetResult() interface{} {
	return p.Success
}

func contactsResetSavedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsResetSavedArgs)
	realResult := result.(*ContactsResetSavedResult)
	success, err := handler.(tg.RPCContacts).ContactsResetSaved(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsResetSavedArgs() interface{} {
	return &ContactsResetSavedArgs{}
}

func newContactsResetSavedResult() interface{} {
	return &ContactsResetSavedResult{}
}

type ContactsResetSavedArgs struct {
	Req *tg.TLContactsResetSaved
}

func (p *ContactsResetSavedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsResetSavedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsResetSavedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsResetSaved)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsResetSavedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsResetSavedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsResetSavedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsResetSaved)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsResetSavedArgs_Req_DEFAULT *tg.TLContactsResetSaved

func (p *ContactsResetSavedArgs) GetReq() *tg.TLContactsResetSaved {
	if !p.IsSetReq() {
		return ContactsResetSavedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsResetSavedArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsResetSavedResult struct {
	Success *tg.Bool
}

var ContactsResetSavedResult_Success_DEFAULT *tg.Bool

func (p *ContactsResetSavedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsResetSavedResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsResetSavedResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResetSavedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsResetSavedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsResetSavedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsResetSavedResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsResetSavedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsResetSavedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsResetSavedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsResetSavedResult) GetResult() interface{} {
	return p.Success
}

func contactsGetSavedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetSavedArgs)
	realResult := result.(*ContactsGetSavedResult)
	success, err := handler.(tg.RPCContacts).ContactsGetSaved(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetSavedArgs() interface{} {
	return &ContactsGetSavedArgs{}
}

func newContactsGetSavedResult() interface{} {
	return &ContactsGetSavedResult{}
}

type ContactsGetSavedArgs struct {
	Req *tg.TLContactsGetSaved
}

func (p *ContactsGetSavedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetSavedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetSavedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetSaved)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetSavedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetSavedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetSavedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetSaved)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetSavedArgs_Req_DEFAULT *tg.TLContactsGetSaved

func (p *ContactsGetSavedArgs) GetReq() *tg.TLContactsGetSaved {
	if !p.IsSetReq() {
		return ContactsGetSavedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetSavedArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetSavedResult struct {
	Success *tg.VectorSavedContact
}

var ContactsGetSavedResult_Success_DEFAULT *tg.VectorSavedContact

func (p *ContactsGetSavedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetSavedResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetSavedResult) Unmarshal(in []byte) error {
	msg := new(tg.VectorSavedContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetSavedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetSavedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetSavedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.VectorSavedContact)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetSavedResult) GetSuccess() *tg.VectorSavedContact {
	if !p.IsSetSuccess() {
		return ContactsGetSavedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetSavedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.VectorSavedContact)
}

func (p *ContactsGetSavedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetSavedResult) GetResult() interface{} {
	return p.Success
}

func contactsToggleTopPeersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsToggleTopPeersArgs)
	realResult := result.(*ContactsToggleTopPeersResult)
	success, err := handler.(tg.RPCContacts).ContactsToggleTopPeers(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsToggleTopPeersArgs() interface{} {
	return &ContactsToggleTopPeersArgs{}
}

func newContactsToggleTopPeersResult() interface{} {
	return &ContactsToggleTopPeersResult{}
}

type ContactsToggleTopPeersArgs struct {
	Req *tg.TLContactsToggleTopPeers
}

func (p *ContactsToggleTopPeersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsToggleTopPeersArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsToggleTopPeersArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsToggleTopPeers)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsToggleTopPeersArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsToggleTopPeersArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsToggleTopPeersArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsToggleTopPeers)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsToggleTopPeersArgs_Req_DEFAULT *tg.TLContactsToggleTopPeers

func (p *ContactsToggleTopPeersArgs) GetReq() *tg.TLContactsToggleTopPeers {
	if !p.IsSetReq() {
		return ContactsToggleTopPeersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsToggleTopPeersArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsToggleTopPeersResult struct {
	Success *tg.Bool
}

var ContactsToggleTopPeersResult_Success_DEFAULT *tg.Bool

func (p *ContactsToggleTopPeersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsToggleTopPeersResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsToggleTopPeersResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsToggleTopPeersResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsToggleTopPeersResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsToggleTopPeersResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsToggleTopPeersResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsToggleTopPeersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsToggleTopPeersResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsToggleTopPeersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsToggleTopPeersResult) GetResult() interface{} {
	return p.Success
}

func contactsAddContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsAddContactArgs)
	realResult := result.(*ContactsAddContactResult)
	success, err := handler.(tg.RPCContacts).ContactsAddContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsAddContactArgs() interface{} {
	return &ContactsAddContactArgs{}
}

func newContactsAddContactResult() interface{} {
	return &ContactsAddContactResult{}
}

type ContactsAddContactArgs struct {
	Req *tg.TLContactsAddContact
}

func (p *ContactsAddContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsAddContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsAddContactArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsAddContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsAddContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsAddContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsAddContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsAddContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsAddContactArgs_Req_DEFAULT *tg.TLContactsAddContact

func (p *ContactsAddContactArgs) GetReq() *tg.TLContactsAddContact {
	if !p.IsSetReq() {
		return ContactsAddContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsAddContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsAddContactResult struct {
	Success *tg.Updates
}

var ContactsAddContactResult_Success_DEFAULT *tg.Updates

func (p *ContactsAddContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsAddContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsAddContactResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsAddContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsAddContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsAddContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsAddContactResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ContactsAddContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsAddContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ContactsAddContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsAddContactResult) GetResult() interface{} {
	return p.Success
}

func contactsAcceptContactHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsAcceptContactArgs)
	realResult := result.(*ContactsAcceptContactResult)
	success, err := handler.(tg.RPCContacts).ContactsAcceptContact(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsAcceptContactArgs() interface{} {
	return &ContactsAcceptContactArgs{}
}

func newContactsAcceptContactResult() interface{} {
	return &ContactsAcceptContactResult{}
}

type ContactsAcceptContactArgs struct {
	Req *tg.TLContactsAcceptContact
}

func (p *ContactsAcceptContactArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsAcceptContactArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsAcceptContactArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsAcceptContact)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsAcceptContactArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsAcceptContactArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsAcceptContactArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsAcceptContact)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsAcceptContactArgs_Req_DEFAULT *tg.TLContactsAcceptContact

func (p *ContactsAcceptContactArgs) GetReq() *tg.TLContactsAcceptContact {
	if !p.IsSetReq() {
		return ContactsAcceptContactArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsAcceptContactArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsAcceptContactResult struct {
	Success *tg.Updates
}

var ContactsAcceptContactResult_Success_DEFAULT *tg.Updates

func (p *ContactsAcceptContactResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsAcceptContactResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsAcceptContactResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsAcceptContactResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsAcceptContactResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsAcceptContactResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsAcceptContactResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ContactsAcceptContactResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsAcceptContactResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ContactsAcceptContactResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsAcceptContactResult) GetResult() interface{} {
	return p.Success
}

func contactsGetLocatedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsGetLocatedArgs)
	realResult := result.(*ContactsGetLocatedResult)
	success, err := handler.(tg.RPCContacts).ContactsGetLocated(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsGetLocatedArgs() interface{} {
	return &ContactsGetLocatedArgs{}
}

func newContactsGetLocatedResult() interface{} {
	return &ContactsGetLocatedResult{}
}

type ContactsGetLocatedArgs struct {
	Req *tg.TLContactsGetLocated
}

func (p *ContactsGetLocatedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsGetLocatedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsGetLocatedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsGetLocated)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsGetLocatedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsGetLocatedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsGetLocatedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsGetLocated)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsGetLocatedArgs_Req_DEFAULT *tg.TLContactsGetLocated

func (p *ContactsGetLocatedArgs) GetReq() *tg.TLContactsGetLocated {
	if !p.IsSetReq() {
		return ContactsGetLocatedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsGetLocatedArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsGetLocatedResult struct {
	Success *tg.Updates
}

var ContactsGetLocatedResult_Success_DEFAULT *tg.Updates

func (p *ContactsGetLocatedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsGetLocatedResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsGetLocatedResult) Unmarshal(in []byte) error {
	msg := new(tg.Updates)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetLocatedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsGetLocatedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsGetLocatedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Updates)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsGetLocatedResult) GetSuccess() *tg.Updates {
	if !p.IsSetSuccess() {
		return ContactsGetLocatedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsGetLocatedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Updates)
}

func (p *ContactsGetLocatedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsGetLocatedResult) GetResult() interface{} {
	return p.Success
}

func contactsEditCloseFriendsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsEditCloseFriendsArgs)
	realResult := result.(*ContactsEditCloseFriendsResult)
	success, err := handler.(tg.RPCContacts).ContactsEditCloseFriends(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsEditCloseFriendsArgs() interface{} {
	return &ContactsEditCloseFriendsArgs{}
}

func newContactsEditCloseFriendsResult() interface{} {
	return &ContactsEditCloseFriendsResult{}
}

type ContactsEditCloseFriendsArgs struct {
	Req *tg.TLContactsEditCloseFriends
}

func (p *ContactsEditCloseFriendsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsEditCloseFriendsArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsEditCloseFriendsArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsEditCloseFriends)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsEditCloseFriendsArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsEditCloseFriendsArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsEditCloseFriendsArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsEditCloseFriends)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsEditCloseFriendsArgs_Req_DEFAULT *tg.TLContactsEditCloseFriends

func (p *ContactsEditCloseFriendsArgs) GetReq() *tg.TLContactsEditCloseFriends {
	if !p.IsSetReq() {
		return ContactsEditCloseFriendsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsEditCloseFriendsArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsEditCloseFriendsResult struct {
	Success *tg.Bool
}

var ContactsEditCloseFriendsResult_Success_DEFAULT *tg.Bool

func (p *ContactsEditCloseFriendsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsEditCloseFriendsResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsEditCloseFriendsResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsEditCloseFriendsResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsEditCloseFriendsResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsEditCloseFriendsResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsEditCloseFriendsResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsEditCloseFriendsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsEditCloseFriendsResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsEditCloseFriendsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsEditCloseFriendsResult) GetResult() interface{} {
	return p.Success
}

func contactsSetBlockedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ContactsSetBlockedArgs)
	realResult := result.(*ContactsSetBlockedResult)
	success, err := handler.(tg.RPCContacts).ContactsSetBlocked(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}

func newContactsSetBlockedArgs() interface{} {
	return &ContactsSetBlockedArgs{}
}

func newContactsSetBlockedResult() interface{} {
	return &ContactsSetBlockedResult{}
}

type ContactsSetBlockedArgs struct {
	Req *tg.TLContactsSetBlocked
}

func (p *ContactsSetBlockedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ContactsSetBlockedArgs")
	}
	return json.Marshal(p.Req)
}

func (p *ContactsSetBlockedArgs) Unmarshal(in []byte) error {
	msg := new(tg.TLContactsSetBlocked)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

func (p *ContactsSetBlockedArgs) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetReq() {
		return fmt.Errorf("No req in ContactsSetBlockedArgs")
	}

	return p.Req.Encode(x, layer)
}

func (p *ContactsSetBlockedArgs) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.TLContactsSetBlocked)
	msg.ClazzID, _ = d.ClazzID()
	msg.Decode(d)
	p.Req = msg
	return nil
}

var ContactsSetBlockedArgs_Req_DEFAULT *tg.TLContactsSetBlocked

func (p *ContactsSetBlockedArgs) GetReq() *tg.TLContactsSetBlocked {
	if !p.IsSetReq() {
		return ContactsSetBlockedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ContactsSetBlockedArgs) IsSetReq() bool {
	return p.Req != nil
}

type ContactsSetBlockedResult struct {
	Success *tg.Bool
}

var ContactsSetBlockedResult_Success_DEFAULT *tg.Bool

func (p *ContactsSetBlockedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ContactsSetBlockedResult")
	}
	return json.Marshal(p.Success)
}

func (p *ContactsSetBlockedResult) Unmarshal(in []byte) error {
	msg := new(tg.Bool)
	if err := json.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsSetBlockedResult) Encode(x *bin.Encoder, layer int32) error {
	if !p.IsSetSuccess() {
		return fmt.Errorf("No req in ContactsSetBlockedResult")
	}

	return p.Success.Encode(x, layer)
}

func (p *ContactsSetBlockedResult) Decode(d *bin.Decoder) (err error) {
	msg := new(tg.Bool)
	if err = msg.Decode(d); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ContactsSetBlockedResult) GetSuccess() *tg.Bool {
	if !p.IsSetSuccess() {
		return ContactsSetBlockedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ContactsSetBlockedResult) SetSuccess(x interface{}) {
	p.Success = x.(*tg.Bool)
}

func (p *ContactsSetBlockedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ContactsSetBlockedResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AccountGetContactSignUpNotification(ctx context.Context, req *tg.TLAccountGetContactSignUpNotification) (r *tg.Bool, err error) {
	var _args AccountGetContactSignUpNotificationArgs
	_args.Req = req
	var _result AccountGetContactSignUpNotificationResult
	if err = p.c.Call(ctx, "account.getContactSignUpNotification", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AccountSetContactSignUpNotification(ctx context.Context, req *tg.TLAccountSetContactSignUpNotification) (r *tg.Bool, err error) {
	var _args AccountSetContactSignUpNotificationArgs
	_args.Req = req
	var _result AccountSetContactSignUpNotificationResult
	if err = p.c.Call(ctx, "account.setContactSignUpNotification", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetContactIDs(ctx context.Context, req *tg.TLContactsGetContactIDs) (r *tg.VectorInt, err error) {
	var _args ContactsGetContactIDsArgs
	_args.Req = req
	var _result ContactsGetContactIDsResult
	if err = p.c.Call(ctx, "contacts.getContactIDs", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetStatuses(ctx context.Context, req *tg.TLContactsGetStatuses) (r *tg.VectorContactStatus, err error) {
	var _args ContactsGetStatusesArgs
	_args.Req = req
	var _result ContactsGetStatusesResult
	if err = p.c.Call(ctx, "contacts.getStatuses", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetContacts(ctx context.Context, req *tg.TLContactsGetContacts) (r *tg.ContactsContacts, err error) {
	var _args ContactsGetContactsArgs
	_args.Req = req
	var _result ContactsGetContactsResult
	if err = p.c.Call(ctx, "contacts.getContacts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsImportContacts(ctx context.Context, req *tg.TLContactsImportContacts) (r *tg.ContactsImportedContacts, err error) {
	var _args ContactsImportContactsArgs
	_args.Req = req
	var _result ContactsImportContactsResult
	if err = p.c.Call(ctx, "contacts.importContacts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsDeleteContacts(ctx context.Context, req *tg.TLContactsDeleteContacts) (r *tg.Updates, err error) {
	var _args ContactsDeleteContactsArgs
	_args.Req = req
	var _result ContactsDeleteContactsResult
	if err = p.c.Call(ctx, "contacts.deleteContacts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsDeleteByPhones(ctx context.Context, req *tg.TLContactsDeleteByPhones) (r *tg.Bool, err error) {
	var _args ContactsDeleteByPhonesArgs
	_args.Req = req
	var _result ContactsDeleteByPhonesResult
	if err = p.c.Call(ctx, "contacts.deleteByPhones", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsBlock(ctx context.Context, req *tg.TLContactsBlock) (r *tg.Bool, err error) {
	var _args ContactsBlockArgs
	_args.Req = req
	var _result ContactsBlockResult
	if err = p.c.Call(ctx, "contacts.block", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsUnblock(ctx context.Context, req *tg.TLContactsUnblock) (r *tg.Bool, err error) {
	var _args ContactsUnblockArgs
	_args.Req = req
	var _result ContactsUnblockResult
	if err = p.c.Call(ctx, "contacts.unblock", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetBlocked(ctx context.Context, req *tg.TLContactsGetBlocked) (r *tg.ContactsBlocked, err error) {
	var _args ContactsGetBlockedArgs
	_args.Req = req
	var _result ContactsGetBlockedResult
	if err = p.c.Call(ctx, "contacts.getBlocked", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsSearch(ctx context.Context, req *tg.TLContactsSearch) (r *tg.ContactsFound, err error) {
	var _args ContactsSearchArgs
	_args.Req = req
	var _result ContactsSearchResult
	if err = p.c.Call(ctx, "contacts.search", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetTopPeers(ctx context.Context, req *tg.TLContactsGetTopPeers) (r *tg.ContactsTopPeers, err error) {
	var _args ContactsGetTopPeersArgs
	_args.Req = req
	var _result ContactsGetTopPeersResult
	if err = p.c.Call(ctx, "contacts.getTopPeers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsResetTopPeerRating(ctx context.Context, req *tg.TLContactsResetTopPeerRating) (r *tg.Bool, err error) {
	var _args ContactsResetTopPeerRatingArgs
	_args.Req = req
	var _result ContactsResetTopPeerRatingResult
	if err = p.c.Call(ctx, "contacts.resetTopPeerRating", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsResetSaved(ctx context.Context, req *tg.TLContactsResetSaved) (r *tg.Bool, err error) {
	var _args ContactsResetSavedArgs
	_args.Req = req
	var _result ContactsResetSavedResult
	if err = p.c.Call(ctx, "contacts.resetSaved", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetSaved(ctx context.Context, req *tg.TLContactsGetSaved) (r *tg.VectorSavedContact, err error) {
	var _args ContactsGetSavedArgs
	_args.Req = req
	var _result ContactsGetSavedResult
	if err = p.c.Call(ctx, "contacts.getSaved", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsToggleTopPeers(ctx context.Context, req *tg.TLContactsToggleTopPeers) (r *tg.Bool, err error) {
	var _args ContactsToggleTopPeersArgs
	_args.Req = req
	var _result ContactsToggleTopPeersResult
	if err = p.c.Call(ctx, "contacts.toggleTopPeers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsAddContact(ctx context.Context, req *tg.TLContactsAddContact) (r *tg.Updates, err error) {
	var _args ContactsAddContactArgs
	_args.Req = req
	var _result ContactsAddContactResult
	if err = p.c.Call(ctx, "contacts.addContact", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsAcceptContact(ctx context.Context, req *tg.TLContactsAcceptContact) (r *tg.Updates, err error) {
	var _args ContactsAcceptContactArgs
	_args.Req = req
	var _result ContactsAcceptContactResult
	if err = p.c.Call(ctx, "contacts.acceptContact", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsGetLocated(ctx context.Context, req *tg.TLContactsGetLocated) (r *tg.Updates, err error) {
	var _args ContactsGetLocatedArgs
	_args.Req = req
	var _result ContactsGetLocatedResult
	if err = p.c.Call(ctx, "contacts.getLocated", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsEditCloseFriends(ctx context.Context, req *tg.TLContactsEditCloseFriends) (r *tg.Bool, err error) {
	var _args ContactsEditCloseFriendsArgs
	_args.Req = req
	var _result ContactsEditCloseFriendsResult
	if err = p.c.Call(ctx, "contacts.editCloseFriends", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContactsSetBlocked(ctx context.Context, req *tg.TLContactsSetBlocked) (r *tg.Bool, err error) {
	var _args ContactsSetBlockedArgs
	_args.Req = req
	var _result ContactsSetBlockedResult
	if err = p.c.Call(ctx, "contacts.setBlocked", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
