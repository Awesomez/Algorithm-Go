package sort

/**
从右边选择最小的
 */

func SelectionSort(data []int){
	length:=len(data)

	for i:=0;i<length;i++ {
		min:=i
		for j:=i;j<length;j++{
			if data[j]<data[min] {
				min=j
			}
		}
		data[i],data[min]=data[min],data[i]
	}
}
