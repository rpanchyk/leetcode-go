package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Definition for singly-linked list
type ListNode struct {
	fmt.Stringer
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
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

func (l *ListNode) push(val int) {
	next := l
	for next != nil {
		if next.Next == nil {
			break
		}
		next = next.Next
	}
	next.Next = &ListNode{Val: val, Next: nil}
}

func (l *ListNode) get(i int) int {
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

func (l *ListNode) len() int {
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

		val := lists[walker.x].get(walker.y)
		if val < minVal {
			minVal = val
			minWalker = walker
		}
	}

	if minWalker != nil {
		val := lists[minWalker.x].get(minWalker.y)
		if holder.result == nil {
			holder.result = &ListNode{Val: val}
		} else {
			holder.result.push(val)
		}

		if minWalker.y + 1 >= lists[minWalker.x].len() {
			var newWalkers []*Walker
			for _, walker := range walkers {
				if walker != minWalker {
					newWalkers = append(newWalkers, walker)
				}
			}
			walkers = newWalkers
		} else {
			minWalker.y++
		}

		walk(lists, walkers, holder)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	walkers := []*Walker{}
	for i := range len(lists) {
		walkers = append(walkers, &Walker{i, 0})
	}
	holder := &Holder{}
	walk(lists, walkers, holder)
	return holder.result
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var result *ListNode
	ids := make([]int, len(lists))
	realIds := make([]int, len(lists))
	// vals := make([]int, len(lists))

	maxLength := -1
	for _, list := range lists {
		if maxLength < list.len() {
			maxLength = list.len()
		}
	}

	for i := 0; i < maxLength; i++ {
		//values := make([]int, 0)

		for j := 0; j < len(lists); j++ {
			list := lists[j]

			if i < list.len() {
				ids[j] = i
				// vals[j] = list.get(i)
			} else {
				ids[j] = -1
			}

			
		}

		// fmt.Printf("%v\n", ids)
		// fmt.Printf("%v\n\n", vals)

		minJ := -1
		minId := math.MaxInt
		for j := 0; j < len(ids); j++ {
			if minId > lists[j].get(ids[j]) {
				minJ = j
				minId = lists[j].get(ids[j])
			}
		}

		minJnext := math.MaxInt
		if minJ != -1 && minId + 1 < lists[minJ].len() {
			minJnext = lists[minJ].get(minId + 1)
		}

		for j := 0; j < len(ids); j++ {
			if j == minJ {
				realIds[j] = j
				continue
			}
			if lists[j].get(ids[j]) < minJnext {
				realIds[j] = j
			}
		}

		fmt.Printf("%v\n", realIds)
	}

	// for _, list := range lists {
	// 	if result == nil {
	// 		result = &ListNode{Val: list.Val, Next: nil}
	// 	} else {
	// 		// next := list
	// 		// for next != nil {
	// 		// 	if next.Next == nil {
	// 		// 		break
	// 		// 	}
	// 		// 	next = next.Next
	// 		// }
	// 		result.push(list.Val)
	// 	}
	// }

	return result
}

func main() {
	lists := make([]*ListNode, 0)
	lists = append(lists, &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}})
	lists = append(lists, &ListNode{Val: 10, Next: &ListNode{Val: 30, Next: &ListNode{Val: 40, Next: nil}}})
	lists = append(lists, &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}})
	result := mergeKLists(lists)
	fmt.Printf("%v", result)
}
