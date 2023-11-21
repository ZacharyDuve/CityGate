package dns

import (
	"fmt"
	"strings"
)

type DomainName struct {
	domains []string
}

func ParseDomainNameFromString(dn string) (*DomainName, error) {
	if dn == "" {
		return nil, fmt.Errorf("Unable to parse blank domain name")
	}
	//Need to trim off the "." from the end if we have it
	dn = strings.TrimSuffix(dn, ".")

	domains := strings.Split(dn, ".")

	domainName := &DomainName{domains: domains}

	return domainName, nil
}

func (this DomainName) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(this.domains); i++ {
		sb.WriteString(this.domains[i])

		if i == len(this.domains)-1 {
			sb.WriteRune('.')
		}
	}

	return sb.String()
}

func (this DomainName) Iter() *DomainNameIterator {
	return &DomainNameIterator{domainName: this, curPosition: len(this.domains) - 1}
}

type DomainNameIterator struct {
	domainName  DomainName
	curPosition int
}

func (this *DomainNameIterator) Next() string {
	if !this.HasNext() {
		return ""
	}

	i := this.curPosition
	this.curPosition--
	return this.domainName.domains[i]
}

func (this *DomainNameIterator) HasNext() bool {
	return this.curPosition >= 0
}
