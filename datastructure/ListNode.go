package datastructure

/**
 * 链表
 */
type ListNode[T comparable] struct {
	Val  T
	Next *ListNode[T]
}

func (ln *ListNode[T]) ArrayToListNode(arr []T) *ListNode[T] {

	if len(arr) == 0 {
		return nil
	}

	head := &ListNode[T]{Val: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode[T]{Val: arr[i]}
		current = current.Next
	}

	return head
}

func (ln *ListNode[T]) ListNodeToArray() []T {
	res := []T{}

	for ln != nil {
		res = append(res, ln.Val)
		ln = ln.Next
	}

	return res
}
