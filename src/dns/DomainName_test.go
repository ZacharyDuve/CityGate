package dns

import "testing"

func TestThatParseDomainNameFromBlankStringReturnsError(t *testing.T) {
	_, err := ParseDomainNameFromString("")

	if err == nil {
		t.Fail()
	}
}

func TestParseDomainNameOfOnePartNoEndDotReturnsNoError(t *testing.T) {
	_, err := ParseDomainNameFromString("com")

	if err != nil {
		t.Fail()
	}
}

func TestParseDomainNameOfOnePartWithEndDotReturnsNoError(t *testing.T) {
	_, err := ParseDomainNameFromString("com.")

	if err != nil {
		t.Fail()
	}
}

func TestParseDomainNameOfOnePartNoEndDotReturnsDomain(t *testing.T) {
	dn, _ := ParseDomainNameFromString("com")

	if dn.Length() == 0 || dn.Get(0) != "com" {
		t.Fail()
	}
}

func TestParseDomainNameOfOnePartWithEndDotReturnsDomain(t *testing.T) {
	dn, _ := ParseDomainNameFromString("com.")

	if dn.Length() == 0 || dn.Get(0) != "com" {
		t.Fail()
	}
}
