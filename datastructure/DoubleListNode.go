package datastructure

/**
 *	双链表
 */
type DoubleListNode[T comparable] struct {
	Key   T
	Value T
	Prev  *DoubleListNode[T]
	Next  *DoubleListNode[T]
}
