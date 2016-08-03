package search

func binarySearch(data []int , e int) int {
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
