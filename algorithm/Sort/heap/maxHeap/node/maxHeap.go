package node

import "fmt"

/**
关于堆的一些结论：二叉堆
	1.结点i的左子结点为2xi+1
	2.右子结点为2xi+2
	3.结点i的父节点为(i-1)/2
	4. 0~n/2 - 1的位置的数需要重新down,之后的数已经是最后一层的叶子结点了，可用来构建堆
	5.大小堆切换只需要修改less函数
	6.只有大堆和小堆
	7.堆始终是一颗完全二叉树
	8.左右结点谁大谁小不确定
	9.优先队列的实现

优先队列(priority queue)
普通的队列是一种先进先出的数据结构，元素在队列尾追加，而从队列头删除。在优先队列中，元素被赋予优先级。当访问元素时，具有最高优先级的元素最先删除。优先队列具有最高级先出 （first in, largest out）的行为特征。通常采用堆数据结构来实现。

*/

// MaxHeap 定义一个大堆结构
type MaxHeap []int

func (h MaxHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 判断i是不是比j小
func (h MaxHeap) less(i, j int) bool {
	// 大小堆切换改此处即可
	return h[i] < h[j]
}

func (h MaxHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h MaxHeap) leftChildIndex(i int) int {
	// 使用这个值的时候要注意，有可能超出len(h)
	return 2*i + 1
}

func (h MaxHeap) rightChildIndex(i int) int {
	// 使用这个值的时候要注意，有可能超出len(h)
	return 2*i + 2
}

func (h MaxHeap) BuildMaxHeap() {
	n := len(h)
	// 当h还不是堆的时候，i > n/2-1 的结点为叶子结点本身已经是堆了
	// 0~n/2 - 1的位置的数需要重新down,之后的数已经是最后一层的叶子结点了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *MaxHeap) Push(x int) {
	// 添加到数组末尾
	*h = append(*h, x)
	// 调用up方法构建大堆
	h.up(len(*h) - 1)
}

func (h MaxHeap) up(i int) {
	for {
		parentIndex := h.parentIndex(i)
		// 自身就是父节点，说明已经到顶部了
		if i == parentIndex {
			break
		}

		// 父节点比自己大，说明不需要交换了
		if h.less(parentIndex, i) {
			break
		}

		h.swap(parentIndex, i)
		// 再比较父节点的父节点
		i = parentIndex
	}
}

// Remove 删除堆中位置为i的元素，返回被删元素的值
func (h *MaxHeap) Remove(i int) (int, bool) {
	// 只能删除数组内的元素
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}

	// 1.先记录被删除的元素
	removeElem := (*h)[i]
	endIndex := len(*h) - 1
	// 2.把数组的最后一个元素填充到要删除的位置上
	(*h)[i] = (*h)[endIndex]
	*h = (*h)[0:endIndex]

	// 3.比父节点大就向上检查是否满足大堆要求
	if (*h)[i] > (*h)[h.parentIndex(i)] {
		h.up(i)
	} else { // 小就向下检查是否满足大堆要求
		h.down(i)
	}

	// 4.返回要删除的元素值和是否成功
	return removeElem, true
}

func (h MaxHeap) down(i int) {
	for {
		l := h.leftChildIndex(i)
		// 如果左孩子的索引已经越界了，就返回
		if l >= len(h) {
			break
		}

		// 右节点不能提前判断返回，左节点可以，右节点有些情况是没有的，完全二叉树就是这种特性
		r := l + 1
		// 假设孩子结点里最大的是left
		childMaxIndex := l
		if r < len(h) && h.less(l, r) {
			childMaxIndex = r
		}

		// 父节点大就停止交换
		if h.less(childMaxIndex, i) {
			break
		}

		h.swap(i, childMaxIndex)
		i = childMaxIndex
	}
}

// Pop 弹出堆顶的元素，并返回其值
// 当i=0时，Remove 就是 Pop
// 每次只能删除数组最后一个元素
func (h *MaxHeap) Pop() int {
	// 0永远是堆顶元素
	heapTop := (*h)[0]

	endIndex := len(*h) - 1
	// 从数组里去掉最后一个元素
	(*h)[0] = (*h)[endIndex]
	*h = (*h)[0:endIndex]
	// 重新构建堆
	h.down(0)
	return heapTop
}

func (h MaxHeap) Desc() []int {
	var res []int
	for len(h) > 0 {
		res = append(res, h.Pop())
	}

	return res
}

func (h MaxHeap) Asc() []int {
	var res = make([]int, len(h))
	for i := len(h) - 1; i >= 0; i-- {
		res[i] = h.Pop()
	}

	return res
}

func TestMaxHeap() {
	arr := []int{20, 30, 90, 40, 70, 110, 60, 10, 100, 50, 80}
	arr1 := []int{20, 30, 90, 40, 70, 110, 60, 10, 100, 50, 80}
	fmt.Println("排序前:", arr)
	hDesc := MaxHeap(arr)
	hDesc.BuildMaxHeap()
	fmt.Println("堆排序(降序):", hDesc.Desc())

	hAsc := MaxHeap(arr1)
	hAsc.BuildMaxHeap()
	fmt.Println("堆排序(升序)", hAsc.Asc())
}
