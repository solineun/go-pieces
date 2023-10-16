package main

type LinkedList struct {
	Value interface{}
	Next *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
	return &LinkedList{
	Value: val,
	Next: ll,
	}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}