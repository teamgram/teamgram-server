package core

import (
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/userupdates"
)

const MaxUserOperationBatchSize = 100

func (c *UserupdatesCore) UserupdatesProcessUserOperationBatch(in *userupdates.TLUserupdatesProcessUserOperationBatch) (*userupdates.VectorUserOperationResult, error) {
	if in == nil || len(in.Operations) == 0 {
		return &userupdates.VectorUserOperationResult{}, nil
	}
	if len(in.Operations) > MaxUserOperationBatchSize {
		return nil, fmt.Errorf("%w: user operation batch size %d exceeds limit %d", userupdates.ErrOperationTerminal, len(in.Operations), MaxUserOperationBatchSize)
	}

	inputs := make([]repository.ApplyUserOperationInput, 0, len(in.Operations))
	for _, op := range in.Operations {
		apply, err := operationToApplyInput(op)
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, apply)
	}

	results, err := c.svcCtx.Repo.ApplyUserOperationBatch(c.ctx, inputs)
	if err != nil {
		return nil, err
	}
	out := make([]userupdates.UserOperationResultClazz, 0, len(results))
	for i := range results {
		result, err := applyResultToTL(&results[i])
		if err != nil {
			return nil, err
		}
		out = append(out, result)
	}
	if c.svcCtx.PushOutboxNotifier != nil {
		c.svcCtx.PushOutboxNotifier.Wake()
	}
	return &userupdates.VectorUserOperationResult{Datas: out}, nil
}
