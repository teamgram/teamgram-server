package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/bff/userchannelprofiles/internal/repository"
	"github.com/teamgram/teamgram-server/v2/app/bff/userchannelprofiles/internal/svc"
	userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"
	userpb "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
	mediaclient "github.com/teamgram/teamgram-server/v2/app/service/media/client"
	mediapb "github.com/teamgram/teamgram-server/v2/app/service/media/media"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

type fakeUserClient struct {
	userclient.UserClient
	getImmutableUser       func(context.Context, *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error)
	updateAbout            func(context.Context, *userpb.TLUserUpdateAbout) (*tg.Bool, error)
	updateFirstAndLastName func(context.Context, *userpb.TLUserUpdateFirstAndLastName) (*tg.Bool, error)
	updateLastSeen         func(context.Context, *userpb.TLUserUpdateLastSeen) (*tg.Bool, error)
	updateBirthday         func(context.Context, *userpb.TLUserUpdateBirthday) (*tg.Bool, error)
	updatePersonalChannel  func(context.Context, *userpb.TLUserUpdatePersonalChannel) (*tg.Bool, error)
	setMainProfileTab      func(context.Context, *userpb.TLUserSetMainProfileTab) (*tg.Bool, error)
	saveMusic              func(context.Context, *userpb.TLUserSaveMusic) (*tg.Bool, error)
	getSavedMusicIDList    func(context.Context, *userpb.TLUserGetSavedMusicIdList) (*userpb.VectorLong, error)
	getBirthdays           func(context.Context, *userpb.TLUserGetBirthdays) (*userpb.VectorContactBirthday, error)
	getMutableUsersV2      func(context.Context, *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error)
	getUserProjection      func(context.Context, *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error)
	updateProfilePhoto     func(context.Context, *userpb.TLUserUpdateProfilePhoto) (*tg.Int64, error)
	deleteProfilePhotos    func(context.Context, *userpb.TLUserDeleteProfilePhotos) (*tg.Int64, error)
	getProfilePhotos       func(context.Context, *userpb.TLUserGetProfilePhotos) (*userpb.VectorLong, error)
}

type fakeMediaClient struct {
	mediaclient.MediaClient
	uploadProfilePhotoFile func(context.Context, *mediapb.TLMediaUploadProfilePhotoFile) (*tg.Photo, error)
	getPhoto               func(context.Context, *mediapb.TLMediaGetPhoto) (*tg.Photo, error)
	getDocumentList        func(context.Context, *mediapb.TLMediaGetDocumentList) (*mediapb.VectorDocument, error)
}

func (f *fakeUserClient) UserGetImmutableUser(ctx context.Context, in *userpb.TLUserGetImmutableUser) (*tg.ImmutableUser, error) {
	return f.getImmutableUser(ctx, in)
}

func (f *fakeUserClient) UserUpdateAbout(ctx context.Context, in *userpb.TLUserUpdateAbout) (*tg.Bool, error) {
	return f.updateAbout(ctx, in)
}

func (f *fakeUserClient) UserUpdateFirstAndLastName(ctx context.Context, in *userpb.TLUserUpdateFirstAndLastName) (*tg.Bool, error) {
	return f.updateFirstAndLastName(ctx, in)
}

func (f *fakeUserClient) UserUpdateLastSeen(ctx context.Context, in *userpb.TLUserUpdateLastSeen) (*tg.Bool, error) {
	return f.updateLastSeen(ctx, in)
}

func (f *fakeUserClient) UserUpdateBirthday(ctx context.Context, in *userpb.TLUserUpdateBirthday) (*tg.Bool, error) {
	return f.updateBirthday(ctx, in)
}

func (f *fakeUserClient) UserUpdatePersonalChannel(ctx context.Context, in *userpb.TLUserUpdatePersonalChannel) (*tg.Bool, error) {
	return f.updatePersonalChannel(ctx, in)
}

