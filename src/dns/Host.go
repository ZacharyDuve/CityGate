package dns

import (
	"net"
)

type Host interface {
	FQDN() DomainName
	Address() net.Addr
}

type host struct {
	fqdn string
	addr net.Addr
}
