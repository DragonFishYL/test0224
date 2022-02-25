package main

import "testing"

func TestNode(t *testing.T) {
	sortSlice = []int{2, 55, 3, 34, 6, 7, 24, 78, 10}
	no := Node{}
	no.GenerNode()
	no.InGenera()
}
