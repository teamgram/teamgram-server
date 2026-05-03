// Copyright (c) 2026 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package core

import (
	"context"
	"testing"

	"github.com/teamgram/teamgram-server/v2/app/bff/configuration/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/bff/configuration/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func newStartupConfigurationCore() *ConfigurationCore {
	return New(context.Background(), svc.NewServiceContext(config.Config{}))
}

func TestHelpGetConfigReturnsValidStartupConfig(t *testing.T) {
	c := newStartupConfigurationCore()

	got, err := c.HelpGetConfig(&tg.TLHelpGetConfig{})
	if err != nil {
		t.Fatalf("HelpGetConfig error = %v", err)
	}
	if got == nil {
		t.Fatal("HelpGetConfig returned nil")
	}
	if got.ThisDc != startupThisDC {
		t.Fatalf("ThisDc = %d, want %d", got.ThisDc, startupThisDC)
	}
	if len(got.DcOptions) == 0 {
		t.Fatal("DcOptions is empty")
	}
	dc := got.DcOptions[0]
	if dc.IpAddress != "127.0.0.1" || dc.Port != 10443 {
		t.Fatalf("DcOptions[0] = %s:%d, want 127.0.0.1:10443", dc.IpAddress, dc.Port)
	}
	if got.Date <= 0 || got.Expires <= got.Date {
		t.Fatalf("Date/Expires = %d/%d, want positive date and later expiry", got.Date, got.Expires)
	}
	if got.TestMode == nil {
		t.Fatal("TestMode is nil")
	}
	if got.DcTxtDomainName == "" {
		t.Fatal("DcTxtDomainName is empty")
	}
	if got.MeUrlPrefix == "" {
		t.Fatal("MeUrlPrefix is empty")
	}
	if err := got.Validate(223); err != nil {
		t.Fatalf("Config Validate(223) error = %v", err)
	}
}

func TestHelpGetNearestDcReturnsCurrentDc(t *testing.T) {
	c := newStartupConfigurationCore()

	got, err := c.HelpGetNearestDc(&tg.TLHelpGetNearestDc{})
	if err != nil {
		t.Fatalf("HelpGetNearestDc error = %v", err)
	}
	if got == nil {
		t.Fatal("HelpGetNearestDc returned nil")
	}
	if got.Country != startupCountry || got.ThisDc != startupThisDC || got.NearestDc != startupThisDC {
		t.Fatalf("nearest dc = %+v, want country=%q this_dc=%d nearest_dc=%d", got, startupCountry, startupThisDC, startupThisDC)
	}
	if err := got.Validate(223); err != nil {
		t.Fatalf("NearestDc Validate(223) error = %v", err)
	}
}

func TestHelpGetAppConfigReturnsConfigThenNotModified(t *testing.T) {
	c := newStartupConfigurationCore()

	got, err := c.HelpGetAppConfig(&tg.TLHelpGetAppConfig{Hash: 0})
	if err != nil {
		t.Fatalf("HelpGetAppConfig(hash=0) error = %v", err)
	}
	appConfig, ok := got.ToHelpAppConfig()
	if !ok {
		t.Fatalf("HelpGetAppConfig(hash=0) returned %s, want help.appConfig", got.ClazzName())
	}
	if appConfig.Hash != startupAppConfigHash {
		t.Fatalf("app config hash = %d, want %d", appConfig.Hash, startupAppConfigHash)
	}
	if appConfig.Config == nil {
		t.Fatal("app config JSONValue is nil")
	}

	got, err = c.HelpGetAppConfig(&tg.TLHelpGetAppConfig{Hash: startupAppConfigHash})
	if err != nil {
		t.Fatalf("HelpGetAppConfig(hash=startupAppConfigHash) error = %v", err)
	}
	if _, ok := got.ToHelpAppConfigNotModified(); !ok {
		t.Fatalf("HelpGetAppConfig(hash=startupAppConfigHash) returned %s, want help.appConfigNotModified", got.ClazzName())
	}
}

func TestHelpGetCountriesListReturnsEmptyListThenNotModified(t *testing.T) {
	c := newStartupConfigurationCore()

	got, err := c.HelpGetCountriesList(&tg.TLHelpGetCountriesList{LangCode: "en", Hash: 0})
	if err != nil {
		t.Fatalf("HelpGetCountriesList(hash=0) error = %v", err)
	}
	list, ok := got.ToHelpCountriesList()
	if !ok {
		t.Fatalf("HelpGetCountriesList(hash=0) returned %s, want help.countriesList", got.ClazzName())
	}
	if list.Hash != startupCountriesHash {
		t.Fatalf("countries hash = %d, want %d", list.Hash, startupCountriesHash)
	}
	if len(list.Countries) != 0 {
		t.Fatalf("countries len = %d, want 0", len(list.Countries))
	}

	got, err = c.HelpGetCountriesList(&tg.TLHelpGetCountriesList{LangCode: "en", Hash: startupCountriesHash})
	if err != nil {
		t.Fatalf("HelpGetCountriesList(hash=startupCountriesHash) error = %v", err)
	}
	if _, ok := got.ToHelpCountriesListNotModified(); !ok {
		t.Fatalf("HelpGetCountriesList(hash=startupCountriesHash) returned %s, want help.countriesListNotModified", got.ClazzName())
	}
}

func TestHelpGetSupportReturnsStaticSupportUser(t *testing.T) {
	c := newStartupConfigurationCore()

	got, err := c.HelpGetSupport(&tg.TLHelpGetSupport{})
	if err != nil {
		t.Fatalf("HelpGetSupport error = %v", err)
	}
	if got == nil {
		t.Fatal("HelpGetSupport returned nil")
	}
	if got.PhoneNumber != supportPhoneNumber {
		t.Fatalf("PhoneNumber = %q, want %q", got.PhoneNumber, supportPhoneNumber)
	}
	user, ok := got.User.(*tg.TLUser)
	if !ok {
		t.Fatalf("User = %T, want *tg.TLUser", got.User)
	}
	if user.Id != supportUserID {
		t.Fatalf("support user id = %d, want %d", user.Id, supportUserID)
	}
	if user.FirstName == nil || *user.FirstName != supportName {
		t.Fatalf("support user first_name = %v, want %q", user.FirstName, supportName)
	}
	if user.AccessHash == nil || *user.AccessHash != supportUserAccessHash {
		t.Fatalf("support user access_hash = %v, want %d", user.AccessHash, supportUserAccessHash)
	}
	if !user.Support {
		t.Fatal("support user Support flag is false")
	}
	if err := got.Validate(223); err != nil {
		t.Fatalf("HelpSupport Validate(223) error = %v", err)
	}
}

func TestHelpGetSupportNameReturnsStaticName(t *testing.T) {
	c := newStartupConfigurationCore()

	got, err := c.HelpGetSupportName(&tg.TLHelpGetSupportName{})
	if err != nil {
		t.Fatalf("HelpGetSupportName error = %v", err)
	}
	if got == nil {
		t.Fatal("HelpGetSupportName returned nil")
	}
	if got.Name != supportName {
		t.Fatalf("Name = %q, want %q", got.Name, supportName)
	}
	if err := got.Validate(223); err != nil {
		t.Fatalf("HelpSupportName Validate(223) error = %v", err)
	}
}
