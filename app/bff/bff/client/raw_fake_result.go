package bffproxyclient

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/iface"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TryReturnRawFakeRpcResult(ctx context.Context, md *metadata.RpcMetadata, payload []byte) ([]byte, bool, error) {
	_ = ctx

	constructorID, err := bin.NewDecoder(payload).PeekClazzID()
	if err != nil {
		return nil, false, fmt.Errorf("raw fake result constructor: %w", err)
	}
	if !iface.CheckClazzID(constructorID) {
		return nil, false, nil
	}
	obj, err := iface.DecodeObject(bin.NewDecoder(payload))
	if err != nil {
		return nil, true, fmt.Errorf("decode raw fake request: %w", err)
	}
	result, err := tryReturnMetadataRawFakeRpcResult(md, obj)
	if result == nil && err == nil {
		result, err = new(BFFProxyClient2).TryReturnFakeRpcResult(obj)
	}
	if err != nil {
		return nil, true, err
	}
	if result == nil {
		return nil, false, nil
	}
	x := bin.NewEncoder()
	defer x.End()
	layer := int32(224)
	if md != nil && md.Layer > 0 {
		layer = md.Layer
	}
	if err := result.Encode(x, layer); err != nil {
		x.Reset()
		if err2 := result.Encode(x, 0); err2 != nil {
			return nil, true, fmt.Errorf("encode raw fake result: %w", err)
		}
	}
	return append([]byte(nil), x.Bytes()...), true, nil
}

func tryReturnMetadataRawFakeRpcResult(md *metadata.RpcMetadata, obj iface.TLObject) (iface.TLObject, error) {
	if md == nil || md.UserId <= 0 {
		return nil, nil
	}
	switch obj.(type) {
	case *tg.TLUsersGetFullUser:
		firstName := "Teamgram"
		return tg.MakeTLUsersUserFull(&tg.TLUsersUserFull{
			FullUser: tg.MakeTLUserFull(&tg.TLUserFull{
				Id:             md.UserId,
				Settings:       tg.MakeTLPeerSettings(&tg.TLPeerSettings{}).ToPeerSettings(),
				NotifySettings: tg.MakeTLPeerNotifySettings(&tg.TLPeerNotifySettings{}).ToPeerNotifySettings(),
			}).ToUserFull(),
			Chats: []tg.ChatClazz{},
			Users: []tg.UserClazz{
				tg.MakeTLUser(&tg.TLUser{
					Self:      true,
					Id:        md.UserId,
					FirstName: &firstName,
					Usernames: []tg.UsernameClazz{},
				}),
			},
		}).ToUsersUserFull(), nil
	default:
		return nil, nil
	}
}