func (f *fakeUserClient) UserSetMainProfileTab(ctx context.Context, in *userpb.TLUserSetMainProfileTab) (*tg.Bool, error) {
	return f.setMainProfileTab(ctx, in)
}

func (f *fakeUserClient) UserSaveMusic(ctx context.Context, in *userpb.TLUserSaveMusic) (*tg.Bool, error) {
	return f.saveMusic(ctx, in)
}

func (f *fakeUserClient) UserGetSavedMusicIdList(ctx context.Context, in *userpb.TLUserGetSavedMusicIdList) (*userpb.VectorLong, error) {
	return f.getSavedMusicIDList(ctx, in)
}

func (f *fakeUserClient) UserGetBirthdays(ctx context.Context, in *userpb.TLUserGetBirthdays) (*userpb.VectorContactBirthday, error) {
	return f.getBirthdays(ctx, in)
}

func (f *fakeUserClient) UserGetMutableUsersV2(ctx context.Context, in *userpb.TLUserGetMutableUsersV2) (*tg.MutableUsers, error) {
	return f.getMutableUsersV2(ctx, in)
}

func (f *fakeUserClient) UserGetUserProjectionBundle(ctx context.Context, in *userpb.TLUserGetUserProjectionBundle) (*userpb.UserProjectionBundle, error) {
	return f.getUserProjection(ctx, in)
}

func (f *fakeUserClient) UserUpdateProfilePhoto(ctx context.Context, in *userpb.TLUserUpdateProfilePhoto) (*tg.Int64, error) {
	return f.updateProfilePhoto(ctx, in)
}

func (f *fakeUserClient) UserDeleteProfilePhotos(ctx context.Context, in *userpb.TLUserDeleteProfilePhotos) (*tg.Int64, error) {
	return f.deleteProfilePhotos(ctx, in)
}

func (f *fakeUserClient) UserGetProfilePhotos(ctx context.Context, in *userpb.TLUserGetProfilePhotos) (*userpb.VectorLong, error) {
	return f.getProfilePhotos(ctx, in)
}

func (f *fakeMediaClient) MediaUploadProfilePhotoFile(ctx context.Context, in *mediapb.TLMediaUploadProfilePhotoFile) (*tg.Photo, error) {
	return f.uploadProfilePhotoFile(ctx, in)
}

func (f *fakeMediaClient) MediaGetPhoto(ctx context.Context, in *mediapb.TLMediaGetPhoto) (*tg.Photo, error) {
	return f.getPhoto(ctx, in)
}

func (f *fakeMediaClient) MediaGetDocumentList(ctx context.Context, in *mediapb.TLMediaGetDocumentList) (*mediapb.VectorDocument, error) {
	return f.getDocumentList(ctx, in)
}

func newUserChannelProfilesCoreForTest(userClient userclient.UserClient, mediaClient mediaclient.MediaClient, selfID int64) *UserChannelProfilesCore {
	c := New(context.Background(), &svc.ServiceContext{
		Repo: &repository.Repository{UserClient: userClient, MediaClient: mediaClient},
	})
	c.MD = &metadata.RpcMetadata{UserId: selfID, PermAuthKeyId: 9001}
	return c
}

func immutableUserFixture(id int64, firstName string, lastName string, username string) *tg.ImmutableUser {
	return tg.MakeTLImmutableUser(&tg.TLImmutableUser{
		User: tg.MakeTLUserData(&tg.TLUserData{
			Id:         id,
			AccessHash: id * 10,
			FirstName:  firstName,
			LastName:   lastName,
			Username:   username,
		}),
	})
}

func photoFixture(id int64) *tg.Photo {
	return &tg.Photo{Clazz: tg.MakeTLPhoto(&tg.TLPhoto{Id: id, DcId: 2, Sizes: []tg.PhotoSizeClazz{}})}
}

func documentFixture(id int64) tg.DocumentClazz {
	return tg.MakeTLDocument(&tg.TLDocument{Id: id, DcId: 2})
}
