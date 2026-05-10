package repository

import (
	"errors"
	"fmt"
)

var ErrUploadStatusTransitionInvalid = errors.New("media: invalid upload status transition")

type UploadStatus string

const (
	UploadStatusUploading         UploadStatus = "uploading"
	UploadStatusFinalized         UploadStatus = "finalized"
	UploadStatusMetadataCommitted UploadStatus = "metadata_committed"
	UploadStatusPublished         UploadStatus = "published"
	UploadStatusFailedGCPending   UploadStatus = "failed_gc_pending"
)

type UploadCommitState struct {
	status UploadStatus
}

func NewUploadCommitState() *UploadCommitState {
	return &UploadCommitState{status: UploadStatusUploading}
}

func (s *UploadCommitState) Status() UploadStatus {
	if s == nil || s.status == "" {
		return UploadStatusUploading
	}
	return s.status
}

func (s *UploadCommitState) MarkFinalized() error {
	return s.transition(UploadStatusUploading, UploadStatusFinalized)
}

func (s *UploadCommitState) MarkMetadataCommitted() error {
	return s.transition(UploadStatusFinalized, UploadStatusMetadataCommitted)
}

func (s *UploadCommitState) MarkPublished() error {
	return s.transition(UploadStatusMetadataCommitted, UploadStatusPublished)
}

func (s *UploadCommitState) MarkFailedGCPending() error {
	status := s.Status()
	if status == UploadStatusPublished || status == UploadStatusFailedGCPending {
		return invalidUploadStatusTransition(status, UploadStatusFailedGCPending)
	}
	s.status = UploadStatusFailedGCPending
	return nil
}

func (s *UploadCommitState) transition(from, to UploadStatus) error {
	if s == nil {
		return invalidUploadStatusTransition("", to)
	}
	if status := s.Status(); status != from {
		return invalidUploadStatusTransition(status, to)
	}
	s.status = to
	return nil
}

func invalidUploadStatusTransition(from, to UploadStatus) error {
	return fmt.Errorf("%w: %s -> %s", ErrUploadStatusTransitionInvalid, from, to)
}
