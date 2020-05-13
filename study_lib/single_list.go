package study_lib

import (
	"strconv"
)

type SingleList struct {
	Head *Node
	Tail *Node
}

func (this *SingleList) String() (s string) {

	pNode := this.Head
	for {
		if pNode == nil {
			break
		}

		s = s + strconv.Itoa(pNode.Data) + "\n"

		pNode = pNode.Next
	}

	return
}

func (this *SingleList) Init() {
	this.Head = nil
	this.Tail = nil
}

func (this *SingleList) PushTail(node *Node) *Node {

	node.Next = nil

	if this.Tail == nil {
		this.Tail = node
		this.Head = node
		return node
	}

	this.Tail.Next = node

	this.Tail = node

	return node
}

func (this *SingleList) PushHead(node *Node) *Node {

	node.Next = this.Head

	this.Head = node

	return node
}

func New() *SingleList {
	return new(SingleList)
}

type Node struct {
	Data int
	Next *Node
}
