package core

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/internal/svc"
	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
)

func TestProcessHandlersReturnInvalidArgumentUntilImplemented(t *testing.T) {
	c := New(context.Background(), &svc.ServiceContext{})

	tests := []struct {
		name string
		call func() error
	}{
		{
			name: "photo",
			call: func() error {
				_, err := c.MediaProcessorProcessPhoto(&mediaprocessor.TLMediaProcessorProcessPhoto{})
				return err
			},
		},
		{
			name: "gif",
			call: func() error {
				_, err := c.MediaProcessorProcessGif(&mediaprocessor.TLMediaProcessorProcessGif{})
				return err
			},
		},
		{
			name: "mp4",
			call: func() error {
				_, err := c.MediaProcessorProcessMp4(&mediaprocessor.TLMediaProcessorProcessMp4{})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, mediaprocessor.ErrMediaProcessorInvalidArgument) {
				t.Fatalf("got err %v, want ErrMediaProcessorInvalidArgument", err)
			}
		})
	}
}
