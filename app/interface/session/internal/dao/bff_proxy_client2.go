// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: Benqi (wubenqi@gmail.com)

package dao

import (
	"context"

	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/rpc/metadata"

	_ "github.com/teamgram/teamgram-server/v2/app/bff/configuration/configuration/configurationservice"
	_ "github.com/teamgram/teamgram-server/v2/app/bff/miscellaneous/miscellaneous/miscellaneousservice"
)

func (d *Dao) InvokeContext(ctx context.Context, rpcMetaData *metadata.RpcMetadata, object iface.TLObject) (iface.TLObject, error) {
	return d.BFFProxyClient2.InvokeContext(ctx, rpcMetaData, object)
}
