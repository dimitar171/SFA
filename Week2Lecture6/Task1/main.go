package main

import (
	"fmt"
)

type MagicList struct {
	LastItem *Item
}

type Item struct {
	Value    int
	PrevItem *Item
}

func add(l *MagicList, val int) {
	item := &Item{Value: val}

	if l.LastItem == nil {
		l.LastItem = item
	} else {
		item.PrevItem = l.LastItem
		l.LastItem = item
	}
}

var i []int

func ReverseLinkedList(head *Item) []int {
	if head != nil {
		i = append(i, head.Value)
		ReverseLinkedList(head.PrevItem)

	}
	return i
}

func main() {
	ml := &MagicList{}
	add(ml, 10)
	add(ml, 15)
	add(ml, 20)
	add(ml, 30)
	add(ml, 40)
	add(ml, 50)
	add(ml, 60)
	add(ml, 70)
	add(ml, 80)
	fmt.Println(ReverseLinkedList(ml.LastItem))
}
