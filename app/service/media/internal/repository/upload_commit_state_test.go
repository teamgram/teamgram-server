package repository

import (
	"errors"
	"testing"
)

func TestUploadCommitStateValidTransitions(t *testing.T) {
	state := NewUploadCommitState()
	if state.Status() != UploadStatusUploading {
		t.Fatalf("initial status = %q, want %q", state.Status(), UploadStatusUploading)
	}

	if err := state.MarkFinalized(); err != nil {
		t.Fatalf("MarkFinalized() error = %v", err)
	}
	if err := state.MarkMetadataCommitted(); err != nil {
		t.Fatalf("MarkMetadataCommitted() error = %v", err)
	}
	if err := state.MarkPublished(); err != nil {
		t.Fatalf("MarkPublished() error = %v", err)
	}
	if state.Status() != UploadStatusPublished {
		t.Fatalf("status = %q, want %q", state.Status(), UploadStatusPublished)
	}
}

func TestUploadCommitStateRejectsInvalidTransitionOrder(t *testing.T) {
	state := NewUploadCommitState()
	if err := state.MarkMetadataCommitted(); !errors.Is(err, ErrUploadStatusTransitionInvalid) {
		t.Fatalf("MarkMetadataCommitted() error = %v, want ErrUploadStatusTransitionInvalid", err)
	}
	if err := state.MarkPublished(); !errors.Is(err, ErrUploadStatusTransitionInvalid) {
		t.Fatalf("MarkPublished() error = %v, want ErrUploadStatusTransitionInvalid", err)
	}
	if state.Status() != UploadStatusUploading {
		t.Fatalf("status after invalid transitions = %q, want %q", state.Status(), UploadStatusUploading)
	}
}

func TestUploadCommitStateFailedGCPending(t *testing.T) {
	state := NewUploadCommitState()
	if err := state.MarkFailedGCPending(); err != nil {
		t.Fatalf("MarkFailedGCPending() error = %v", err)
	}
	if state.Status() != UploadStatusFailedGCPending {
		t.Fatalf("status = %q, want %q", state.Status(), UploadStatusFailedGCPending)
	}
	if err := state.MarkFinalized(); !errors.Is(err, ErrUploadStatusTransitionInvalid) {
		t.Fatalf("MarkFinalized() after failed_gc_pending error = %v, want ErrUploadStatusTransitionInvalid", err)
	}
}

func TestUploadCommitStateRejectsFailedGCPendingAfterPublished(t *testing.T) {
	state := NewUploadCommitState()
	if err := state.MarkFinalized(); err != nil {
		t.Fatalf("MarkFinalized() error = %v", err)
	}
	if err := state.MarkMetadataCommitted(); err != nil {
		t.Fatalf("MarkMetadataCommitted() error = %v", err)
	}
	if err := state.MarkPublished(); err != nil {
		t.Fatalf("MarkPublished() error = %v", err)
	}

	if err := state.MarkFailedGCPending(); !errors.Is(err, ErrUploadStatusTransitionInvalid) {
		t.Fatalf("MarkFailedGCPending() after published error = %v, want ErrUploadStatusTransitionInvalid", err)
	}
}
