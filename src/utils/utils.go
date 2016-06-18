package utils

import (
	"bufio"
	"go/build"
	"os"
	"path/filepath"
	"strconv"
	"fmt"
)

func MakeTestData(n int) []int{
	file,err:=os.Open("../utils/sortData.txt");
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

func GetArrayOfSize(n int) []int {
	p, err := build.Default.Import("github.com/arnauddri/algorithms/algorithms/sorting/utils", "", build.FindOnly)

	if err != nil {
		// handle error
	}
	fname := filepath.Join(p.Dir, "IntegerArray.txt")
	f, _ := os.Open(fname)
	defer f.Close()

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, s)
	}

	return numbers[0:n]
}
