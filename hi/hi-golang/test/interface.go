package main

import "fmt"

type A interface {
	Name() string
	Cal(params ...int) int
}

type A1 struct {
}

func (a1 *A1) Name() string {
	return "Add"
}

func (a1 *A1) Cal(params ...int) int {
	all := 0
	for _, n := range params {
		all += n
	}
	return all
}

type A2 struct {
}

func (a2 *A2) Name() string {
	return "Reduce"
}

func (a2 *A2) Cal(params ...int) int {
	all := 0
	for _, n := range params {
		all -= n
	}
	return all
}

func GetA(vtype string) A {
	var a A = nil
	switch vtype {
	case "A1":
		//a = A1{}
		// ERROR: cannot use A1 literal (type A1) as type A in assignment:
		a = new(A1)
	case "A2":
		a = new(A2)
	}
	return a
}

func main() {
	param := []int{100, 10, 1, 0}
	var a1, a2 A
	a1 = GetA("A1")
	a2 = GetA("A2")
	fmt.Println(a1.Name(), a1.Cal(param...))
	fmt.Println(a2.Name(), a2.Cal(param...))
}
