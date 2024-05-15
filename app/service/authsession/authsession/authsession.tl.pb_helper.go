// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package authsession

import (
	"strings"
)

func (m *AuthKeyStateData) Layer() int32 {
	return m.GetClient().GetLayer()
}

func (m *AuthKeyStateData) ApiId() int32 {
	return m.GetClient().GetApiId()
}

func (m *AuthKeyStateData) DeviceModel() string {
	return m.GetClient().GetDeviceModel()
}

func (m *AuthKeyStateData) SystemVersion() string {
	return m.GetClient().GetSystemVersion()
}

func (m *AuthKeyStateData) AppVersion() string {
	return m.GetClient().GetAppVersion()
}

func (m *AuthKeyStateData) SystemLangCode() string {
	return m.GetClient().GetSystemLangCode()
}

func (m *AuthKeyStateData) LangPack() string {
	c := m.GetClient().GetLangPack()

	if c == "" {
		if strings.HasSuffix(m.GetClient().GetAppVersion(), " A") {
			c = "weba"
		} else if strings.HasSuffix(m.GetClient().GetAppVersion(), " Z") {
			c = "weba"
		}
	}

	return c
}

func (m *AuthKeyStateData) LangCode() string {
	return m.GetClient().GetLangCode()
}

func (m *AuthKeyStateData) ClientIp() string {
	return m.GetClient().GetIp()
}

func (m *AuthKeyStateData) Proxy() string {
	return m.GetClient().GetProxy()
}

func (m *AuthKeyStateData) Params() string {
	return m.GetClient().GetParams()
}

func (m *AuthKeyStateData) ClientName() string {
	c := m.GetClient().GetLangPack()

	if c == "android" {
		if strings.Index(m.GetClient().GetAppVersion(), "TDLib") >= 0 {
			c = "react"
		}
	} else if c == "" {
		if strings.HasSuffix(m.GetClient().GetAppVersion(), " A") {
			c = "weba"
		} else if strings.HasSuffix(m.GetClient().GetAppVersion(), " Z") {
			c = "weba"
		}
	}

	return c
}
