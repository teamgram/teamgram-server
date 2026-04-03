package user

import (
	"encoding/json"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// These compatibility types keep older generated service/client surfaces buildable
// while the user service schema generation is brought back in sync.

type TLUserSaveMusic struct {
	ClazzID uint32 `json:"_id"`
	Unsave  bool   `json:"unsave,omitempty"`
	UserId  int64  `json:"user_id,omitempty"`
	Id      int64  `json:"id,omitempty"`
	AfterId *int64 `json:"after_id,omitempty"`
}

type TLUserGetSavedMusicIdList struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id,omitempty"`
}

type TLUserSetMainProfileTab struct {
	ClazzID uint32         `json:"_id"`
	UserId  int64          `json:"user_id,omitempty"`
	Tab     *tg.ProfileTab `json:"tab,omitempty"`
}

type TLUserSetDefaultHistoryTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id,omitempty"`
	Ttl     int32  `json:"ttl,omitempty"`
}

type TLUserGetDefaultHistoryTTL struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id,omitempty"`
}

type TLUserGetAccountUsername struct {
	ClazzID uint32 `json:"_id"`
	UserId  int64  `json:"user_id,omitempty"`
}

type TLUserCheckAccountUsername struct {
	ClazzID  uint32 `json:"_id"`
	UserId   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type TLUserGetChannelUsername struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id,omitempty"`
}

type TLUserCheckChannelUsername struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id,omitempty"`
	Username  string `json:"username,omitempty"`
}

type TLUserUpdateUsernameByPeer struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type,omitempty"`
	PeerId   int64  `json:"peer_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type TLUserCheckUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username,omitempty"`
}

type TLUserUpdateUsernameByUsername struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type,omitempty"`
	PeerId   int64  `json:"peer_id,omitempty"`
	Username string `json:"username,omitempty"`
}

type TLUserDeleteUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username,omitempty"`
}

type TLUserResolveUsername struct {
	ClazzID  uint32 `json:"_id"`
	Username string `json:"username,omitempty"`
}

type TLUserGetListByUsernameList struct {
	ClazzID uint32   `json:"_id"`
	Names   []string `json:"names,omitempty"`
}

type TLUserDeleteUsernameByPeer struct {
	ClazzID  uint32 `json:"_id"`
	PeerType int32  `json:"peer_type,omitempty"`
	PeerId   int64  `json:"peer_id,omitempty"`
}

type TLUserSearchUsername struct {
	ClazzID          uint32  `json:"_id"`
	Q                string  `json:"q,omitempty"`
	ExcludedContacts []int64 `json:"excluded_contacts,omitempty"`
	Limit            int32   `json:"limit,omitempty"`
}

type TLUserToggleUsername struct {
	ClazzID  uint32   `json:"_id"`
	PeerType int32    `json:"peer_type,omitempty"`
	PeerId   int64    `json:"peer_id,omitempty"`
	Username string   `json:"username,omitempty"`
	Active   *tg.Bool `json:"active,omitempty"`
}

type TLUserReorderUsernames struct {
	ClazzID      uint32   `json:"_id"`
	PeerType     int32    `json:"peer_type,omitempty"`
	PeerId       int64    `json:"peer_id,omitempty"`
	UsernameList []string `json:"username_list,omitempty"`
}

type TLUserDeactivateAllChannelUsernames struct {
	ClazzID   uint32 `json:"_id"`
	ChannelId int64  `json:"channel_id,omitempty"`
}

type UsernameData struct {
	ClazzID  uint32   `json:"_id"`
	Username string   `json:"username,omitempty"`
	Peer     *tg.Peer `json:"peer,omitempty"`
	Editable bool     `json:"editable,omitempty"`
	Active   bool     `json:"active,omitempty"`
}

type UsernameExist struct {
	ClazzID uint32 `json:"_id"`
}

type VectorUsernameData struct {
	ClazzID uint32          `json:"_id"`
	Datas   []*UsernameData `json:"datas,omitempty"`
}

func (m *TLUserSaveMusic) String() string                     { return compatString(m) }
func (m *TLUserGetSavedMusicIdList) String() string           { return compatString(m) }
func (m *TLUserSetMainProfileTab) String() string             { return compatString(m) }
func (m *TLUserSetDefaultHistoryTTL) String() string          { return compatString(m) }
func (m *TLUserGetDefaultHistoryTTL) String() string          { return compatString(m) }
func (m *TLUserGetAccountUsername) String() string            { return compatString(m) }
func (m *TLUserCheckAccountUsername) String() string          { return compatString(m) }
func (m *TLUserGetChannelUsername) String() string            { return compatString(m) }
func (m *TLUserCheckChannelUsername) String() string          { return compatString(m) }
func (m *TLUserUpdateUsernameByPeer) String() string          { return compatString(m) }
func (m *TLUserCheckUsername) String() string                 { return compatString(m) }
func (m *TLUserUpdateUsernameByUsername) String() string      { return compatString(m) }
func (m *TLUserDeleteUsername) String() string                { return compatString(m) }
func (m *TLUserResolveUsername) String() string               { return compatString(m) }
func (m *TLUserGetListByUsernameList) String() string         { return compatString(m) }
func (m *TLUserDeleteUsernameByPeer) String() string          { return compatString(m) }
func (m *TLUserSearchUsername) String() string                { return compatString(m) }
func (m *TLUserToggleUsername) String() string                { return compatString(m) }
func (m *TLUserReorderUsernames) String() string              { return compatString(m) }
func (m *TLUserDeactivateAllChannelUsernames) String() string { return compatString(m) }
func (m *UsernameData) String() string                        { return compatString(m) }
func (m *UsernameExist) String() string                       { return compatString(m) }
func (m *VectorUsernameData) String() string                  { return compatString(m) }

