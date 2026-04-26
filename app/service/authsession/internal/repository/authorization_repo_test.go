package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/infra/geoip/geoip"
	"github.com/teamgram/teamgram-server/v2/app/service/authsession/authsession"
)

type fakeGeoipClient struct {
	region *geoip.Region
	err    error
	closed *bool
}

func (f fakeGeoipClient) GeoipGetCountryAndRegionByIp(context.Context, *geoip.TLGeoipGetCountryAndRegionByIp) (*geoip.Region, error) {
	return f.region, f.err
}

func (f fakeGeoipClient) Close() error {
	if f.closed != nil {
		*f.closed = true
	}
	return nil
}

func TestAuthorizationCurrentHashIsZero(t *testing.T) {
	auth := authDataToAuthorization(&cacheAuthData{
		Client:   &authsession.ClientSession{Ip: "127.0.0.1"},
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
		Client:   &authsession.ClientSession{Ip: "127.0.0.1"},
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

func TestRepositoryCloseClosesGeoipClient(t *testing.T) {
	closed := false
	r := &Repository{
		geoipClient: fakeGeoipClient{closed: &closed},
	}

	if err := r.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if !closed {
		t.Fatal("Close() did not close geoip client")
	}
}
