package dns

import (
	"errors"
	"log"
	"time"

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
	domainName  string
	hostRecords list.SingleLinkedList[hostRecord]
	subDomains  list.SingleLinkedList[*subDomain]
}

type hostRecord struct {
	expirationTime time.Time
	host           Host
}

func NewHostResolver() HostResolver {
	return &hostResolverTree{rootDomain: &subDomain{domainName: root_domain_name}}
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
			newDomain := &subDomain{domainName: curHostDomain}
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
		}
	} else {
		foundHostRecordResult.Value.expirationTime = time.Now()
	}

	return errors.ErrUnsupported
}

func newHostRecord(h Host) hostRecord {
	return hostRecord{host: h, expirationTime: time.Now()}
}

func (this *hostResolverTree) ResolveDomain(DomainName) Host {
	panic(errors.ErrUnsupported)
	return nil
}
