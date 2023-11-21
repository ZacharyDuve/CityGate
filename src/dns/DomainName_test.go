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

	iter := dn.Iter()

	if !iter.HasNext() {
		t.Fatal("Expected domain name iterator to have next value but didn't")
	}

	if iter.Next() != "com" {
		t.Fatal("Expected iterator for domain name to return com")
	}
}

func TestParseDomainNameMultiPartIteratesOverAllParts(t *testing.T) {
	dn, _ := ParseDomainNameFromString("something.www.z.com")

	iter := dn.Iter()

	count := 0
	for {
		if !iter.HasNext() {
			break
		}
		curName := iter.Next()

		switch {
		case count == 0 && curName != "com":
			t.Fatalf("For first index expected %s but got %s instead", "com", curName)
		case count == 1 && curName != "z":
			t.Fatalf("For second index expected %s but got %s instead", "z", curName)
		case count == 2 && curName != "www":
			t.Fatalf("For third index expected %s but got %s instead", "www", curName)
		case count == 3 && curName != "something":
			t.Fatalf("For fourth index expected %s but got %s instead", "something", curName)
		}
		count++
	}

	if count != 4 {
		t.Fatalf("Failed to process all domains, processed %d instead of %d", count, 4)
	}
}
