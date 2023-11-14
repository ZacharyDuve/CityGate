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

func (this *DomainName) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(this.domains); i++ {
		sb.WriteString(this.domains[i])

		if i == len(this.domains)-1 {
			sb.WriteRune('.')
		}
	}

	return sb.String()
}

func (this *DomainName) Iter() func() (domain string, done bool) {
	curIndex := len(this.domains)

	return func() (domain string, done bool) {
		curIndex--
		if curIndex < 0 {
			return "", true
		}
		return this.domains[curIndex], false
	}
}

func (this *DomainName) Length() int {
	return len(this.domains)
}
