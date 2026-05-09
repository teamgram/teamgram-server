package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

func (c *UserupdatesCore) UserupdatesProcessUserOperationBatch(in *userupdates.TLUserupdatesProcessUserOperationBatch) (*userupdates.VectorUserOperationResult, error) {
	return nil, fmt.Errorf("%w: userupdates.processUserOperationBatch is not wired until Task 3", userupdates.ErrOperationTerminal)
}
