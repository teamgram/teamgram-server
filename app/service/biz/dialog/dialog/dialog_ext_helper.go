// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dialog

type (
	DialogExtList []*DialogExt
)

// sort
func (m DialogExtList) Len() int {
	return len(m)
}
func (m DialogExtList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m DialogExtList) Less(i, j int) bool {
	// TODO(@benqi): if date[i] == date[j]
	return m[i].Order < m[j].Order
}

type (
	DialogPinnedExtList []*DialogPinnedExt
)

func (m DialogPinnedExtList) Add(peerType int32, peerId int64, order int64) DialogPinnedExtList {
	return append(m, &DialogPinnedExt{
		Order:    order,
		PeerType: peerType,
		PeerId:   peerId,
	})
}

func (m DialogPinnedExtList) Len() int {
	return len(m)
}
func (m DialogPinnedExtList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m DialogPinnedExtList) Less(i, j int) bool {
	// TODO(@benqi): if date[i] == date[j]
	return m[i].Order < m[j].Order
}
