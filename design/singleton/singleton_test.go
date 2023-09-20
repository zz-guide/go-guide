package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	fmt.Println(ins1)
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("实例不相等")
	}

	t.Log("是同一个实例")
}
