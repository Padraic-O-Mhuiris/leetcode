package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfsHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := dfsHeight(root.Left)
	if lh == -1 {
		return -1
	}

	rh := dfsHeight(root.Right)
	if rh == -1 {
		return -1
	}

	if math.Abs(float64(lh-rh)) > 1 {
		return -1
	}

	return int(math.Max(float64(lh), float64(rh))) + 1
}

func isBalanced(root *TreeNode) bool {
	return dfsHeight(root) != -1
}

func (t *TreeNode) Insert(val int) error {
	if t == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	case val == t.Val:
		return nil
	case val <= t.Val:
		if t.Left == nil {
			t.Left = &TreeNode{val, nil, nil}
			return nil
		}
		return t.Left.Insert(val)

	case val > t.Val:
		if t.Right == nil {
			t.Right = &TreeNode{val, nil, nil}
			return nil
		}
		return t.Right.Insert(val)
	}

	return nil
}

func stringify(n *TreeNode, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "    "
		}
		format += "-- "
		level++
		stringify(n.Left, level)
		fmt.Printf(format+"%d\n", n.Val)
		stringify(n.Right, level)
	}
}

func main() {
	r := &TreeNode{5, nil, nil}

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "-1" {
			break
		}
		x, err := strconv.Atoi(text)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		r.Insert(x)
		stringify(r, 0)
		fmt.Println(isBalanced(r))
	}
}
