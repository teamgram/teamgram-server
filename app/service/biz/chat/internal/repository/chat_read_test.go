package repository

import "testing"

func TestGetMutableChatRequiresIntegrationDB(t *testing.T) {
	t.Skip("repository aggregate read requires a MySQL fixture")
}

func TestGetChatListRequiresIntegrationDB(t *testing.T) {
	t.Skip("repository aggregate list read requires a MySQL fixture")
}

func TestSearchRequiresIntegrationDB(t *testing.T) {
	t.Skip("repository search read requires a MySQL fixture")
}
