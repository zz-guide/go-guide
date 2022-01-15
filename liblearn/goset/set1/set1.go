package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	setTest()
}

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{m: make(map[interface{}]bool)}
}

func (self *Set) Add(e interface{}) bool {
	self.Lock()
	defer self.Unlock()
	if self.m[e] {
		return false
	}

	self.m[e] = true
	return true
}

func (self *Set) Remove(e interface{}) bool {
	self.Lock()
	defer self.Unlock()
	delete(self.m, e)
	return true
}

func (self *Set) Clear() bool {
	self.Lock()
	defer self.Unlock()
	self.m = make(map[interface{}]bool)
	return true
}

func (self *Set) Contains(e interface{}) bool {
	self.Lock()
	defer self.Unlock()
	//return self.m[e]
	_, ok := self.m[e]
	return ok
}

func (self *Set) IsEmpty() bool {
	return self.Len() == 0
}

func (self *Set) Len() int {
	self.Lock()
	defer self.Unlock()
	return len(self.m)
}

func (self *Set) Same(other *Set) bool {
	if other == nil {
		return false
	}

	if self.Len() != other.Len() {
		return false
	}

	for k, _ := range other.m {
		if !self.Contains(k) {
			return false
		}
	}
	return true
}

func (self *Set) Elements() interface{} {
	self.Lock()
	defer self.Unlock()
	// for k := range self.m{
	//    snapshot = snapshot(snapshot, k)
	// }
	initialLen := self.Len()
	actualLen := 0
	snapshot := make([]interface{}, initialLen)
	for k := range self.m {
		if actualLen < initialLen {
			snapshot[actualLen] = k
		} else {
			snapshot = append(snapshot, k)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (self *Set) String() string {
	self.Lock()
	defer self.Unlock()
	var buf bytes.Buffer
	buf.WriteString("Set{")
	flag := true
	for k := range self.m {
		if flag {
			flag = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", k))
	}
	buf.WriteString("}")

	return buf.String()
}

func (self *Set) IsSuperSet(other *Set) bool {
	self.Lock()
	defer self.Unlock()
	if other == nil {
		return false
	}
	selfLen := self.Len()
	otherLen := other.Len()
	if otherLen == 0 || selfLen == otherLen {
		return false
	}
	if selfLen > 0 && otherLen == 0 {
		return true
	}
	for v := range other.m {
		if !self.Contains(v) {
			return false
		}
	}
	return true
}

//属于A或属于B的元素
func (self *Set) Union(other *Set) *Set {
	self.Lock()
	defer self.Unlock()
	// if other == nil || other.Len() == 0{
	//    return self
	// }
	//
	// for v := range other.m{
	//    self.Add(v)
	// }
	// return self
	//不能改变集合A的范围
	union := NewSet()
	for v := range self.m {
		union.Add(v)
	}
	for v := range other.m {
		union.Add(v)
	}
	return union
}

//属于A且属于B的元素
func (self *Set) Intersect(other *Set) *Set {
	self.Lock()
	defer self.Unlock()
	if other == nil || other.Len() == 0 {
		return NewSet()
	}
	intsSet := NewSet()
	for v, _ := range other.m {
		if self.Contains(v) {
			intsSet.Add(v)
		}
	}
	return intsSet
}

//属于A且不属于B的元素
func (self *Set) Difference(other *Set) *Set {
	self.Lock()
	defer self.Unlock()
	diffSet := NewSet()
	if other == nil || other.Len() == 0 {
		diffSet.Union(self)
	} else {
		for v := range self.m {
			if !other.Contains(v) {
				diffSet.Add(v)
			}
		}
	}

	return diffSet
}

//集合A与集合B中所有不属于A∩B的元素的集合
func (self *Set) SymmetricDifference(other *Set) *Set {
	self.Lock()
	defer self.Unlock()
	//此时A∩B=∅，A中所有元素均不属于空集
	// if other == nil || other.Len() == 0{
	//    return self
	// }
	// ints := self.Intersect(other)
	// //此时A∩B=∅，A为空或B为空,B为空前面已经判断，此时B不能为空，即A为空
	// if ints == nil || ints.Len() == 0 {
	//    return other
	// }
	//
	// unionSet := self.Union(other)
	// result := New()
	// for v := range unionSet.m{
	//    if !ints.Contains(v){
	//       result.Add(v)
	//    }
	// }
	ints := self.Difference(other)
	union := self.Union(other)
	return union.Difference(ints)
}

func setTest() {
	set1 := NewSet()
	set1.Add(1)
	set1.Add("e2")
	set1.Add(3)
	set1.Add("e4")
	fmt.Println("set1:", set1)
	fmt.Printf("set1 Elements:%v\n", set1.Elements())

	//set2 := NewSet()
	//set2.Add(3)
	//set2.Add("e2")
	//set2.Add(5)
	//set2.Add("e6")

	//fmt.Println("set2:", set2)
	//fmt.Printf("set1 union set2:%v\n", set1.Union(set2))
	//fmt.Printf("set1 intersect set2:%v\n", set1.Intersect(set2))
	//fmt.Println(set1, set2)
	//fmt.Printf("set1 difference set2:%v\n", set1.Difference(set2))
	//fmt.Printf("set1 SymmetricDifference set2:%v\n", set1.SymmetricDifference(set2))
	//set1.Clear()
	fmt.Println(set1)
}
