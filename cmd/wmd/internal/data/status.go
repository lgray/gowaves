package data

import (
	"github.com/wavesplatform/gowaves/pkg/proto"
)

//easyjson:json
type WMDStatus struct {
	CurrentHeight int           `json:"current_height"`
	LastBlockID   proto.BlockID `json:"last_block_id"`
}
