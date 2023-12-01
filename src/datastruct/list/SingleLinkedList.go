package list

/*
Attempt at making a generic single linked list
*/

type SingleLinkedList[T comparable] struct {
	head *slListNode[T]
}

type slListNode[T comparable] struct {
	next  *slListNode[T]
	value T
}

func (this *SingleLinkedList[T]) Contains(v T) bool {
	for curNode := this.head; curNode != nil; curNode = curNode.next {
		if curNode.value == v {
			return true
		}
	}

	return false
}

func (this *SingleLinkedList[T]) Add(v T) bool {
	var lastNode *slListNode[T]

	if this.head == nil {
		this.head = &slListNode[T]{value: v}
		return true
	}

	for curNode := this.head; curNode != nil; curNode = curNode.next {
		if curNode.value == v {
			//If we found that we already have then no need to add
			return false
		}
		lastNode = curNode
	}

	//If we are here then we didn't add it
	lastNode.next = &slListNode[T]{value: v}
	return true

}

type Results[T comparable] struct {
	Value T
}

func (this *SingleLinkedList[T]) Find(compFunc func(T) bool) *Results[T] {
	for curNode := this.head; curNode != nil; curNode = curNode.next {
		if compFunc(curNode.value) {
			//If we found that we already have then no need to add
			return &Results[T]{Value: curNode.value}
		}
	}

	return nil
}

func (this *SingleLinkedList[T]) Remove(t T) bool {
	var prevNode *slListNode[T]
	for curNode := this.head; curNode != nil; curNode = curNode.next {
		if curNode.value == t {
			//If we found that we already have then no need to add
			if prevNode == nil {
				this.head = curNode.next
			} else {
				prevNode.next = curNode.next
			}
			return true
		}

		prevNode = curNode
	}

	return false
}
