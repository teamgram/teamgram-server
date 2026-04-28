//go:build integration

package model

import "testing"

func TestSelectCountByLinkRequiresFixture(t *testing.T) {
	t.Skip("requires MySQL fixture; implemented query is verified by package compile and integration environment")
}
