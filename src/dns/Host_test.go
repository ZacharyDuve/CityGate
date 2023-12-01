package dns

import "net"

type testHost struct {
	FullyDomainName *DomainName
	Addr            net.Addr
}

func (this *testHost) FQDN() *DomainName {
	return this.FullyDomainName
}

func (this *testHost) Address() net.Addr {
	return this.Addr
}
