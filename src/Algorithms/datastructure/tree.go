package datastructure

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func NewNode(num int) *TreeNode {
	return &TreeNode{
		Val: num,
		Left: nil,
		Right: nil,
	}
}

func NewStringNode(data string) *TreeNode {
	if data == "#" {
		return nil
	}

	num, err := strconv.Atoi(data)
	if err != nil {
		fmt.Printf("strconv.atoi failed. data:%v, err:%v", data, err.Error())
		return nil
	}

	return &TreeNode{
		Val: num,
		Left: nil,
		Right: nil,
	}
}

func insert(root *TreeNode, num int) *TreeNode {
	if root == nil {
		return NewNode(num)
	}

	ptr := root
	lastptr := root
	for ptr != nil {
		lastptr = ptr
		if num >= ptr.Val {
			ptr = ptr.Right
		} else {
			ptr = ptr.Left
		}
	}
	if num >= lastptr.Val {
		lastptr.Right = NewNode(num)
	}else {
		lastptr.Left = NewNode(num)
	}
	return root
}

//创建二叉搜索树
func NewBinarySearchTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}

	root := (* TreeNode)(nil)
	for _,num := range nums {
		root = insert(root, num)
	}
	return root
}


//"1,2,#,#,3,4,5"
func NewTree(datas string) *TreeNode {
	nodes := strings.Split(datas,",")
	if len(nodes) <= 0 {
		return nil
	}

	fmt.Printf("nodes:%v\n", nodes)
	var root *TreeNode = nil
	i := 0
	root = CreateTree(nodes, &i, len(nodes))
	return root
}

func CreateTree(data []string, i *int, len int) *TreeNode {
	if *i >= len {
		return nil
	}
	root := NewStringNode(data[*i])
	(*i)++
	if root == nil {
		return nil
	}

	root.Left = CreateTree(data, i, len)
	root.Right = CreateTree(data, i, len)
	return root
}

func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		fmt.Printf("%v ",root.Val)
		inorder(root.Right)
	}
}

func preorder(root *TreeNode) {
	if root != nil {
		fmt.Printf("%v ",root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
}

func postorder(root *TreeNode) {
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		fmt.Printf("%v ", root.Val)
	}
}

func leverorder(root *TreeNode) {
	slice := make([]*TreeNode, 0)
	if root != nil {
		slice = append(slice, root)
	}

	for len(slice) > 0 {
		length := len(slice)
		for i := 0; i < length; i++ {
			node := slice[0]
			if node.Left != nil {
				slice = append(slice, node.Left)
			}
			if node.Right != nil {
				slice = append(slice, node.Right)
			}

			slice = slice[1:]
			fmt.Printf("%v ", node.Val)
		}
		fmt.Println()
	}
}

//逆层次遍历树节点---广度遍历
func reverselevelOrder(root *TreeNode) {
	slice := make([]*TreeNode, 0)
	if root != nil {
		slice = append(slice, root)
	}

	result := make([][]int, 0)
	for len(slice) > 0 {
		length := len(slice)
		tmpt := make([]int, 0)
		for i := 0; i < length; i++ {
			node := slice[0]
			if node.Left != nil {
				slice = append(slice, node.Left)
			}
			if node.Right != nil {
				slice = append(slice, node.Right)
			}

			slice = slice[1:]
			tmpt = append(tmpt, node.Val)
		}
		result = append(result, tmpt)
	}

	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
	fmt.Println(result)
}

//逆层次遍历树节点--深度遍历
func reverselevelOrder2(root *TreeNode) {
	if root == nil {
		return
	}

	result := make([][]int, 0)
	dfs(root, 0, &result)
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}
}

func dfs(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result)-1 < level {
		tmpt := make([]int, 0)
		*result = append(*result, tmpt)
	}
	(*result)[level] = append((*result)[level], root.Val)

	dfs(root.Left, level+1, result)
	dfs(root.Right, level+1, result)
}


//非递归版本
func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0)
	pre := (*TreeNode)(nil)
	cur := root

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == pre {
			fmt.Printf("%v ", cur.Val)
			pre = cur
			stack = stack[:len(stack)-1]
			cur = nil
		}else{
			cur = cur.Right
		}
	}
}

func inRangeLR(root *TreeNode, L int, R int, num *int) {
	if root != nil {
		if root.Val >= L && root.Val <= R {
			*num = *num + root.Val
			inRangeLR(root.Left, L, R, num)
			inRangeLR(root.Right, L, R, num)
		} else if root.Val < L {
			inRangeLR(root.Right, L, R, num)
		} else if root.Val > R {
			inRangeLR(root.Left, L, R, num)
		}
	}
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	num := 0
	inRangeLR(root, L, R, &num)
	return num
}

func mergeTrees1(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t2 == nil {
		return t1
	}
	if t1 == nil {
		return t2
	}

	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees1(t1.Left, t2.Left)
	t1.Right = mergeTrees1(t1.Right, t2.Right)
	return t1
}

func mergeTrees2(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	type TreeNodePair struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	stack := make([]*TreeNodePair, 0)
	stack = append(stack,&TreeNodePair{t1, t2})
	for len(stack) > 0 {
		nodepair := stack[0]
		stack = stack[1:]

		if nodepair.t1 == nil || nodepair.t2 == nil {
			continue
		}
		if nodepair.t1 != nil && nodepair.t2 != nil {
			nodepair.t1.Val = nodepair.t1.Val + nodepair.t2.Val
		}

		if nodepair.t1.Left == nil {
			nodepair.t1.Left = nodepair.t2.Left
		}else {
			stack = append(stack, &TreeNodePair{nodepair.t1.Left, nodepair.t2.Left})
		}

		if nodepair.t1.Right == nil {
			nodepair.t1.Right = nodepair.t2.Right
		}else {
			stack = append(stack, &TreeNodePair{nodepair.t1.Right, nodepair.t2.Right})
		}
	}
	return t1
}


func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

