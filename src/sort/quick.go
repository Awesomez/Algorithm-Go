package sort

func QuickSort(data []int){
	quick(data,0,len(data)-1)
}

func quick(data []int,low int,hight int){
	if low<hight {
		mid:=partition(data,low,hight)
		quick(data,low,mid-1)
		quick(data,mid+1,hight)
	}
}

func partition(data []int, low int, hight int) int {
	flag:=data[low]
	for low<hight {
		for low<hight && flag<=data[hight] {
			hight--
		}
		data[low]=data[hight]

		for low<hight && data[low]<=flag {
			low++
		}
		data[hight]=data[low]
	}
	data[low]=flag
	return low
}
