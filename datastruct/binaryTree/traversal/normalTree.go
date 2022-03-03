package traversal

import (
	"go-guide/datastruct/binaryTree/traversal/inorder"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	"go-guide/datastruct/binaryTree/traversal/morris"
	"go-guide/datastruct/binaryTree/traversal/postorder"
	"go-guide/datastruct/binaryTree/traversal/preorder"
	"go-guide/datastruct/binaryTree/treeNode"
	"log"
)

func NormalTreeIteration() {
	root := treeNode.NewNormalTree()
	log.Println("根节点：", root.Val)

	log.Println("----------------前序遍历-------------------")
	preorderResult1 := preorder.TraversalRecursive(root)
	log.Println("前序遍历(递归1 根->左->右):", preorderResult1)

	preorderResult2 := preorder.TraversalRecursive1(root)
	log.Println("前序遍历(递归2 根->左->右):", preorderResult2)

	preorderResult3 := preorder.TraversalStack(root)
	log.Println("前序遍历(栈  1 根->左->右):", preorderResult3)

	preorderResult4 := preorder.TraversalStack1(root)
	log.Println("前序遍历(栈  2 根->左->右):", preorderResult4)

	log.Println("--------------中序遍历---------------------")
	inorderResult1 := inorder.TraversalRecursive(root)
	log.Println("中序遍历(递归1 左->根->右):", inorderResult1)

	inorderResult2 := inorder.TraversalRecursive1(root)
	log.Println("中序遍历(递归2 左->根->右):", inorderResult2)

	inorderResult3 := inorder.TraversalStack(root)
	log.Println("中序遍历(栈  1 左->根->右):", inorderResult3)

	inorderResult4 := inorder.TraversalStack1(root)
	log.Println("中序遍历(栈  2 左->根->右):", inorderResult4)

	log.Println("------------------后序遍历-----------------")
	postorderResult1 := postorder.TraversalRecursive(root)
	log.Println("后序遍历(递归1 左->右->根):", postorderResult1)

	postorderResult2 := postorder.TraversalRecursive1(root)
	log.Println("后序遍历(递归2 左->右->根):", postorderResult2)

	postorderResult3 := postorder.TraversalStack(root)
	log.Println("后序遍历(栈  1 左->右->根):", postorderResult3)

	postorderResult4 := postorder.TraversalStack1(root)
	log.Println("后序遍历(栈  2 左->右->根):", postorderResult4)

	postorderResult5 := postorder.TraversalStack2(root)
	log.Println("后序遍历(栈  3 左->右->根):", postorderResult5)

	log.Println("----------------Morris序列-------------------")
	morrisResult, preorderResult5, inorderResult5, postorderResult5, isBST := morris.TraversalMorris(root)
	log.Println("Morris序列:", morrisResult)
	log.Println("前序遍历(Morris 根->左->右):", preorderResult5)
	log.Println("中序遍历(Morris 左->根->右):", inorderResult5)
	log.Println("后序遍历(Morris 左->右->根):", postorderResult5)
	log.Println("是否搜索二叉树:", isBST)

	log.Println("----------------层序遍历-------------------")
	levelOrderResult1 := levelorder.TraversalRecursive(root)
	log.Println("层序遍历(递归):", levelOrderResult1)

	levelOrderResult2 := levelorder.TraversalQueue(root)
	log.Println("层序遍历(队列):", levelOrderResult2)
}
