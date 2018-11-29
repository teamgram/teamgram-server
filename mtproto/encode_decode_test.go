// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package mtproto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestTLAccountUpdateStatus(t *testing.T) {
	o := &TLAccountUpdateStatus{}
	//t2 := &TLBoolTrue{}
	// o.Offline = MakeBool(t2)
	fmt.Println("status1: ", o)

	b := o.Encode()
	fmt.Println(hex.EncodeToString(b))
	d := NewDecodeBuf(b)
	// o2 := d.Object()
	d.Int()

	o2 := &TLAccountUpdateStatus{}
	o2.Decode(d)

	//o2.Offline = &Bool{}
	//m := o2.Offline
	//classId := d.Int()
	//switch classId {
	//case int32(TLConstructor_CRC32_boolFalse):
	//	fmt.Println("Bool_BoolFalse")
	//	m2 := &Bool_BoolFalse{}
	//	m2.BoolFalse = &TLBoolFalse{}
	//	m.Payload = m2
	//	//  = &TLBoolFalse{}
	//
	//case int32(TLConstructor_CRC32_boolTrue):
	//	fmt.Println("Bool_BoolTrue")
	//	m2 := &Bool_BoolTrue{}
	//	m2.BoolTrue = &TLBoolTrue{}
	//	// m2.BoolTrue.Decode(d)
	//	m.Payload = m2 // &Bool_BoolTrue{}
	//
	//default:
	//	fmt.Println("classId error: ", classId)
	//}
	// return dbuf.err

	// 妈的，有问题
	fmt.Println("status2: ", o2)
}

func TestUsersGetFullUser(t *testing.T) {
	//fullUser := &TLUserFull{}
	//fullUser.PhoneCallsAvailable = true
	//fullUser.PhoneCallsPrivate = true
	//fullUser.About = "@Benqi"
	//fullUser.CommonChatsCount = 0
	//
	//// switch request.Id.Payload.(type) {
	//// case *InputUser_InputUserSelf:
	//	// User
	//	// userDO, _ := s.UsersDAO.SelectById(2)
	//	user := &TLUser{}
	//	user.Self = true
	//	user.Contact = false
	//	user.Id = 2
	//	user.FirstName = "benqi"
	//	user.LastName = "benqi"
	//	user.Username = "benqi"
	//	user.AccessHash = 1234567890
	//	user.Phone = "+86 111 1111 1111"
	//	fullUser.User = MakeUser(user)
	//
	//	// Link
	//	link := &TLContactsLink{}
	//	link.MyLink = MakeContactLink(&TLContactLinkContact{})
	//	link.ForeignLink = MakeContactLink(&TLContactLinkContact{})
	//	link.User = fullUser.User
	//	fullUser.Link = MakeContacts_Link(link)
	//
	//// case *InputUser_InputUser:
	//// case *InputUser_InputUserEmpty:
	////	// TODO(@benqi): BAD_REQUEST: 400
	//// }
	//
	//// NotifySettings
	//peerNotifySettings := &TLPeerNotifySettings{}
	//peerNotifySettings.ShowPreviews = true
	//peerNotifySettings.MuteUntil = 0
	//peerNotifySettings.Sound = "default"
	//fullUser.NotifySettings = MakePeerNotifySettings(peerNotifySettings)
	//
	//reply := MakeUserFull(fullUser)
	//
	//fmt.Println(reply)

}

func TestContactsFound(t *testing.T) {
	// TODO(@benqi) 使用ES查询
	// usersDOList, _ := s.UsersDAO.SelectByQueryString(request.Q, request.Q, request.Q, request.Q)

	// found := &TLContactsFound{}

	//found.Results = append(found.Results, MakePeer(&TLPeerUser{UserId: 1}))
	//
	//user := &TLUser{}
	//user.Self = true
	//user.Contact = false
	//user.Id = 2
	//user.FirstName = "benqi"
	//user.LastName = "benqi"
	//user.Username = "benqi"
	//user.AccessHash = 1234567890
	//user.Phone = "+86 111 1111 1111"
	//
	//found.Users = append(found.Users, MakeUser(user))
	//fmt.Println(found)
	//
	//// found := MakeContacts_Found(tlContactsFound)
	//
	//
	//b := found.Encode()
	//fmt.Println(hex.EncodeToString(b))
	//d := NewDecodeBuf(b)
	//
	//c := d.Int()
	//fmt.Println(c)
	//
	//found2 := &TLContactsFound{}
	//err := found2.Decode(d)
	//if err != nil {
	//	fmt.Println("decode2 error: ", err)
	//}
	//
	//fmt.Println(found2)
}

// contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
func (m *TLContactsFound) Encode2() []byte {
	x := NewEncodeBuf(512)
	//x.Int(int32(TLConstructor_CRC32_contacts_found))
	//
	//// x.VectorMessage(m.Results);
	//x1 := make([]byte, 8)
	//binary.LittleEndian.PutUint32(x1, uint32(TLConstructor_CRC32_vector))
	//binary.LittleEndian.PutUint32(x1[4:], uint32(len(m.Results)))
	//x.Bytes(x1)
	//
	//for _, v := range m.Results {
	//	x.buf = append(x.buf, (*v).Encode()...)
	//}
	//// x.VectorMessage(m.Chats);
	//x2 := make([]byte, 8)
	//binary.LittleEndian.PutUint32(x2, uint32(TLConstructor_CRC32_vector))
	//binary.LittleEndian.PutUint32(x2[4:], uint32(len(m.Chats)))
	//x.Bytes(x2)
	//
	//for _, v := range m.Chats {
	//	x.buf = append(x.buf, (*v).Encode()...)
	//}
	//// x.VectorMessage(m.Users);
	//x3 := make([]byte, 8)
	//binary.LittleEndian.PutUint32(x3, uint32(TLConstructor_CRC32_vector))
	//binary.LittleEndian.PutUint32(x3[4:], uint32(len(m.Users)))
	//x.Bytes(x3)
	//
	//for _, v := range m.Users {
	//	x.buf = append(x.buf, (*v).Encode()...)
	//}
	return x.buf
}

func (m *TLContactsFound) Decode2(dbuf *DecodeBuf) error {
	//// x.VectorMessage(m.Results);
	//cc0 := dbuf.Int()
	//fmt.Println("cc0: ", cc0)
	//
	//c1 := dbuf.Int()
	//fmt.Println("c1: ", c1)
	//
	//if c1 != int32(TLConstructor_CRC32_vector) {
	//	return fmt.Errorf("Not vector, classID: %d", c1)
	//}
	//l1 := dbuf.Int()
	//m.Results = make([]*Peer, l1)
	//for i := 0; i < int(l1); i++ {
	//	m.Results[i] = &Peer{}
	//	// dbuf.Int()
	//	err := (*m.Results[i]).Decode(dbuf)
	//	if err != nil {
	//		return fmt.Errorf("Decode 1 , error")
	//	}
	//}
	//// x.VectorMessage(m.Chats);
	//// dbuf.Int()
	//c2 := dbuf.Int()
	//if c2 != int32(TLConstructor_CRC32_vector) {
	//	return fmt.Errorf("Not vector, classID: ", c2)
	//}
	//l2 := dbuf.Int()
	//m.Chats = make([]*Chat, l2)
	//for i := 0; i < int(l2); i++ {
	//	m.Chats[i] = &Chat{}
	//	// dbuf.Int()
	//	err := (*m.Chats[i]).Decode(dbuf)
	//	if err != nil {
	//		return fmt.Errorf("Decode 2 , error")
	//	}
	//}
	//// x.VectorMessage(m.Users);
	//// dbuf.Int()
	//c3 := dbuf.Int()
	//if c3 != int32(TLConstructor_CRC32_vector) {
	//	return fmt.Errorf("Not vector, classID: ", c3)
	//}
	//l3 := dbuf.Int()
	//m.Users = make([]*User, l3)
	//for i := 0; i < int(l3); i++ {
	//	m.Users[i] = &User{}
	//	// dbuf.Int()
	//	err := (*m.Users[i]).Decode(dbuf)
	//	if err != nil {
	//		return fmt.Errorf("Decode 3 , error")
	//	}
	//}
	return dbuf.err
}

// import (
// )

// req_pq#60469778 nonce:int128 = ResPQ;
//type TLReqPq struct {
//	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
//}

/*
func (m []*TLFutureSalt) Encode() []byte {
	x := NewEncodeBuf(512)
	fmt.Println(len(m))
	// x.UInt(CRC32_req_pq)
	return x.buf
}

func (m* TLReqPq) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(CRC32_req_pq)
	x.Bytes(m.GetNonce())
	return x.buf
}

func (m *DecodeBuf) Object2() (r proto.Message) {
	constructor := m.UInt()
	if m.err != nil {
		return nil
	}

	switch constructor {

	case CRC32_req_pq:
		r = &TLReqPq{
			Nonce: m.Bytes(16),
		}

	default:
		r = nil
	}
	return
}

func TestTLReqPqEncode(t *testing.T) {
	m := TLReqPq{
		Nonce: GenerateNonce(16),
	}

	b := m.encode()
	fmt.Println(hex.EncodeToString(b))
	fmt.Println(m.String())

	dbuf := NewDecodeBuf(b)
	m2 := dbuf.Object2()

	reqPq, _ := m2.(*TLReqPq)
	fmt.Println(reqPq.String())
}
*/
