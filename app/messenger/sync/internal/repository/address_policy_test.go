package repository

import (
	"context"
	"errors"
	"net"
	"testing"

	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
)

func TestAddressPolicyRejectsSchemeAndMetadataIP(t *testing.T) {
	p := AddressPolicy{
		AllowedCIDRs:  []string{"127.0.0.0/8"},
		AllowedPorts:  []int{20110},
		AllowLoopback: true,
	}
	for _, addr := range []string{"http://127.0.0.1:20110", "169.254.169.254:80", "0.0.0.0:20110"} {
		if _, err := p.Validate(context.Background(), addr); !errors.Is(err, syncpb.ErrSyncAddressRejected) {
			t.Fatalf("Validate(%q) error = %v, want ErrSyncAddressRejected", addr, err)
		}
	}
}

func TestAddressPolicyAcceptsAllowedHostPort(t *testing.T) {
	p := AddressPolicy{
		AllowedCIDRs:  []string{"127.0.0.0/8"},
		AllowedPorts:  []int{20110},
		AllowLoopback: true,
	}
	resolved, err := p.Validate(context.Background(), "127.0.0.1:20110")
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if resolved.HostPort != "127.0.0.1:20110" {
		t.Fatalf("HostPort = %q, want 127.0.0.1:20110", resolved.HostPort)
	}
}

func TestAddressPolicyPinsHostnameToValidatedIP(t *testing.T) {
	p := AddressPolicy{
		AllowedCIDRs: []string{"10.0.0.0/8"},
		AllowedPorts: []int{20110},
		Resolver: func(ctx context.Context, host string) ([]net.IP, error) {
			if host != "gateway.internal" {
				t.Fatalf("resolver host = %q, want gateway.internal", host)
			}
			return []net.IP{net.ParseIP("10.20.30.40")}, nil
		},
	}
	resolved, err := p.Validate(context.Background(), "gateway.internal:20110")
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if resolved.HostPort != "10.20.30.40:20110" {
		t.Fatalf("HostPort = %q, want pinned validated IP", resolved.HostPort)
	}
}

func TestAddressPolicyRejectsIPv6UnlessConfigured(t *testing.T) {
	p := AddressPolicy{
		AllowedCIDRs:  []string{"127.0.0.0/8"},
		AllowedPorts:  []int{20110},
		AllowLoopback: true,
	}
	_, err := p.Validate(context.Background(), "[::1]:20110")
	if !errors.Is(err, syncpb.ErrSyncAddressRejected) {
		t.Fatalf("Validate() error = %v, want ErrSyncAddressRejected", err)
	}
}

func TestAddressPolicyAcceptsIPv6LoopbackWhenAllowedCIDRConfigured(t *testing.T) {
	p := AddressPolicy{
		AllowedIPv6CIDRs: []string{"::1/128"},
		AllowedPorts:     []int{20110},
		AllowLoopback:    true,
	}
	resolved, err := p.Validate(context.Background(), "[::1]:20110")
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	if resolved.HostPort != "[::1]:20110" {
		t.Fatalf("HostPort = %q, want [::1]:20110", resolved.HostPort)
	}
}
