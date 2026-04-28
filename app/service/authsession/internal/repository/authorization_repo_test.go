package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/internal/repository/model"
)

type fakeGeoipClient struct {
	region *geoip.Region
	err    error
}

func (f fakeGeoipClient) GeoipGetCountryAndRegionByIp(context.Context, *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error) {
	return f.region, f.err
}

func TestAuthorizationCurrentHashIsZero(t *testing.T) {
	auth := authDataToAuthorization(&cacheAuthData{
		Client:   &clientSessionCacheData{Ip: "127.0.0.1"},
		BindUser: &bindUser{Hash: 12345},
	}, true, "US", "California")
	if auth.Hash != 0 {
		t.Fatalf("Hash = %d, want 0", auth.Hash)
	}
	if !auth.Current {
		t.Fatal("Current = false, want true")
	}
}

func TestAuthorizationNonCurrentPreservesHash(t *testing.T) {
	auth := authDataToAuthorization(&cacheAuthData{
		Client:   &clientSessionCacheData{Ip: "127.0.0.1"},
		BindUser: &bindUser{Hash: 12345},
	}, false, "US", "California")
	if auth.Hash != 12345 {
		t.Fatalf("Hash = %d, want 12345", auth.Hash)
	}
}

func TestAuthorizationGeoipMapping(t *testing.T) {
	r := &Repository{
		geoipClient: fakeGeoipClient{region: &geoip.Region{IsoCode: "US", Region: "California"}},
	}
	country, region := r.getCountryAndRegionByIP(context.Background(), "127.0.0.1")
	if country != "US" || region != "California" {
		t.Fatalf("geoip mapping = (%q, %q), want (US, California)", country, region)
	}
}

func TestGeoipLookupDegradesOnError(t *testing.T) {
	r := &Repository{
		geoipClient: fakeGeoipClient{err: errors.New("boom")},
	}
	country, region := r.getCountryAndRegionByIP(context.Background(), "127.0.0.1")
	if country != "" || region != "" {
		t.Fatalf("geoip mapping = (%q, %q), want empty values", country, region)
	}
}

type fakeAuthUsersModel struct {
	model.AuthUsersModel
	selectListByUserId func(ctx context.Context, userId int64) ([]model.AuthUsers, error)
}

func (m fakeAuthUsersModel) SelectListByUserId(ctx context.Context, userId int64) ([]model.AuthUsers, error) {
	return m.selectListByUserId(ctx, userId)
}

type fakeAuthsModel struct {
	model.AuthsModel
	findListByAuthKeyId func(ctx context.Context, authKeyId ...int64) ([]model.Auths, error)
}

func (m fakeAuthsModel) FindListByAuthKeyIdList(ctx context.Context, authKeyId ...int64) ([]model.Auths, error) {
	return m.findListByAuthKeyId(ctx, authKeyId...)
}

func TestGetAuthorizationsBatchesAuthRowsAndSkipsMissingAuthData(t *testing.T) {
	repo := &Repository{
		model: &model.Models{
			AuthUsersModel: fakeAuthUsersModel{
				selectListByUserId: func(ctx context.Context, userId int64) ([]model.AuthUsers, error) {
					return []model.AuthUsers{
						{AuthKeyId: 1001, UserId: userId, Hash: 11},
						{AuthKeyId: 1002, UserId: userId, Hash: 22},
					}, nil
				},
			},
			AuthsModel: fakeAuthsModel{
				findListByAuthKeyId: func(ctx context.Context, authKeyId ...int64) ([]model.Auths, error) {
					if len(authKeyId) != 2 || authKeyId[0] != 1001 || authKeyId[1] != 1002 {
						t.Fatalf("FindListByAuthKeyIdList ids = %v, want [1001 1002]", authKeyId)
					}
					return []model.Auths{
						{AuthKeyId: 1002, ClientIp: "127.0.0.1", DeviceModel: "phone"},
					}, nil
				},
			},
		},
	}

	got, err := repo.GetAuthorizations(context.Background(), 777, 1002)
	if err != nil {
		t.Fatalf("GetAuthorizations() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("len(authorizations) = %d, want 1", len(got))
	}
	if !got[0].Current || got[0].DeviceModel != "phone" {
		t.Fatalf("authorization = %#v, want current phone authorization", got[0])
	}
}
