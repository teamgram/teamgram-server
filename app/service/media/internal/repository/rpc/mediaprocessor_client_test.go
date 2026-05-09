package rpc

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/service/mediaprocessor/mediaprocessor"
)

type fakeMediaProcessorGeneratedClient struct {
	photoReq *mediaprocessor.TLMediaProcessorProcessPhoto
	gifReq   *mediaprocessor.TLMediaProcessorProcessGif
	mp4Req   *mediaprocessor.TLMediaProcessorProcessMp4
	photo    *mediaprocessor.ProcessedPhoto
	document *mediaprocessor.ProcessedDocument
	err      error
}

func (f *fakeMediaProcessorGeneratedClient) MediaProcessorProcessPhoto(_ context.Context, in *mediaprocessor.TLMediaProcessorProcessPhoto) (*mediaprocessor.ProcessedPhoto, error) {
	f.photoReq = in
	return f.photo, f.err
}

func (f *fakeMediaProcessorGeneratedClient) MediaProcessorProcessGif(_ context.Context, in *mediaprocessor.TLMediaProcessorProcessGif) (*mediaprocessor.ProcessedDocument, error) {
	f.gifReq = in
	return f.document, f.err
}

func (f *fakeMediaProcessorGeneratedClient) MediaProcessorProcessMp4(_ context.Context, in *mediaprocessor.TLMediaProcessorProcessMp4) (*mediaprocessor.ProcessedDocument, error) {
	f.mp4Req = in
	return f.document, f.err
}

func TestMediaProcessorClientForwardsProcessPhoto(t *testing.T) {
	photo := &mediaprocessor.ProcessedPhoto{OriginalObjectId: "object-photo"}
	fake := &fakeMediaProcessorGeneratedClient{photo: photo}
	client := NewMediaProcessorClient(fake)
	req := &mediaprocessor.TLMediaProcessorProcessPhoto{OwnerId: 1001}

	got, err := client.ProcessPhoto(context.Background(), req)
	if err != nil {
		t.Fatalf("ProcessPhoto() error = %v", err)
	}
	if got != photo {
		t.Fatalf("ProcessPhoto() result = %#v, want forwarded photo", got)
	}
	if fake.photoReq != req {
		t.Fatalf("forwarded request = %#v, want original", fake.photoReq)
	}
}

func TestMediaProcessorClientForwardsProcessGifAndMp4(t *testing.T) {
	document := &mediaprocessor.ProcessedDocument{ObjectId: "object-document"}
	fake := &fakeMediaProcessorGeneratedClient{document: document}
	client := NewMediaProcessorClient(fake)
	gifReq := &mediaprocessor.TLMediaProcessorProcessGif{OwnerId: 1001}
	mp4Req := &mediaprocessor.TLMediaProcessorProcessMp4{OwnerId: 1001}

	gotGif, err := client.ProcessGif(context.Background(), gifReq)
	if err != nil {
		t.Fatalf("ProcessGif() error = %v", err)
	}
	gotMp4, err := client.ProcessMp4(context.Background(), mp4Req)
	if err != nil {
		t.Fatalf("ProcessMp4() error = %v", err)
	}
	if gotGif != document || gotMp4 != document {
		t.Fatalf("results = %#v/%#v, want forwarded document", gotGif, gotMp4)
	}
	if fake.gifReq != gifReq {
		t.Fatalf("forwarded gif request = %#v, want original", fake.gifReq)
	}
	if fake.mp4Req != mp4Req {
		t.Fatalf("forwarded mp4 request = %#v, want original", fake.mp4Req)
	}
}