func (m *TLUserSaveMusic) Encode(x *bin.Encoder, layer int32) error { return compatEncodeJSON(x, m) }
func (m *TLUserGetSavedMusicIdList) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserSetMainProfileTab) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserSetDefaultHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserGetDefaultHistoryTTL) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserGetAccountUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserCheckAccountUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserGetChannelUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserCheckChannelUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserUpdateUsernameByPeer) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserCheckUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserUpdateUsernameByUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserDeleteUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserResolveUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserGetListByUsernameList) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserDeleteUsernameByPeer) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserSearchUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserToggleUsername) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserReorderUsernames) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *TLUserDeactivateAllChannelUsernames) Encode(x *bin.Encoder, layer int32) error {
	return compatEncodeJSON(x, m)
}
func (m *UsernameData) Encode(x *bin.Encoder, layer int32) error       { return compatEncodeJSON(x, m) }
func (m *UsernameExist) Encode(x *bin.Encoder, layer int32) error      { return compatEncodeJSON(x, m) }
func (m *VectorUsernameData) Encode(x *bin.Encoder, layer int32) error { return compatEncodeJSON(x, m) }

func (m *TLUserSaveMusic) Decode(d *bin.Decoder) error                { return compatDecodeJSON(d, m) }
func (m *TLUserGetSavedMusicIdList) Decode(d *bin.Decoder) error      { return compatDecodeJSON(d, m) }
func (m *TLUserSetMainProfileTab) Decode(d *bin.Decoder) error        { return compatDecodeJSON(d, m) }
func (m *TLUserSetDefaultHistoryTTL) Decode(d *bin.Decoder) error     { return compatDecodeJSON(d, m) }
func (m *TLUserGetDefaultHistoryTTL) Decode(d *bin.Decoder) error     { return compatDecodeJSON(d, m) }
func (m *TLUserGetAccountUsername) Decode(d *bin.Decoder) error       { return compatDecodeJSON(d, m) }
func (m *TLUserCheckAccountUsername) Decode(d *bin.Decoder) error     { return compatDecodeJSON(d, m) }
func (m *TLUserGetChannelUsername) Decode(d *bin.Decoder) error       { return compatDecodeJSON(d, m) }
func (m *TLUserCheckChannelUsername) Decode(d *bin.Decoder) error     { return compatDecodeJSON(d, m) }
func (m *TLUserUpdateUsernameByPeer) Decode(d *bin.Decoder) error     { return compatDecodeJSON(d, m) }
func (m *TLUserCheckUsername) Decode(d *bin.Decoder) error            { return compatDecodeJSON(d, m) }
func (m *TLUserUpdateUsernameByUsername) Decode(d *bin.Decoder) error { return compatDecodeJSON(d, m) }
func (m *TLUserDeleteUsername) Decode(d *bin.Decoder) error           { return compatDecodeJSON(d, m) }
func (m *TLUserResolveUsername) Decode(d *bin.Decoder) error          { return compatDecodeJSON(d, m) }
func (m *TLUserGetListByUsernameList) Decode(d *bin.Decoder) error    { return compatDecodeJSON(d, m) }
func (m *TLUserDeleteUsernameByPeer) Decode(d *bin.Decoder) error     { return compatDecodeJSON(d, m) }
func (m *TLUserSearchUsername) Decode(d *bin.Decoder) error           { return compatDecodeJSON(d, m) }
func (m *TLUserToggleUsername) Decode(d *bin.Decoder) error           { return compatDecodeJSON(d, m) }
func (m *TLUserReorderUsernames) Decode(d *bin.Decoder) error         { return compatDecodeJSON(d, m) }
func (m *TLUserDeactivateAllChannelUsernames) Decode(d *bin.Decoder) error {
	return compatDecodeJSON(d, m)
}
func (m *UsernameData) Decode(d *bin.Decoder) error       { return compatDecodeJSON(d, m) }
func (m *UsernameExist) Decode(d *bin.Decoder) error      { return compatDecodeJSON(d, m) }
func (m *VectorUsernameData) Decode(d *bin.Decoder) error { return compatDecodeJSON(d, m) }

func compatString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func compatEncodeJSON(x *bin.Encoder, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	x.PutBytes(b)
	return nil
}

func compatDecodeJSON(d *bin.Decoder, v interface{}) error {
	b, err := d.Bytes()
	if err != nil {
		return nil
	}
	if len(b) == 0 {
		return nil
	}
	return json.Unmarshal(b, v)
}

var (
	_ fmt.Stringer = (*TLUserSaveMusic)(nil)
	_ fmt.Stringer = (*UsernameData)(nil)
)
