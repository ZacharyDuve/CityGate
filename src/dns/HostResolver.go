package dns

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ZacharyDuve/CityGate/src/datastruct/list"
)

const (
	root_domain_name = "_root_"
)

type HostResolver interface {
	AddUpdateHost(Host) error
	ResolveDomain(DomainName) Host
	PrintTree()
}

type hostResolverTree struct {
	rootDomain *subDomain
}

type subDomain struct {
	domainName  string
	hostRecords list.SingleLinkedList[hostRecord]
	subDomains  list.SingleLinkedList[*subDomain]
	depth       int
}

type hostRecord struct {
	expirationTime time.Time
	host           Host
}

func NewHostResolver() HostResolver {
	return &hostResolverTree{rootDomain: &subDomain{domainName: root_domain_name, depth: 0}}
}

func (this *hostResolverTree) AddUpdateHost(h Host) error {
	curDomain := this.rootDomain
	hostDomainNameIter := h.FQDN().Iter()

	for hostDomainNameIter.HasNext() {
		curHostDomain := hostDomainNameIter.Next()

		matchingSubDomainResult := curDomain.subDomains.Find(func(curSubDomain *subDomain) bool {
			return curSubDomain.domainName == curHostDomain
		})

		if matchingSubDomainResult == nil {
			newDomain := &subDomain{domainName: curHostDomain, depth: curDomain.depth + 1}
			curDomain.subDomains.Add(newDomain)
			curDomain = newDomain
		} else {
			curDomain = matchingSubDomainResult.Value
		}

	}

	//At this point curDomain should be the domain that the host is hosting

	hosts := curDomain.hostRecords
	//Lets see if we have a matching host
	foundHostRecordResult := hosts.Find(func(curHost hostRecord) bool {
		return h.Address().String() == curHost.host.Address().String()
	})

	//Now lets check the results to if actually found our host or if

	if foundHostRecordResult == nil {
		//We don't have a host so lets add it
		if !hosts.Add(newHostRecord(h)) {
			log.Println("Error adding host", h)
		} else {
			log.Println("Host was added")
		}
	} else {
		foundHostRecordResult.Value.expirationTime = time.Now()
	}

	return errors.ErrUnsupported
}

// Function useful for testing
func (this *hostResolverTree) PrintTree() {
	fmt.Print("-\n")
	this.rootDomain.subDomains.ForEach(printSubDomain)

}

func printSubDomain(curSubDomain *subDomain) {

	indent := strings.Repeat(" ", curSubDomain.depth)
	fmt.Print(indent, "- ", curSubDomain.domainName, " ", curSubDomain.hostRecords.Len(), "\n")
	curSubDomain.subDomains.ForEach(printSubDomain)
}

func newHostRecord(h Host) hostRecord {
	return hostRecord{host: h, expirationTime: time.Now()}
}

func (this *hostResolverTree) ResolveDomain(DomainName) Host {
	panic(errors.ErrUnsupported)
	return nil
}
