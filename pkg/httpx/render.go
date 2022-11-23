// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package httpx

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/pkg/httpx/render"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// JSONError
// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func JSONError(w http.ResponseWriter, err error) {
	if rErr, ok := err.(*mtproto.TLRpcError); ok {
		httpx.WriteJson(w, http.StatusOK, render.JSON{
			Ok:          false,
			ErrorCode:   rErr.Code(),
			Description: rErr.Message(),
		})
	} else {
		httpx.WriteJson(w, http.StatusOK, render.JSON{
			Ok:          false,
			ErrorCode:   500,
			Description: err.Error(),
		})
	}
}

func JSON(w http.ResponseWriter, data interface{}) {
	// r, err := jsonx.Marshal(data)
	var (
		err error
		r   []byte
	)
	switch data.(type) {
	case proto.Message:
		m := &jsonpb.Marshaler{
			OrigName: true,
		}
		var buf bytes.Buffer
		if err = m.Marshal(&buf, data.(proto.Message)); err != nil {
			JSONError(w, err)
		}
		r = buf.Bytes()
	default:
		r, err = json.Marshal(data)
	}

	if err != nil {
		JSONError(w, err)
	} else {
		httpx.WriteJson(w, http.StatusOK, render.JSON{
			Ok:     true,
			Result: r,
		})
	}
}
