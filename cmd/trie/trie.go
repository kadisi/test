package main

import (
	"bufio"
	"fmt"
	"io"
)

type Node struct {
	Root  bool
	Count uint32
	C     rune
	List  []*Node
}

func NewNode(root bool) *Node {
	node := Node{}
	node.List = make([]*Node, 26, 26)
	node.Root = root
	return &node
}

func ExpendTrieTree(root *Node, words string) error {
	if root == nil {
		panic("root of tree is nil")
	}
	// change to lower case
	words = Lower(words)
	//fmt.Printf("Add words [%v]\n", words)
	cur := root
	var index int
	for _, c := range words {
		index = GetIndex(rune(c))
		if cur.List[index] == nil {
			newNode := NewNode(false)
			newNode.C = c
			cur.List[index] = newNode
		}
		cur = cur.List[index]
	}

	cur.Count++
	return nil
}

func GetCounts(root *Node, words string) uint32 {
	words = Lower(words)
	cur := root
	for _, c := range words {
		index := GetIndex(rune(c))
		if cur.List[index] == nil {
			return 0
		}
		cur = cur.List[index]
	}
	return cur.Count
}

func TravelTree(n *Node, stack *Stack) {
	if n == nil {
		return
	}

	if !n.Root {
		stack.Put(byte(n.C))
	}
	if n.Count != 0 && !n.Root {
		fmt.Printf("[%s] count %d\n", stack.String(), n.Count)
	}
	for _, child := range n.List {
		TravelTree(child, stack)
	}
	if !n.Root {
		stack.Pop()
	}
}

func CreateTree(r io.Reader) *Node {

	scan := bufio.NewScanner(r)
	scan.Split(SplitEnglishWords)

	tree := NewNode(true)
	for scan.Scan() {
		ExpendTrieTree(tree, scan.Text())
	}
	return tree
}
