package repository

import "encoding/hex"

func AuthSeqPayloadID(hash []byte) string {
	if len(hash) == 0 {
		return ""
	}
	return hex.EncodeToString(hash)
}
