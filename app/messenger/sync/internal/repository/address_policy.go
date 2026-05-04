package repository

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/teamgram/teamgram-server/v2/app/messenger/sync/internal/metrics"
	syncpb "github.com/teamgram/teamgram-server/v2/app/messenger/sync/sync"
)

type AddressPolicy struct {
	AllowedCIDRs     []string
	AllowedIPv6CIDRs []string
	AllowedPorts     []int
	AllowLoopback    bool
	Resolver         func(ctx context.Context, host string) ([]net.IP, error)
}

type ResolvedGatewayAddr struct {
	HostPort string
	IPs      []net.IP
}

func (p AddressPolicy) Validate(ctx context.Context, addr string) (ResolvedGatewayAddr, error) {
	if addr == "" {
		return ResolvedGatewayAddr{}, addressRejected("empty address")
	}
	if strings.Contains(addr, "://") || strings.ContainsAny(addr, "/?@") {
		return ResolvedGatewayAddr{}, addressRejected("invalid address syntax")
	}
	host, portText, err := net.SplitHostPort(addr)
	if err != nil {
		return ResolvedGatewayAddr{}, addressRejected("invalid hostport: %v", err)
	}
	port, err := strconv.Atoi(portText)
	if err != nil {
		return ResolvedGatewayAddr{}, addressRejected("invalid port")
	}
	if !p.portAllowed(port) {
		return ResolvedGatewayAddr{}, addressRejected("port %d rejected", port)
	}
	ips, err := p.resolve(ctx, host)
	if err != nil {
		metrics.AddressRejected("rejected")
		return ResolvedGatewayAddr{}, fmt.Errorf("%w: resolve %q: %w", syncpb.ErrSyncAddressRejected, host, err)
	}
	if len(ips) == 0 {
		return ResolvedGatewayAddr{}, addressRejected("host %q resolved no ips", host)
	}
	for _, ip := range ips {
		if err := p.validateIP(ip); err != nil {
			return ResolvedGatewayAddr{}, err
		}
	}
	return ResolvedGatewayAddr{HostPort: net.JoinHostPort(ips[0].String(), portText), IPs: ips}, nil
}

func (p AddressPolicy) portAllowed(port int) bool {
	for _, allowed := range p.AllowedPorts {
		if allowed == port {
			return true
		}
	}
	return false
}

func (p AddressPolicy) resolve(ctx context.Context, host string) ([]net.IP, error) {
	if ip := net.ParseIP(host); ip != nil {
		return []net.IP{ip}, nil
	}
	resolver := p.Resolver
	if resolver != nil {
		return resolver(ctx, host)
	}
	addrs, err := net.DefaultResolver.LookupIPAddr(ctx, host)
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0, len(addrs))
	for _, addr := range addrs {
		ips = append(ips, addr.IP)
	}
	return ips, nil
}

func (p AddressPolicy) validateIP(ip net.IP) error {
	if ip == nil {
		return addressRejected("nil ip")
	}
	if ip.IsUnspecified() || ip.IsMulticast() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() || ip.Equal(net.ParseIP("169.254.169.254")) {
		return addressRejected("ip %s rejected", ip.String())
	}
	if ip.To4() == nil {
		if !cidrAllowed(ip, p.AllowedIPv6CIDRs) {
			return addressRejected("ipv6 ip %s rejected", ip.String())
		}
		if ip.IsLoopback() && !p.AllowLoopback {
			return addressRejected("loopback ip %s rejected", ip.String())
		}
		return nil
	}
	if ip.IsLoopback() {
		if p.AllowLoopback {
			return nil
		}
		return addressRejected("loopback ip %s rejected", ip.String())
	}
	if cidrAllowed(ip, p.AllowedCIDRs) {
		return nil
	}
	return addressRejected("ipv4 ip %s rejected", ip.String())
}

func cidrAllowed(ip net.IP, cidrs []string) bool {
	for _, raw := range cidrs {
		_, network, err := net.ParseCIDR(raw)
		if err == nil && network.Contains(ip) {
			return true
		}
	}
	return false
}

func addressRejected(format string, args ...interface{}) error {
	metrics.AddressRejected("rejected")
	return fmt.Errorf("%w: "+format, append([]interface{}{syncpb.ErrSyncAddressRejected}, args...)...)
}
