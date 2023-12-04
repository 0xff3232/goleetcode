package leettest 

import (
	"fmt"
)

func main() {
	test := LargestGoodInteger("1233312223")
	fmt.Println(test)
}

func LargestGoodInteger(n string) string {
	N := len(n) - 1
	max_n := ""
	count := 1

	for i := 0; i < N; i++ {
		if n[i] == n[i+1] {
			count += 1
			if count == 3 {
				curr_n := string(n[i])
				if curr_n > max_n {
					max_n = curr_n
				}
				count = 1
			}
		} else {
			count = 1
		}
	}
    
	if max_n != "" {
		return max_n + max_n + max_n
	}
	return ""

}
