package datastructure

import (
	"fmt"
	"testing"
)

func Test_tree(t *testing.T) {
	nums := []int{1,2,3,77,6,5,9}
	root := NewBinarySearchTree(nums)
	inorder(root)
}

func Test_rangeSumBST(t *testing.T) {
	nums := []int{1,2,3,77,6,5,9}
	root := NewBinarySearchTree(nums)
	sum := rangeSumBST(root,2,10)
	t.Logf("nums:%v, sum:%v",nums, sum)
}

func Test_invertTree(t *testing.T) {
	nums := []int{15,5,20,1,6,30,40}
	root := NewBinarySearchTree(nums)
	postorderTraversal(root)
	fmt.Printf("\n")
	root = invertTree(root)
	postorderTraversal(root)
	fmt.Printf("\n")
}

func Test_leverorder(t *testing.T) {
	data := "1,2,7,#,#,8,#,#,3,4,#,1,#,#,5,6"
	root := NewTree(data)
	reverselevelOrder2(root)
}
