package main

import "fmt"

type t1 interface {
	Add()
	Sub()
}

type s1 struct{}

func (s *s1) Add() {
	fmt.Printf("Add s1")
}

func (s *s1) Sub() {
	fmt.Printf("Sub s1")
}

type t2 interface {
	Add()
	Sub()
}

type s2 struct{}

func (s *s2) Add() {
	fmt.Printf("Add s2")
}

func (s *s2) Sub() {
	fmt.Printf("Sub s2")
}

func main() {

	sample1:=new(s1)
	
	if v,ok:=interface.(*s1);ok{
		fmt.Printf("t1 interface",v)
	}

	if v,ok:=interface.(*s2);ok{
		fmt.Printf("t2 interface",v)
	}
}
