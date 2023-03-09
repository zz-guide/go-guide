package main

func main() {
	/**
		栈内存的分配
	1.其实，span除了用作堆内存分配外，也用于栈内存分配，只是用途不同的span对应的mspan状态不同。用做堆内存的mspan状态为mSpanInUse，而用做栈内存的状态为mSpanManual。
	为提高栈内存分配效率，调度器初始化时会初始化两个用于栈分配的全局对象：stackpool 和stackLarge。


	stackpool面向32KB以下的栈分配，栈大小必须是2的幂，最小2KB，在Linux环境下，stackpool提供了2kB、4KB、8KB、16KB四种规格的mspan链表。

	大于等于32KB的栈，由stackLarge来分配，这也是个mspan链表的数组，长度为25。

	mspan规格从8KB开始，之后每个链表的mspan规格，都是前一个的两倍。
	8KB和16KB这两个链表，实际上会一直是空的，留着它们是为了方便使用mspan包含页面数的（以2为底）对数作为数组下标。


























































	*/
}
