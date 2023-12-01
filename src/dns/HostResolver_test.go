package dns

import (
	"net"
	"testing"
)

func TestAddSingleHostWorksAsExpected(t *testing.T) {
	resovler := NewHostResolver()
	h := createTestHost("www.somewhere.go", "10.3.4.5")
	resovler.AddUpdateHost(h)
}

func TestAddSameHostTwiceDoesNotCrashAndBurn(t *testing.T) {
	resovler := NewHostResolver()
	h := createTestHost("www.somewhere.go", "10.3.4.5")
	resovler.AddUpdateHost(h)
	resovler.AddUpdateHost(h)
}

func TestAddTwoHostsForSameDomainDoesNotCrashAndBurn(t *testing.T) {

	resovler := NewHostResolver()
	h := createTestHost("www.somewhere.go", "10.3.4.6")
	resovler.AddUpdateHost(h)
	h2 := createTestHost("www.somewhere.go", "10.3.4.5")
	resovler.AddUpdateHost(h2)
}

func TestAddTwoHostsForDifferentDomainDoesNotCrashAndBurn(t *testing.T) {

	resovler := NewHostResolver()
	h := createTestHost("www.somewhere.else.com", "10.3.4.6")
	resovler.AddUpdateHost(h)
	h2 := createTestHost("www.somewhere.go", "10.3.4.5")
	resovler.AddUpdateHost(h2)
}

func createTestHost(dn string, addr string) Host {
	d, _ := ParseDomainNameFromString(dn)
	a := &net.IPAddr{IP: net.ParseIP(addr)}
	return &testHost{FullyDomainName: d, Addr: a}
}
