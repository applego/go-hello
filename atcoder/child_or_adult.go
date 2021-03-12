package main

import "fmt"

func main() {

	複数入力()

	違う型()

	var n int
	fmt.Scan(&n)

	if n < 20 {
		fmt.Println("Child")
		return
	}

	fmt.Println("Adult")
}

func 複数入力(){
	var x,n int
	fmt.Scan(&x,&n)
	fmt.Println(x)
	fmt.Println(n)
}

func 違う型(){
	var(
		n int
		s string
	)
	fmt.Scan(&n,&s)
	fmt.Println(n,s)
}
