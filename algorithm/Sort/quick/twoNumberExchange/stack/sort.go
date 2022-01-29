package main

import (
	"log"
)

func main() {
	arr := []int{9, 0, 7, 1, 4, 8, 4, 9, 9, 7}
	log.Println("快速排序(迭代法):", sortArray(arr))
}

// Stack 定义栈
type Stack struct {
	S []*Index
}

// Index 栈内元素结构
type Index struct {
	Left  int
	Right int
}

// Push 入栈
func (s *Stack) Push(l, r int) {
	index := Index{
		Left:  l,
		Right: r,
	}
	s.S = append(s.S, &index)
}

// Pop 出栈
func (s *Stack) Pop() *Index {
	if len(s.S) == 0 {
		return nil
	}
	rt := s.S[0]
	s.S = s.S[1:]
	return rt
}

func quickSort(l []int) {
	length := len(l)
	var stack Stack
	stack.Push(0, length-1)
	for {
		s := stack.Pop()
		if s == nil {
			break
		}

		if s.Left <= s.Right {
			i := s.Left
			j := s.Right
			pivot := l[(s.Left+s.Right)/2]

			for i <= j {
				for l[i] < pivot {
					i++
				}
				for l[j] > pivot {
					j--
				}
				if i <= j {
					l[i], l[j] = l[j], l[i]
					i++
					j--
				}
			}

			if s.Left < j {
				stack.Push(s.Left, j)
			}
			if s.Right > i {
				stack.Push(i, s.Right)
			}
		}
	}
}

func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}
