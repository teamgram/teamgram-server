package userupdates

import "errors"

var (
	ErrNotOwner                 = errors.New("userupdates: not owner")
	ErrOwnerFenceFailed         = errors.New("userupdates: owner fence failed")
	ErrOperationPayloadConflict = errors.New("userupdates: operation payload conflict")
	ErrPtsContinuityViolation   = errors.New("userupdates: pts continuity violation")
	ErrOperationTerminal        = errors.New("userupdates: operation terminal")
	ErrUserupdatesStorage       = errors.New("userupdates: storage failure")
	ErrDialogQueryTooLarge      = errors.New("userupdates: dialog query too large")
	ErrAuthSeqLedgerUnavailable = errors.New("userupdates: auth seq ledger unavailable")
)
