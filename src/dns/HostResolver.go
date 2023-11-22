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
	domains []subDomain
}

type subDomain struct {
	domainName string
	hosts      []Host
	subDomains []subDomain
}

type slList[T comparable] struct {
	head *slListNode[T]
}

type slListNode[T comparable] struct {
	next  *slListNode[T, U]
	value T
}

func (this *slList[T]) contains(v T) bool {
	for curNode := this.head; curNode != nil; curNode = curNode.next {
		if curNode.value == v {
			return true
		}
	}

	return false
}

func (this *slList[T]) add(v T) bool {
	var lastNode *slListNode[T]

	for curNode := this.head; curNode != nil; curNode = curNode.next {
		if curNode.value == v {
			//If we found that we already have then no need to add
			return false
		}
		lastNode = curNode
	}

	//If we are here then we didn't add it
	lastNode.next = &slListNode[T]{value: v}
	return true

}

func (this *slList[T]) find(compFunc func(T, T)) {

}

type subDomainTreeNode struct {
	domain         string
	hostsHead      *hostListNode
	subDomainsHead *subDomainListNode
}

func NewHostResolver() HostResolver {
	return &hostResolverTree{rootNode: newSubDomainTreeNode("_root_")}
}

func newSubDomainTreeNode(domain string) *subDomainTreeNode {
	return &subDomainTreeNode{domain: domain}
}

func (this *hostResolverTree) AddUpdateHost(h Host) error {
	curDomainNode := this.rootNode
	hostDomainNameIter := h.FQDN().Iter()

	for hostDomainNameIter.HasNext() {
		curHostSubDomain := hostDomainNameIter.Next()
		//Lets find the subdomain for this host in our current domain node

	}

	// for curSubDomainNode := curDomainNode.subDomainsHead; curSubDomainNode != nil; curSubDomainNode = curSubDomainNode.next {
	// 	//if we have not found our domain yet
	// 	if curSubDomainNode == nil {
	// 		//Then lets add it
	// 	} else if curSubDomainNode.domainTreeNode.domain ==
	// }

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
	next              *subDomainListNode
	subDomainTreeNode *subDomainTreeNode
}

type hostListNode struct {
	next           *hostListNode
	host           Host
	expirationTime time.Time
}

func (this *hostResolverTree) ResolveDomain(d DomainName) Host {
	return nil
}

// For the passed in subDomainTreeNode
func findSubDomainByName(domain *subDomainTreeNode, subDomainName string) *subDomainTreeNode {
	var subDomain *subDomainTreeNode
	//See if this node has any subdomains
	if domain.subDomainsHead == nil {
		//If not then we get to add it
		subDomain = newSubDomainTreeNode(subDomainName)
		domain.subDomainsHead = &subDomainListNode{subDomainTreeNode: subDomain}
	} else {
		//We have sub domains so we need to check to see if we have it first
		curSubDomain := domain.subDomainsHead
		for {
			//If we found the subdomain that we want then lets return it
			if curSubDomain.subDomainTreeNode.domain == subDomainName {
				return curSubDomain.subDomainTreeNode
			}

			//Now we need to look at the next node

			if curSubDomain.next == nil {
				//However if we don't have one then we need to create it
				subDomain = newSubDomainTreeNode(subDomainName)
				curSubDomain.next = &subDomainListNode{subDomainTreeNode: subDomain}
			} else {
				//We actually have more nodes
				curSubDomain = curSubDomain.next
			}
		}
	}

	return subDomain
}
