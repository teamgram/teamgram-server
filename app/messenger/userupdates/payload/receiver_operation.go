package payload

type ReceiverOperationEnvelopeV1 struct {
	UserID        int64
	BucketID      int32
	PartitionID   int32
	OperationID   string
	OpType        int32
	PeerType      int32
	PeerID        int64
	PayloadCodec  int32
	Payload       []byte
	PayloadHash   string
	DependencyPts []int64
}
