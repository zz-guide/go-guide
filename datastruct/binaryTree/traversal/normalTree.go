package traversal

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/inorder"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	"go-guide/datastruct/binaryTree/traversal/morris"
	"go-guide/datastruct/binaryTree/traversal/postorder"
	"go-guide/datastruct/binaryTree/traversal/preorder"
	"go-guide/datastruct/binaryTree/treeNode"
)

func NormalTreeIteration() {
	root := treeNode.NewNormalTree()
	fmt.Println("根节点：", root.Val)

	fmt.Println("----------------前序遍历-------------------")
	preorderResult1 := preorder.TraversalRecursive(root)
	fmt.Println("前序遍历(递归1 根->左->右):", preorderResult1)

	preorderResult2 := preorder.TraversalRecursive1(root)
	fmt.Println("前序遍历(递归2 根->左->右):", preorderResult2)

	preorderResult3 := preorder.TraversalStack(root)
	fmt.Println("前序遍历(栈  1 根->左->右):", preorderResult3)

	preorderResult4 := preorder.TraversalStack1(root)
	fmt.Println("前序遍历(栈  2 根->左->右):", preorderResult4)

	fmt.Println("--------------中序遍历---------------------")
	inorderResult1 := inorder.TraversalRecursive(root)
	fmt.Println("中序遍历(递归1 左->根->右):", inorderResult1)

	inorderResult2 := inorder.TraversalRecursive1(root)
	fmt.Println("中序遍历(递归2 左->根->右):", inorderResult2)

	inorderResult3 := inorder.TraversalStack(root)
	fmt.Println("中序遍历(栈  1 左->根->右):", inorderResult3)

	inorderResult4 := inorder.TraversalStack1(root)
	fmt.Println("中序遍历(栈  2 左->根->右):", inorderResult4)

	fmt.Println("------------------后序遍历-----------------")
	postorderResult1 := postorder.TraversalRecursive(root)
	fmt.Println("后序遍历(递归1 左->右->根):", postorderResult1)

	postorderResult2 := postorder.TraversalRecursive1(root)
	fmt.Println("后序遍历(递归2 左->右->根):", postorderResult2)

	postorderResult3 := postorder.TraversalStack(root)
	fmt.Println("后序遍历(栈  1 左->右->根):", postorderResult3)

	postorderResult4 := postorder.TraversalStack1(root)
	fmt.Println("后序遍历(栈  2 左->右->根):", postorderResult4)

	postorderResult5 := postorder.TraversalStack2(root)
	fmt.Println("后序遍历(栈  3 左->右->根):", postorderResult5)

	fmt.Println("----------------Morris序列-------------------")
	morrisResult, preorderResult5, inorderResult5, postorderResult5, isBST := morris.TraversalMorris(root)
	fmt.Println("Morris序列:", morrisResult)
	fmt.Println("前序遍历(Morris 根->左->右):", preorderResult5)
	fmt.Println("中序遍历(Morris 左->根->右):", inorderResult5)
	fmt.Println("后序遍历(Morris 左->右->根):", postorderResult5)
	fmt.Println("是否搜索二叉树:", isBST)

	fmt.Println("----------------层序遍历-------------------")
	levelOrderResult1 := levelorder.TraversalRecursive(root)
	fmt.Println("层序遍历(递归):", levelOrderResult1)

	levelOrderResult2 := levelorder.TraversalQueue(root)
	fmt.Println("层序遍历(队列):", levelOrderResult2)
}
