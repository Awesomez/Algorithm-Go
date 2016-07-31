package sort

func QuickSort(data []int){
	quick(data,0,len(data)-1)
}

func quick(data []int,low int,high int){
	if low<high {
		mid:=partition(data,low,high)
		quick(data,low,mid-1)
		quick(data,mid+1,high)
	}
}

func partition(data []int, low int, high int) int {
	flag:=data[low]
	for low<high {
		for low<high && flag<=data[high] {
			high--
		}
		data[low]=data[high]

		for low<high && data[low]<=flag {
			low++
		}
		data[high]=data[low]
	}
	data[low]=flag
	return low
}
