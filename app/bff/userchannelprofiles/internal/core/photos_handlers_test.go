package core

import (
	"context"
	"errors"
	"testing"

	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func TestPhotosUpdateProfilePhoto(t *testing.T) {
	var gotUpdate *userpb.TLUserUpdateProfilePhoto
	var gotPhoto *mediapb.TLMediaGetPhoto
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		updateProfilePhoto: func(_ context.Context, in *userpb.TLUserUpdateProfilePhoto) (*tg.Int64, error) {
			gotUpdate = in
			return &tg.Int64{V: 333}, nil
		},
	}, &fakeMediaClient{
		getPhoto: func(_ context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error) {
			gotPhoto = in
			return photoFixture(in.PhotoId), nil
		},
	}, 1001)
	got, err := core.PhotosUpdateProfilePhoto(&tg.TLPhotosUpdateProfilePhoto{Id: tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 222})})
	if err != nil {
		t.Fatalf("PhotosUpdateProfilePhoto returned error: %v", err)
	}
	if gotUpdate == nil || gotUpdate.UserId != 1001 || gotUpdate.Id != 222 {
		t.Fatalf("update request = %+v", gotUpdate)
	}
	if gotPhoto == nil || gotPhoto.PhotoId != 333 {
		t.Fatalf("photo request = %+v", gotPhoto)
	}
	if got.Photo == nil || len(got.Users) != 0 {
		t.Fatalf("response = %+v", got)
	}
}

func TestPhotosDeletePhotosReturnsDeletedIDs(t *testing.T) {
	var got *userpb.TLUserDeleteProfilePhotos
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		deleteProfilePhotos: func(_ context.Context, in *userpb.TLUserDeleteProfilePhotos) (*tg.Int64, error) {
			got = in
			return &tg.Int64{V: 0}, nil
		},
	}, &fakeMediaClient{}, 1001)
	out, err := core.PhotosDeletePhotos(&tg.TLPhotosDeletePhotos{Id: []tg.InputPhotoClazz{
		tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 10}),
		tg.MakeTLInputPhoto(&tg.TLInputPhoto{Id: 20}),
	}})
	if err != nil {
		t.Fatalf("PhotosDeletePhotos returned error: %v", err)
	}
	if got == nil || got.UserId != 1001 || len(got.Id) != 2 || got.Id[1] != 20 {
		t.Fatalf("delete request = %+v", got)
	}
	if len(out.Datas) != 2 || out.Datas[0] != 10 || out.Datas[1] != 20 {
		t.Fatalf("deleted ids = %+v", out.Datas)
	}
}

func TestPhotosGetUserPhotosSkipsPerPhotoErrors(t *testing.T) {
	core := newUserChannelProfilesCoreForTest(&fakeUserClient{
		getProfilePhotos: func(_ context.Context, in *userpb.TLUserGetProfilePhotos) (*userpb.VectorLong, error) {
			if in.UserId != 2002 {
				t.Fatalf("profile photos user id = %d, want 2002", in.UserId)
			}
			return &userpb.VectorLong{Datas: []int64{1, 2, 3}}, nil
		},
	}, &fakeMediaClient{
		getPhoto: func(_ context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error) {
			if in.PhotoId == 2 {
				return nil, errors.New("missing photo")
			}
			return photoFixture(in.PhotoId), nil
		},
	}, 1001)
	got, err := core.PhotosGetUserPhotos(&tg.TLPhotosGetUserPhotos{UserId: tg.MakeTLInputUser(&tg.TLInputUser{UserId: 2002})})
	if err != nil {
		t.Fatalf("PhotosGetUserPhotos returned error: %v", err)
	}
	photos, ok := got.Clazz.(*tg.TLPhotosPhotos)
	if !ok || len(photos.Photos) != 2 || len(photos.Users) != 0 {
		t.Fatalf("photos response = %#v", got)
	}
}
