package main

import (
	"fmt"
	"math/big"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertList(num int, l *ListNode) *ListNode {
	if l.Val == 0 && l.Next == nil {
		return &ListNode{num, nil}
	}
	return &ListNode{num, l}

}

func printList(l *ListNode) {
	fmt.Println(convListToArr(l))
}

func convArrToRevNum(nums []int) *big.Int {

	var x = big.NewInt(0)
	for i, num := range nums {
		a := big.NewInt(int64(num))
		b := big.NewInt(int64(i))
		c := big.NewInt(int64(10))

		pow := c.Exp(c, b, nil)
		x.Add(x, a.Mul(a, pow))
	}
	return x
}

func convListToArr(l1 *ListNode) []int {
	var list []int
	for true {
		list = append(list, l1.Val)
		if l1.Next == nil {
			break
		} else {
			l1 = l1.Next
		}
	}
	return list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var list1 = convListToArr(l1)
	var list2 = convListToArr(l2)
	fmt.Println(list1, list2)
	sumVal := convArrToRevNum(list1).Add(convArrToRevNum(list1), convArrToRevNum(list2))

	s := (len(sumVal.String()))

	var x = &ListNode{}
	tn := big.NewInt(int64(10))
	for i := s - 1; i >= 0; i-- {
		ipow := big.NewInt(int64(i))
		pow := ipow.Exp(tn, ipow, nil)
		powsv := pow.Div(sumVal, pow)
		mod := powsv.Mod(powsv, tn)
		fmt.Println(i, mod)
		x = insertList(int(mod.Int64()), x)

	}
	return x
}

func convArrToList(nums []int) *ListNode {
	var x = &ListNode{}
	for _, num := range nums {
		x = insertList(num, x)
	}
	xi := convListToArr(x)
	x = &ListNode{}
	for _, num := range xi {
		x = insertList(num, x)
	}
	return x
}

func main() {

	aa := []int{2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9}
	bb := []int{5, 6, 4, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9, 9, 9, 9}

	aal := convArrToList(aa)
	bbl := convArrToList(bb)

	// a := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// b := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	cc := addTwoNumbers(aal, bbl)
	printList(cc)

}

// INPUT
// [2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9]
// [5,6,4,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,2,4,3,9,9,9,9]

// OUTPUT
// [7,0,8,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,4,8,6,1,4,3,9,1]
