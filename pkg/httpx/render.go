// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/pkg/httpx/render"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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
		r, err = protojson.Marshal(data.(proto.Message))
		if err != nil {
			JSONError(w, err)
		}
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
