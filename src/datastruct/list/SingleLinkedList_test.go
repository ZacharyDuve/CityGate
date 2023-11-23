package list

import "testing"

type testThing struct {
	name   string
	number int64
}

func TestListLifeCycle(t *testing.T) {
	l := &SingleLinkedList[testThing]{}

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

	if l.Find(func(t *testThing) bool { return t.name == myTestThing.name }) != myTestThing {
		t.Fatal("Expected to find my testThing due to name match")
	}

	if !l.Remove(myTestThing) {
		t.Fatal("Expected to remove existing test thing")
	}

	if l.Remove(myTestThing) {
		t.Fatal("Expected second remove to not occur")
	}
}
