package linkedlist

// A linked list contains an address, a record and a pointer to the next node's address
type linkedListNode struct {
	record *int32
	next   *linkedListNode
}

type linkedList struct {
	head *linkedListNode
}

// lets start with inserting a node at the end
func (ll *linkedList) insert(record *int32) {
	if ll.head == nil {
		ll.head = &linkedListNode{record: record, next: nil}
		return
	}

	ptr := ll.head

	for ptr.next != nil {
		ptr = ptr.next
	}

	ptr.next = &linkedListNode{record: record, next: nil}
}

// deleting the first occurance of a node with a certain value
func (ll *linkedList) deleteValue(record *int32) (status int) {
	ptr := ll.head

	if ptr == nil {
		return 0
	}

	if ptr.record == record {
		ll.head = ptr.next
		return 1
	}

	for ptr.next != nil {
		if ptr.next.record == record {
			ptr.next = ptr.next.next
			return 1
		}
		ptr = ptr.next
	}

	return 0
}

// Helper function to create a linked list from an array of integers
func createLinkedListFromArray(arr []int32) *linkedList {
	ll := &linkedList{}
	for i := range arr {
		// Issue lied here, because all the record values were actually pointing to the same
		// address, i.e, the address of the variable val, since it was being reused here
		// ll.insert(&val)
		// Changed to the following to avoid that issue.
		ll.insert(&arr[i])
	}
	return ll
}

// Helper function to get the linked list as an array of integers
func getLinkedListAsArray(ll *linkedList) []int32 {
	var arr []int32
	ptr := ll.head
	for ptr != nil {
		arr = append(arr, *ptr.record)
		ptr = ptr.next
	}
	return arr
}
