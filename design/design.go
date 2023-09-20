package main

//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

//OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

//SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

//SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

//Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

func main() {

}
