package mtproto

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/teamgram/teamgram-server/v2/pkg/proto/bin"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type WrapperMetadata struct {
	Layer          int32
	Client         string
	ApiId          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          string
	Params         string
}

func UnwrapClientRPC(payload []byte) (inner []byte, md WrapperMetadata, err error) {
	d := bin.NewDecoder(payload)
	clazzID, err := d.PeekClazzID()
	if err != nil {
		return nil, WrapperMetadata{}, fmt.Errorf("unwrap client rpc constructor: %w", err)
	}

	switch clazzID {
	case tg.ClazzID_invokeWithLayer:
		invoke := &tg.TLInvokeWithLayer{}
		if err := invoke.Decode(bin.NewDecoder(payload)); err != nil {
			return nil, WrapperMetadata{}, fmt.Errorf("unwrap invokeWithLayer: %w", err)
		}
		inner, childMD, err := UnwrapClientRPC(invoke.Query)
		if err != nil {
			return nil, WrapperMetadata{}, err
		}
		childMD.Layer = invoke.Layer
		return inner, childMD, nil
	case tg.ClazzID_initConnection_c1cd5ea9, tg.ClazzID_initConnection_785188b8:
		initConn := &tg.TLInitConnection{}
		if err := initConn.Decode(bin.NewDecoder(payload)); err != nil {
			return nil, WrapperMetadata{}, fmt.Errorf("unwrap initConnection: %w", err)
		}
		return initConn.Query, WrapperMetadata{
			Client:         strings.TrimSpace(initConn.DeviceModel + " " + initConn.SystemVersion + " " + initConn.AppVersion),
			ApiId:          initConn.ApiId,
			DeviceModel:    initConn.DeviceModel,
			SystemVersion:  initConn.SystemVersion,
			AppVersion:     initConn.AppVersion,
			SystemLangCode: initConn.SystemLangCode,
			LangPack:       initConn.LangPack,
			LangCode:       initConn.LangCode,
			Proxy:          encodeWrapperJSON(initConn.Proxy),
			Params:         encodeWrapperJSON(initConn.Params),
		}, nil
	default:
		return payload, WrapperMetadata{}, nil
	}
}

func encodeWrapperJSON(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
