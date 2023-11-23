package dns

import (
	"errors"

	"github.com/ZacharyDuve/CityGate/src/datastruct/list"
)

const (
	root_domain_name = "_root_"
)

type HostResolver interface {
	AddUpdateHost(Host) error
	ResolveDomain(DomainName) Host
}

type hostResolverTree struct {
	rootDomain *subDomain
}

type subDomain struct {
	domainName string
	hosts      list.SingleLinkedList[Host]
	subDomains list.SingleLinkedList[subDomain]
}

// type subDomainTreeNode struct {
// 	domain         string
// 	hostsHead      *hostListNode
// 	subDomainsHead *subDomainListNode
// }

func NewHostResolver() HostResolver {
	return &hostResolverTree{rootDomain: &subDomain{domainName: root_domain_name}}
}

func (this *hostResolverTree) AddUpdateHost(h Host) error {
	curDomain := this.rootDomain
	hostDomainNameIter := h.FQDN().Iter()

	for hostDomainNameIter.HasNext() {
		curHostDomain := hostDomainNameIter.Next()

		matchingSubDomain := curDomain.subDomains.Find(func(curSubDomain *subDomain) bool {
			return curSubDomain.domainName == curHostDomain
		})

		if matchingSubDomain == nil {
			matchingSubDomain = &subDomain{domainName: curHostDomain}
			curDomain.subDomains.Add(matchingSubDomain)
		}

		curDomain = matchingSubDomain
	}

	//At this point curDomain should be the domain that the host is hosting

	return errors.ErrUnsupported
}

func (this *hostResolverTree) ResolveDomain(DomainName) Host {
	panic(errors.ErrUnsupported)
	return nil
}
