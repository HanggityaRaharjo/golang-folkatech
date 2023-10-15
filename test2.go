package main

import "fmt"

func shiftArray(arrayData []int, i int) []int {
	if i <= 0 {
		return arrayData
	} else {
		temporaryArray := make([]int, 0, len(arrayData))

		for j := 0; j < i; j++ {
			temporaryArray = append(temporaryArray, arrayData[3])
			temporaryArray = append(temporaryArray, arrayData[0])
			temporaryArray = append(temporaryArray, arrayData[1])
			temporaryArray = append(temporaryArray, arrayData[6])
			temporaryArray = append(temporaryArray, arrayData[4])
			temporaryArray = append(temporaryArray, arrayData[2])
			temporaryArray = append(temporaryArray, arrayData[7])
			temporaryArray = append(temporaryArray, arrayData[8])
			temporaryArray = append(temporaryArray, arrayData[5])
		}

		return shiftArray(temporaryArray, i-1)
	}
}

func main() {
	originalArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 1 // Ganti nilai i sesuai kebutuhan Anda.
	shiftedArray := shiftArray(originalArray, i)
	fmt.Println(shiftedArray)
}
