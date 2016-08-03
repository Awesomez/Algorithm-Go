package main

import (
	"fmt"
	"./sort"
	"./tool"
	"./stack"
	"./search"
)

func main() {

	/*
	array:=[2]string{"a","b"}

	sl:=array[1:]

	fmt.Println(array)

	array[1]="c"


	fmt.Println(array)

	fmt.Println(sl)
	*/

    //test_stack()
    test_binarySearch();
}

func test_binarySearch(){
    array := []int{1, 3, 5, 7, 9, 11, 13, 15}
    e:=12;
    a:=search.BinarySearch(array,e)
    fmt.Println(a)

    a=search.BinarySearchRecursion(array,e)
    fmt.Println(a)
}

func test_stack(){
	stack := &stack.Stack{}
	stack.Push("iOS")
	stack.Push("Android")
	stack.Push("WindowsPhone")
	stack.Push("Symbian")
	stack.Push("BalckBerry")
	for stack.Size() > 0 {
        fmt.Println(stack.Pop())
	}
}

func test_sort(){
	data :=tool.MakeTestData(10);

	fmt.Println("init data :")
	fmt.Println(data)
	tool.CheckIsSorted(data)
	fmt.Println("")


	//sort.SelectionSort(data)
	//sort.BubbleSort(data)
	//sort.InsertionSort(data)
	sort.QuickSort(data)

	fmt.Println("sorted data :")
	fmt.Println(data)
	tool.CheckIsSorted(data)

}
