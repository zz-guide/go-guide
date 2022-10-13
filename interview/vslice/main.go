package main

import "log"

func main() {
	var nums1 []interface{}
	nums2 := []interface{}{1, 2, 3}
	// 把nums2各项分别追加到nums1
	nums3 := append(nums1, nums2...)
	// 把nums2整体追加到nums1
	nums4 := append(nums1, nums2)
	log.Printf("nums3=%+v, len(nums3)=%d\n", nums3, len(nums3))
	log.Printf("nums4=%+v, len(nums4)=%d\n", nums4, len(nums4))
	log.Printf("nums1=%+v, len(nums1)=%d\n", nums1, len(nums1))
}
