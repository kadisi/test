package main

type Stack struct {
	cache []byte
}

func NewStack() *Stack {
	return &Stack{
		cache: make([]byte, 0, 2048),
	}
}

func (this *Stack) Put(c byte) {
	this.cache = append(this.cache, c)
}

func (this *Stack) Pop() {
	if this.Len() == 0 {
		return
	}
	this.cache = this.cache[:len(this.cache)-1]
	return
}

func (this *Stack) String() string {
	return string(this.cache)
}

func (this *Stack) Len() int {
	return len(this.cache)
}
