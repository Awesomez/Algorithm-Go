package tool

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func swap(data []int,a int, b int){
	data[a],data[b]=data[b],data[a];
}

func MakeTestData(n int) []int{
	file,err:=os.Open("./tool/sortData.txt");
	if err!=nil {
		fmt.Println(err);
		os.Exit(1)
	}
	defer file.Close();

	data:=make([]int,0,n);
	scanner:=bufio.NewScanner(file);

	i:=1;
	for scanner.Scan(){
		number,_:=strconv.Atoi(scanner.Text())
		data=append(data,number)
		if i>=n {
			break
		}
		i++
	}

	return data;
}

func CheckIsSorted(data []int){
	length:=len(data)

	fmt.Println("")
	fmt.Println("unsort: ")
	for i:=1;i<length;i++{
		if data[i-1]>data[i] {
			fmt.Print(" [")
			fmt.Print(i)
			fmt.Print("->")
			fmt.Print(data[i])
			fmt.Print("] ")
		}
	}
	fmt.Println("")
}
