package repository

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"time"

	gatewayclient "github.com/teamgram/teamgram-server/v2/app/interface/gateway/client"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

const (
	GatewayClientCacheMaxEntriesDefault    = 1000
	GatewayClientCacheTTLSecondsDefault    = 300
	GatewayClientIdleTimeoutSecondsDefault = 60
)

type GatewayClientCache struct {
	mu         sync.Mutex
	maxEntries int
	ttl        time.Duration
	idle       time.Duration
	policy     AddressPolicy
	base       kitex.RpcClientConf
	entries    map[string]*list.Element
	lru        *list.List
	now        func() time.Time
	newPusher  func(conf kitex.RpcClientConf) (GatewayPusher, error)
}

type gatewayClientCacheEntry struct {
	addr         string
	pusher       GatewayPusher
	validatedAt  time.Time
	lastAccessAt time.Time
}

func NewGatewayClientCache(base kitex.RpcClientConf, policy AddressPolicy, maxEntries int, ttl, idle time.Duration) *GatewayClientCache {
	if maxEntries <= 0 {
		maxEntries = GatewayClientCacheMaxEntriesDefault
	}
	if ttl <= 0 {
		ttl = GatewayClientCacheTTLSecondsDefault * time.Second
	}
	if idle <= 0 {
		idle = GatewayClientIdleTimeoutSecondsDefault * time.Second
	}
	return &GatewayClientCache{
		maxEntries: maxEntries,
		ttl:        ttl,
		idle:       idle,
		policy:     policy,
		base:       base,
		entries:    make(map[string]*list.Element),
		lru:        list.New(),
		now:        time.Now,
		newPusher:  newGatewayClientPusher,
	}
}

func (c *GatewayClientCache) Get(ctx context.Context, gatewayRPCAddr string) (GatewayPusher, error) {
	now := c.now()
	c.mu.Lock()
	if elem := c.entries[gatewayRPCAddr]; elem != nil {
		entry := elem.Value.(*gatewayClientCacheEntry)
		if c.idle > 0 && now.Sub(entry.lastAccessAt) > c.idle {
			c.removeElement(elem)
		} else {
			if now.Sub(entry.validatedAt) >= c.ttl {
				c.mu.Unlock()
				if _, err := c.policy.Validate(ctx, gatewayRPCAddr); err != nil {
					return nil, err
				}
				c.mu.Lock()
				if elem = c.entries[gatewayRPCAddr]; elem != nil {
					entry = elem.Value.(*gatewayClientCacheEntry)
					entry.validatedAt = now
					entry.lastAccessAt = now
					c.lru.MoveToFront(elem)
					pusher := entry.pusher
					c.mu.Unlock()
					return pusher, nil
				}
				c.mu.Unlock()
				return c.newClient(ctx, gatewayRPCAddr, now)
			}
			entry.lastAccessAt = now
			c.lru.MoveToFront(elem)
			pusher := entry.pusher
			c.mu.Unlock()
			return pusher, nil
		}
	}
	c.mu.Unlock()
	return c.newClient(ctx, gatewayRPCAddr, now)
}

func (c *GatewayClientCache) PushSessionUpdates(ctx context.Context, route SessionRoute, updates tg.UpdatesClazz) error {
	pusher, err := c.Get(ctx, route.GatewayRPCAddr)
	if err != nil {
		return err
	}
	return pusher.PushSessionUpdates(ctx, route, updates)
}

func (c *GatewayClientCache) PushRpcResult(ctx context.Context, route RpcResultRoute) error {
	pusher, err := c.Get(ctx, route.GatewayRPCAddr)
	if err != nil {
		return err
	}
	return pusher.PushRpcResult(ctx, route)
}

func (c *GatewayClientCache) newClient(ctx context.Context, gatewayRPCAddr string, now time.Time) (GatewayPusher, error) {
	resolved, err := c.policy.Validate(ctx, gatewayRPCAddr)
	if err != nil {
		return nil, err
	}
	conf := c.base
	conf.DestService = firstNonEmpty(conf.DestService, "interface.gateway")
	conf.ServiceName = firstNonEmpty(conf.ServiceName, "RPCGateway")
	conf.Endpoints = []string{resolved.HostPort}
	conf.Etcd.Hosts = nil
	conf.Etcd.Key = ""
	pusher, err := c.newPusher(conf)
	if err != nil {
		return nil, fmt.Errorf("new gateway client %q: %w", resolved.HostPort, err)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem := c.entries[gatewayRPCAddr]; elem != nil {
		entry := elem.Value.(*gatewayClientCacheEntry)
		entry.pusher = pusher
		entry.validatedAt = now
		entry.lastAccessAt = now
		c.lru.MoveToFront(elem)
		return pusher, nil
	}
	elem := c.lru.PushFront(&gatewayClientCacheEntry{
		addr:         gatewayRPCAddr,
		pusher:       pusher,
		validatedAt:  now,
		lastAccessAt: now,
	})
	c.entries[gatewayRPCAddr] = elem
	for len(c.entries) > c.maxEntries {
		c.removeElement(c.lru.Back())
	}
	return pusher, nil
}

func newGatewayClientPusher(conf kitex.RpcClientConf) (GatewayPusher, error) {
	cli, err := gatewayclient.NewKitexClient(conf)
	if err != nil {
		return nil, err
	}
	return gatewayClientPusher{client: gatewayclient.NewGatewayClient(cli)}, nil
}

func (c *GatewayClientCache) removeElement(elem *list.Element) {
	if elem == nil {
		return
	}
	c.lru.Remove(elem)
	entry := elem.Value.(*gatewayClientCacheEntry)
	delete(c.entries, entry.addr)
}

func firstNonEmpty(value, fallback string) string {
	if value != "" {
		return value
	}
	return fallback
}
