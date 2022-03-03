package node

import "fmt"

//红黑树的定义和实现
const (
	Red   = true
	Black = false
)

type node struct {
	key   int
	value int
	color bool
	left  *node
	right *node
}

type redBlackTree struct {
	size int
	root *node
}

func newNode(key, val int) *node {
	// 默认添加红节点
	return &node{key, val, Red, nil, nil}
}

func NewRedBlackTree() *redBlackTree {
	return new(redBlackTree)
}

func (nd *node) isRed() bool {
	if nd == nil {
		return Black
	}
	return nd.color
}

func (tree *redBlackTree) GetSize() int {
	return tree.size
}

// 向红黑树中添加元素
func (tree *redBlackTree) Add(key, val int) {
	isAdd, nd := tree.root.add(key, val)
	tree.size += isAdd
	tree.root = nd
	tree.root.color = Black //根节点为黑色节点
}

// 递归写法:向树的root节点中插入key,val
// 返回1,代表加了节点
// 返回0,代表没有添加新节点,只更新key对应的value值
func (nd *node) add(key, val int) (int, *node) {
	if nd == nil { // 默认插入红色节点
		return 1, newNode(key, val)
	}

	isAdd := 0
	if key < nd.key {
		isAdd, nd.left = nd.left.add(key, val)
	} else if key > nd.key {
		isAdd, nd.right = nd.right.add(key, val)
	} else { // nd.key == key
		// 对value值更新,节点数量不增加,isAdd = 0
		nd.value = val
	}

	// 维护红黑树
	nd = nd.updateRedBlackTree(isAdd)
	return isAdd, nd
}

// 红黑树维护
func (nd *node) updateRedBlackTree(isChange int) *node {
	// 0说明无新节点,不必维护
	if isChange == 0 {
		return nd
	}

	// 维护
	// 判断是否为情形2，是需要左旋转
	if nd.right.isRed() == Red && nd.left.isRed() != Red {
		nd = nd.leftRotate()
	}

	// 判断是否为情形3，是需要右旋转
	if nd.left.isRed() == Red && nd.left.left.isRed() == Red {
		nd = nd.rightRotate()
	}

	// 判断是否为情形4，是需要颜色翻转
	if nd.left.isRed() == Red && nd.right.isRed() == Red {
		nd.flipColors()
	}
	return nd
}

//    nd                      x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (nd *node) leftRotate() *node {
	// 左旋转
	retNode := nd.right
	nd.right = retNode.left

	retNode.left = nd
	retNode.color = nd.color
	nd.color = Red

	return retNode
}

//      nd                    x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func (nd *node) rightRotate() *node {
	//右旋转
	retNode := nd.left
	nd.left = retNode.right
	retNode.right = nd

	retNode.color = nd.color
	nd.color = Red

	return retNode
}

// 颜色翻转
func (nd *node) flipColors() {
	nd.color = Red
	nd.left.color = Black
	nd.right.color = Black
}

// 前序遍历打印出key,val,color
func (tree *redBlackTree) PrintPreOrder() {
	resp := [][]interface{}{}
	tree.root.printPreOrder(&resp)
	fmt.Println(resp)
}

func (nd *node) printPreOrder(resp *[][]interface{}) {
	if nd == nil {
		return
	}
	*resp = append(*resp, []interface{}{nd.key, nd.value, nd.color})
	nd.left.printPreOrder(resp)
	nd.right.printPreOrder(resp)
}
