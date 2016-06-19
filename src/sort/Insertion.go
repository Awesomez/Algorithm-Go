package sort

/**
插入排序,左边是已排序的,往左插序
 */

func InsertionSort(data []int){
	length:=len(data)

	for i:=1;i<length;i++{
		for j:=i;j>0;j--{
			if data[j]<data[j-1] {
				data[j],data[j-1]=data[j-1],data[j]
			}
		}
	}
}
