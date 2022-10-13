package main

func main() {
	/**
	1.GC种类
		手动：C语言

		自动：php,go,java等



	2.gc主要是针对堆内存而言的

	3.数据“可达性”近似等价于数据有用性的
	原因：虽然能够追踪得到不代表后续一定会用到，但是从这些根节点追踪不到的数据，一定不会被用到，也就一定是无用的垃圾


	4.STW(STOP THE WORLD)
		标记——清扫

		标记——压缩（整理）:防止碎片化
		分代回收
		引用计数

	5.增量式垃圾回收
	增量式垃圾回收是指用户程序与垃圾回收交替执行，将垃圾回收工作分多次完成，也将暂停的时间分摊到多次，进而缩短每次暂停的时间。
	缺点：但是这也带来了额外的问题，交替执行的过程中，保不齐垃圾回收程序前脚刚把一个变量标记为垃圾，用户程序后脚又用到了它。
	若是放任不管，垃圾回收程序就会把有用数据“误判”为垃圾，进而影响程序正常执行。
	若用户程序对回收相关节点进行了写操作，通常的办法就是建立写屏障。

	6.写屏障
	写屏障会在写操作中插入指令，目的是把数据对象的修改通知到垃圾回收器。所以写屏障通常都要有一个记录集。

	7.读屏障
	读操作在非移动式垃圾回收器中无甚影响，但是在复制式回收器或者压缩回收器中，由于会移动数据来避免碎片化，所以垃圾回收器和用户程序交替执行时，读数据便也不那么安全了。

	8.三色抽象
	实际上，“三色抽象”适用于描述多种垃圾回收算法的推进过程：
	（1）最初所有数据都可描述为白色对象，代表尚未处理；
	（2）灰色对象代表尚未处理完；
	（3）黑色对象表示处理结束且不为垃圾。
	所以不要把给对象着色狭隘的理解为标记——清扫类算法中修改数据颜色标记的操作，着色操作可以是广义的。

	9.弱三色不变式
	我们知道，黑色对象是不会被回收器重新处理的，而会被回收器处理的灰色对象又不能抵达这个白色对象，那么它就会被作为垃圾回收，但实际上它是可达的。

	如果能保障被黑色对象引用的白色对象处在灰色对象的可达路径内，就能保护它不被错误地回收了，这被称为“弱三色不变式”。
	若直接不允许存在黑色对象到白色对象的引用，那就更安全了，这被称为“强三色不变式”。
	*/

	/**
	Golang中垃圾回收支持三种模式：
	（1）gcBackgroundMode，默认模式，标记与清扫过程都是并发执行的；
	（2）gcForceMode，只在清扫阶段支持并发；
	（3）gcForceBlockMode，GC全程需要STW。
	*/
}