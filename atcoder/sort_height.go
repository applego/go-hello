package main

import (
	"fmt"
	"sort"
)

func main() {
	var n,x int
	fmt.Scan(&n,&x)
	t := make([]int,n)
	for i:= 0;i<n;i++{
		fmt.Scan(&t[i])
	}
	sort.Ints(t)
	for i:=0;i<x;i++{
		fmt.Println(t[i])
	}
}

/*
//* Not good
var t []int
for i := 0; i < n; i++ {
    var tmp int
    fmt.Scan(&tmp)
    t = append(t, tmp)
}

//* Good
t := make([]int, n)
for i := 0; i < n; i++ {
    fmt.Scan(&t[i])
}
*/
