package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func toString(l *ListNode) string {
	vals := make([]string, 0)
	next := l
	for next != nil {
		vals = append(vals, strconv.Itoa(next.Val))
		if next.Next == nil {
			break
		}
		next = next.Next
	}
	return strings.Join(vals, ",")
}

func push(l *ListNode, val int) {
	next := l
	for next != nil {
		if next.Next == nil {
			break
		}
		next = next.Next
	}
	next.Next = &ListNode{Val: val, Next: nil}
}

func get(l *ListNode, i int) int {
	counter := 0
	next := l
	for next != nil {
		if i == counter {
			return next.Val
		}
		if next.Next == nil {
			break
		} else {
			next = next.Next
			counter++
		}
	}
	return -1
}

func size(l *ListNode) int {
	counter := 0
	next := l
	for next != nil {
		counter++
		if next.Next == nil {
			break
		}
		next = next.Next
	}
	return counter
}

type Walker struct {
	x int
	y int
}

type Holder struct {
	result *ListNode
}

func walk(lists []*ListNode, walkers []*Walker, holder *Holder) {
	minVal := math.MaxInt
	var minWalker *Walker
	for i := 0; i < len(walkers); i++ {
		walker := walkers[i]

		val := get(lists[walker.x], walker.y)
		if val < minVal {
			minVal = val
			minWalker = walker
		}
	}

	if minWalker != nil {
		val := get(lists[minWalker.x], minWalker.y)
		if holder.result == nil {
			holder.result = &ListNode{Val: val}
		} else {
			push(holder.result, val)
		}

		if minWalker.y+1 < size(lists[minWalker.x]) {
			minWalker.y++
		} else {
			var newWalkers []*Walker
			for _, walker := range walkers {
				if walker != minWalker {
					newWalkers = append(newWalkers, walker)
				}
			}
			walkers = newWalkers
		}

		walk(lists, walkers, holder)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	walkers := []*Walker{}
	for i := range len(lists) {
		if lists[i] != nil {
			walkers = append(walkers, &Walker{i, 0})
		}
	}
	holder := &Holder{}
	walk(lists, walkers, holder)
	return holder.result
}

func main() {
	lists := make([]*ListNode, 0)
	lists = append(lists, &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}})
	lists = append(lists, &ListNode{Val: 10, Next: &ListNode{Val: 30, Next: &ListNode{Val: 40, Next: nil}}})
	lists = append(lists, &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}})
	result := mergeKLists(lists)
	fmt.Printf("%v", toString(result))
}
