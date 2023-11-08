package dns

import (
	"fmt"
	"strings"
)

type DomainName []string

func ParseDomainNameFromString(dn string) (DomainName, error) {
	if dn == "" {
		return nil, fmt.Errorf("Unable to parse blank domain name")
	}
	//Need to trim off the "." from the end if we have it
	dn = strings.TrimSuffix(dn, ".")

	domains := strings.Split(dn, ".")

	return domains, nil
}

func DomainNameToString(dn DomainName) string {
	if len(dn) == 0 {
		return "."
	}

	retDN := ""

	for _, curDomain := range dn {
		retDN += curDomain + "."
	}

	return retDN
}
