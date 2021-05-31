package offer

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	root := buildTree(preorder, inorder)
	preorderN := preBST(root)
	inorderN := inBST(root)
	if !compareTwoArray(preorderN, preorder) && !compareList(inorderN, inorder) {
		t.Errorf("test error buildTree()")
	}

	var preorder2 []int
	var inorder2 []int
	root2 := buildTree(preorder2, inorder2)
	preorderN2 := preBST(root2)
	inorderN2 := inBST(root2)
	if !compareTwoArray(preorder2, preorderN2) && !compareList(inorderN2, inorder2){
		t.Errorf("test error buildTree()")
	}
}

func compareTwoArray(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return  false
	}

	for i := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func preBST(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var bst func(node *TreeNode)
	bst = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		bst(node.Left)
		bst(node.Right)
	}
	bst(root)
	return res
}

func inBST(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var bst func(node *TreeNode)
	bst = func(node *TreeNode) {
		if node == nil {
			return
		}
		bst(node.Left)
		res = append(res, node.Val)
		bst(node.Right)
	}
	bst(root)
	return res
}
