package sort

import (
	"testing"
	"fmt"
	"../utils"
)

func Test_bubble_sort(t *testing.T){
	data :=utils.MakeTestData(20);

	fmt.Println("Initial array:")
	fmt.Println(data)
	sort(data);
	fmt.Println("Sorted array:")
	fmt.Println(data)
}
