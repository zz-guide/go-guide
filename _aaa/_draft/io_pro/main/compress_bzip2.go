﻿package main

import (
	"compress/bzip2"
	"fmt"
	"os"
)

func main2352() {
	//file, e := os.Open("main/bzip2/e.txt.bz2")
	file, e := os.Open("main/bzip2/aa.txt.bz2")
	check_err_bzip2(e)
	// NewReader返回一个io.Reader，它将io中的bzip2数据解压缩。
	//如果r还没有实现io.ByteReader，则解压缩器可能会从r读取比所需更多的数据。
	r_bz2 := bzip2.NewReader(file)

	//info, e := file.Stat()
	check_err_bzip2(e)
	//dst_byte:=make([]byte,info.Size())//这里获取的是压缩后的字节的大小，但是容纳的却是解压后的字节的大小，是不一样的额！
	//dst_byte:=make([]byte,102400)//给一个很大的数就行，事实上他的真实字节数是100003，对于这种有没有更好的解决办法呢？
	dst_byte:=make([]byte,600)//这个大小主要是为了解压aa.txt.bz2用的，如果是e.txt.bz2请用上面的大小
	// 目前还不知道，一般[]byte大小对于这种格式给源bzip2文件的大小的2.2倍即可
	n, e := r_bz2.Read(dst_byte)
	check_err_bzip2(e)
	fmt.Println("解压了的字节数是：",n)
	fmt.Println("解压了的字节是：",dst_byte)
	fmt.Println("解压了的字符串是：",string(dst_byte))
	//e.txt.bz2输出：
	//	解压了的字节数是： 43149
	//	解压了的字节是： [50 46 55 49 56 50 56 49 56 50 56 52 53 57 48 52 53 50 51 53 ...](这里我省略了，因为太长了)
	//	解压了的字符串是： 2.718281828459045235360287471352662497757247093699959574966967...(这里我省略了，因为太长了)
	//aa.txt.bz2输出：
	//	解压了的字节数是： 577
	//	解压了的字节是： [98 99 50 48 49 57 229 185 180 49 48 230 156 136 49 53 230 151 165 239 188 140 229 140 151 228 186 172 227 128 130 32
	//	233 184 189 229 173 144 228 187 142 232 145 163 229 133 136 231 148 159 231 154 132 233 184 189 231 172 188 228 184 173 233 163 158
	//	229 135 186 239 188 140 232 145 163 229 133 136 231 148 159 230 152 175 229 164 167 229 133 180 228 191 161 233 184 189 229 141 143
	//	228 188 154 230 136 144 229 145 152 239 188 140 232 191 153 230 172 161 228 184 186 229 155 189 229 186 134 230 180 187 229 138 168
	//	230 143 144 228 190 155 228 186 134 228 184 128 231 153 190 229 164 154 228 184 170 233 184 189 229 173 144 227 128 130 49 48 230 156
	//	136 49 230 151 165 231 154 132 230 150 176 228 184 173 229 155 189 230 136 144 231 171 139 55 48 229 145 168 229 185 180 229 186 134
	//	229 133 184 228 184 138 230 148 190 233 163 158 228 186 134 55 228 184 135 231 190 189 229 146 140 229 185 179 233 184 189 239 188 140
	//	229 156 186 233 157 162 232 148 154 228 184 186 229 163 174 232 167 130 227 128 130 10 230 141 174 228 186 134 232 167 163 239 188 140
	//	232 191 153 228 186 155 229 146 140 229 185 179 233 184 189 233 131 189 230 152 175 229 140 151 228 186 172 228 191 161 233 184 189 229
	//	141 143 228 188 154 228 187 142 233 184 189 229 143 139 229 164 132 229 190 129 233 155 134 232 128 140 230 157 165 227 128 130 229 156
	//	168 229 185 191 229 156 186 233 155 134 228 189 147 230 148 190 233 163 158 229 144 142 239 188 140 232 191 153 228 186 155 228 191 161
	//	233 184 189 231 187 157 229 164 167 233 131 168 229 136 134 233 131 189 229 155 158 229 136 176 228 186 134 229 174 182 228 184 173 239
	//	188 140 232 128 140 232 191 153 228 186 155 228 191 161 233 184 189 229 185 179 230 151 182 229 156 168 229 174 182 228 184 173 230 152
	//	175 229 166 130 228 189 149 232 174 173 231 187 131 229 146 140 231 148 159 230 180 187 231 154 132 229 145 162 239 188 159 229 140 151
	//	228 186 172 233 157 146 229 185 180 230 138 165 232 174 176 232 128 133 233 135 135 232 174 191 228 186 134 233 131 168 229 136 134 233
	//	184 189 229 143 139 239 188 140 230 143 173 231 167 152 228 191 161 233 184 189 231 154 132 230 151 165 229 184 184 231 148 159 230 180
	//	187 229 146 140 232 174 173 231 187 131 227 128 130 228 190 155 229 155 190 239 188 154 232 167 134 232 167 137 228 184 173 229 155 189
	//	97 98 99 10 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	//	解压了的字符串是： bc2019年10月15日，北京。 鸽子从董先生的鸽笼中飞出，董先生是大兴信鸽协会成员，这次为国庆活动提供了一百多个鸽子。10月1日的新中国成立70周年庆典上放飞了7万羽和平鸽，场面蔚为壮观。
	//	据了解，这些和平鸽都是北京信鸽协会从鸽友处征集而来。在广场集体放飞后，这些信鸽绝大部分都回到了家中，而这些信鸽平时在家中是如何训练和生活的呢？北京青年报记者采访了部分鸽友，揭秘信鸽的日常生活和训练。供图：视觉中国abc

	//目前go没提供压缩成这种格式的api,只能解压！
	//如果你希望得到自己的这种格式的压缩文件进行测试，请到linux虚拟机中：
	//	1.命令“yum -y install bzip2”安装bzip2包（前提要有yum包）
	//	2.创建你需要压缩的文件
	//	3.命令“bzip2 -k 你要压缩的完整文件名”,这样的话就可以得到对应的bzip2压缩文件了

}
func check_err_bzip2(err2 error)  {
	if err !=nil{
		fmt.Println(err)
	}
}