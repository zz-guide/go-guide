package main

import "log"

func main() {

	p := Person{Name: "许磊"}
	log.Printf("地址1=%p\n", &p)
	p.setName("张三")
	p.setName1("张三")
	log.Println(p)
}

type Person struct {
	Name string
}

func (p Person) setName(name string) {
	log.Printf("地址2=%p\n", &p)
	(&p).Name = name
}

func (p *Person) setName1(name string) {
	log.Printf("地址3=%p\n", p)
	log.Printf("地址4=%p\n", &p)
	p.Name = name
}
