package dns

import (
	"net"
	"sync"
	"time"
)

type DomainName []string

type domain struct {
	name       DomainName
	subDomains []*domain
	hosts      []*host
}

func newDomain(n DomainName) *domain {
	return &domain{name: n, subDomains: make([]*domain, 0), hosts: make([]*host, 0)}
}

type Host interface {
	FQDN() DomainName
	Address() net.Addr
	Expires() time.Time
}

type host struct {
	fqdn           string
	addr           net.Addr
	expirationTime time.Time
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

func (this *hostedDomainsManager) AddHost(h Host) (added bool) {
	this.rwLock.RLock()
	_, contains := this.hosts[h.Address()]
	this.rwLock.RUnlock()

	if contains {
		return false
	}
	//we do not contain the host yet
	this.rwLock.Lock()
	this.hosts[h.Address()] = h

	this.rwLock.Unlock()
}
