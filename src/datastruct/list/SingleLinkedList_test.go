package list

import "testing"

type testThing struct {
	name   string
	number int64
}

func TestListLifeCycle(t *testing.T) {
	l := &SingleLinkedList[*testThing]{}

	myTestThing := &testThing{name: "bib", number: 43}
	MyOtherTestThing := &testThing{name: "chichi", number: 97}

	if l.Contains(myTestThing) {
		t.Fatal("Expected list to not contain test object")
	}

	if !l.Add(myTestThing) {
		t.Fatal("Expected to have Added test thing as it is the first time")
	}

	if l.Add(myTestThing) {
		t.Fatal("Expect second Add to not Add test object")
	}

	if !l.Add(MyOtherTestThing) {
		t.Fatal("Expected to be able to Add a different thing")
	}

	if !l.Contains(myTestThing) {
		t.Fatal("Expect to contains as already exists")
	}

	if res := l.Find(func(t *testThing) bool { return t.name == myTestThing.name }); res != nil && res.Value != myTestThing {
		t.Fatal("Expected to find my testThing due to name match")
	}

	if !l.Remove(myTestThing) {
		t.Fatal("Expected to remove existing test thing")
	}

	if l.Remove(myTestThing) {
		t.Fatal("Expected second remove to not occur")
	}
}

//======================================== Len() ==================================================

func TestThatLenStartsAt0(t *testing.T) {
	l := &SingleLinkedList[int]{}

	if l.Len() != 0 {
		t.Fatalf("Expected Len of new list to be 0 but it was %d", l.Len())
	}
}

func TestThatLenIsNAfterAddingNElements(t *testing.T) {
	l := &SingleLinkedList[int]{}

	numElemsToAdd := 10

	for i := 0; i < numElemsToAdd; i++ {
		l.Add(i * i)
	}

	if l.Len() != numElemsToAdd {
		t.Fatalf("Expected Len of %d after adding that many elements but instead got %d", numElemsToAdd, l.Len())
	}
}

func TestThatLenIsOriginalMinusNAfterRemovingNElements(t *testing.T) {
	l := &SingleLinkedList[int]{}

	numElemsToAdd := 10

	for i := 0; i < numElemsToAdd; i++ {
		l.Add(i)
	}

	numElemsToRemove := 5

	if numElemsToAdd < numElemsToRemove {
		t.Fatal("BAD TEST as test is trying to remove more elements than it added")
	}

	for i := 0; i < numElemsToRemove; i++ {
		l.Remove(i)
	}

	if l.Len() != numElemsToAdd-numElemsToRemove {
		t.Fatalf("Expected Len of %d after removing that many elements but instead got %d", numElemsToAdd-numElemsToRemove, l.Len())
	}
}
