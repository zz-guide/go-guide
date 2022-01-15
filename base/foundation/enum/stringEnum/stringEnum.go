package main

import (
	"fmt"
)

type PolicyType int32

const (
	PolicyMin PolicyType = 0
	PolicyMax PolicyType = 1
	PolicyMid PolicyType = 2
	PolicyAvg PolicyType = 3
)

func (p PolicyType) String() string {
	switch p {
	case PolicyMin:
		return "MIN"
	case PolicyMax:
		return "MAX"
	case PolicyMid:
		return "MID"
	case PolicyAvg:
		return "AVG"
	default:
		return "UNKNOWN"
	}
}

func foo(p PolicyType) {
	fmt.Printf("enum value: %v\n", p)
}

func main() {
	foo(PolicyMax)
}
