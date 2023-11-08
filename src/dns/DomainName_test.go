package dns

import "testing"

func TestThatParseDomainNameFromBlankStringReturnsError(t *testing.T) {
	_, err := ParseDomainNameFromString("")

	if err == nil {
		t.Fail()
	}
}
