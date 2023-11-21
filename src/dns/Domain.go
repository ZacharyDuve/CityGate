package dns

import (
	"net"
	"sync"
)

type domain struct {
	name       DomainName
	subDomains []*domain
}

func newDomain(n DomainName) *domain {
	return &domain{name: n, subDomains: make([]*domain, 0)}
}

type hostedDomainsManager struct {
	domains []*domain
	hosts   map[net.Addr]Host
	rwLock  *sync.RWMutex
}

func NewHostedDomainsManager() *hostedDomainsManager {
	dM := &hostedDomainsManager{domains: make([]*domain, 0), hosts: make(map[net.Addr]Host)}

	return dM
}
