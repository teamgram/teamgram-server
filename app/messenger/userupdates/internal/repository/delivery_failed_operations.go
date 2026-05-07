package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/internal/repository/model"
	"github.com/teamgram/teamgram-server/v2/app/messenger/userupdates/payload"
)

func (r *Repository) RecordDeliveryFailure(ctx context.Context, in RecordDeliveryFailureInput) error {
	if in.FailedId == 0 {
		id, err := r.idgen.NextID(ctx)
		if err != nil {
			return err
		}
		in.FailedId = id
	}
	_, _, err := r.models.DeliveryFailedOperationsModel.Insert(ctx, &model.DeliveryFailedOperations{
		FailedId:             in.FailedId,
		UserId:               in.UserId,
		OperationId:          in.OperationId,
		OpType:               in.OpType,
		BucketId:             in.BucketId,
		KafkaTopic:           in.KafkaTopic,
		KafkaPartition:       in.KafkaPartition,
		KafkaOffset:          in.KafkaOffset,
		PayloadSchemaVersion: payload.ReceiverKafkaMessageSchemaVersion,
		PayloadHash:          in.PayloadHash,
		FailureCategory:      in.FailureCategory,
		FailureCode:          in.FailureCode,
		FailureMessage:       safeFailureMessage(in.FailureMessage),
		RetryCount:           0,
		Status:               DeliveryFailedOperationStatusOpen,
		FailedAt:             unixNow(),
	})
	if isDuplicateKey(err) {
		return nil
	}
	return storageError("record delivery failure", err)
}

func isDuplicateKey(err error) bool {
	if err == nil {
		return false
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}
	return strings.Contains(err.Error(), "Duplicate entry")
}

func safeFailureMessage(message string) string {
	if message == "" {
		return ""
	}
	message = strings.ReplaceAll(message, "\n", " ")
	message = strings.ReplaceAll(message, "\r", " ")
	if len(message) > 256 {
		return message[:256]
	}
	return message
}
