package search

func BinarySearch(data []int , e int) int {
    start ,stop :=0, len(data)-1

    for start<=stop {
        mid:=(start+stop)/2

        if e<data[mid] {
            stop=mid-1
        } else if e>data[mid] {
            start=mid+1
        }else {
            return mid
        }
    }

    return -1
}

/**
递归写法
 */
func BinarySearchRecursion(data []int ,e int) int {
    return binarySearchRecursionLoop(data,e,0,len(data)-1)
}

func binarySearchRecursionLoop(data []int ,e int ,low int ,high int) int {
    if(low>high){
        return -1
    }
    mid:=(low+high)/2
    if e<data[mid] {
        return binarySearchRecursionLoop(data,e,low,mid-1)
    } else if e>data[mid] {
        high=mid+1
        return binarySearchRecursionLoop(data,e,mid+1,high)
    }else {
        return mid
    }
}

