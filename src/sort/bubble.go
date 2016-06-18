package sort

func sort(data []int){
	length := len(data);
	for i:=0;i<length;i++ {
		for j:=0;j<length-i-1;j++{
			if data[j+1]<data[j]{
				data[j+1],data[j]=data[j],data[j+1];
			}
		}
	}
}
