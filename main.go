package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"Colorado","Utah","Wisconsin","Oregon"}
	sort.Strings(str)
	for _, s := range str {
		fmt.Println(len(s)%3)
	}
	fmt.Println(RemainderSorting(str))

}

func RemainderSorting(strArr []string) []string {

	sort.SliceStable(strArr, func(i, j int) bool {
		return len(strArr[i])%3 < len(strArr[j])%3
	})
	return strArr
}
	
