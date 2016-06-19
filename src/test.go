package main

import (
	"fmt"
	"./sort"
	"./tool"
)

func main() {
	test_sort()
}

func test_sort(){
	data :=tool.MakeTestData(200);

	fmt.Println("init data :")
	fmt.Println(data)
	tool.CheckIsSorted(data)
	fmt.Println("")


	//sort.SelectionSort(data)
	sort.BubbleSort(data)

	fmt.Println("sorted data :")
	fmt.Println(data)
	tool.CheckIsSorted(data)

}
