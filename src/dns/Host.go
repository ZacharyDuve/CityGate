package dns

import (
	"net"
)

type Host interface {
	FQDN() *DomainName
	Address() net.Addr
}

type host struct {
	fqdn DomainName
	addr net.Addr
}
