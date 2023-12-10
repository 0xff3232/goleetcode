package removeDuplicatesFromSortedArray

func RemoveDups(arr []int) int { 
    if len(arr) == 0 {
        return 0
    }

    i := 0
    for j := 1; j < len(arr); j++ {
        if arr[j] != arr[i] {
            i++
            arr[i] = arr[j]
        }
    }

    return i + 1
}