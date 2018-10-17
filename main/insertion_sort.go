package main

import (
	"fmt"
	"github.com/spradeepv/golang/sort"
)

func main() {
	i := []int{14, 11, 7, 2, 4, 10, 50}
	fmt.Println(i)
	sort.InsertionSort(i)
	fmt.Println(i)
}
