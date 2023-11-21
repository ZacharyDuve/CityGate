package dns

import (
	"errors"
	"time"
)

type HostResolver interface {
	AddUpdateHost(Host) error
	ResolveDomain(DomainName) Host
}

type hostResolverTree struct {
	rootNode *resolverTreeNode
}

type resolverTreeNode struct {
	domain         string
	hostsHead      *hostListNode
	subDomainsHead *subDomainListNode
}

func NewHostResolver() HostResolver {
	return &hostResolverTree{rootNode: newResolverTreeNode("_root_")}
}

func newResolverTreeNode(domain string) *resolverTreeNode {
	return &resolverTreeNode{domain: domain}
}

func (this *hostResolverTree) AddUpdateHost(h Host) error {
	curDomainNode := this.rootNode

	for curSubDomainNode := curDomainNode.subDomainsHead; curSubDomainNode != nil; curSubDomainNode = curSubDomainNode.next {
		if curSubDomainNode
	}

	//OLD CODE
	// hostDNIter := h.FQDN().Iter()
	// for hostDNIter.HasNext() {
	// 	curHostDomain := hostDNIter.Next()
	// 	for _, curHostedDomain := range curHostedDomains {
	// 		//Look to see if we have a match of the hosts domain to the hosted domain
	// 		if curHostedDomain.domain == curHostDomain {
	// 			//If so
	// 		}
	// 	}
	// }

	return errors.ErrUnsupported
}

type subDomainListNode struct {
	next           *subDomainListNode
	domainTreeNode *resolverTreeNode
}

type hostListNode struct {
	next           *hostListNode
	host           Host
	expirationTime time.Time
}

func (this *hostResolverTree) ResolveDomain(d DomainName) Host {
	return nil
}
