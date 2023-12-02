package main

import (
	"net"

	"github.com/ZacharyDuve/CityGate/src/dns"
)

type testHost struct {
	FullyDomainName *dns.DomainName
	Addr            net.Addr
}

func (this *testHost) FQDN() *dns.DomainName {
	return this.FullyDomainName
}

func (this *testHost) Address() net.Addr {
	return this.Addr
}

func main() {
	resovler := dns.NewHostResolver()
	h := createTestHost("www.somewhere.else.com", "10.3.4.6")
	resovler.AddUpdateHost(h)
	h2 := createTestHost("www.somewhere.com", "10.3.4.5")
	resovler.AddUpdateHost(h2)

	resovler.PrintTree()
}

func createTestHost(dn string, addr string) dns.Host {
	d, _ := dns.ParseDomainNameFromString(dn)
	a := &net.IPAddr{IP: net.ParseIP(addr)}
	return &testHost{FullyDomainName: d, Addr: a}
}
