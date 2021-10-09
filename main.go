package main

import "fmt"

func deleteProducts(ids []int32, m int32) int32 {
	// Write your code here
	mapIds := make(map[int32]int32)
	for _, val := range ids {
		if mapIds[val] >= 1 {
			mapIds[val] += 1
		} else {
			mapIds[val] = 1
		}
	}

	mapIdsR := make(map[int32]int32)
	for k, v := range mapIds {
		mapIdsR[v] = k
	}

	for m > 0 {
		for k, v := range mapIdsR {
			if m >= k {
				delete(mapIdsR, k)
				m = m - k
			} else if m < k {
				delete(mapIdsR, k)
				mapIdsR[k-m] = v
			}
		}
	}
	result := int32(len(mapIdsR))
	return result
}

func main() {
	tes1 := []int32{1, 2, 3, 1, 2, 2}
	var m int32 = 3
	fmt.Println(deleteProducts(tes1, m))
}
