/*
 * @Descripttion:
 * @version:
 * @Author: Winston
 * @Date: 2020-04-16 17:43:08
 * @LastEditors: Winston
 * @LastEditTime: 2020-04-17 10:11:59
 */

package main

import "fmt"

func multiplicationTable() {
	for i := 1; i < 10; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d*%d=%2d ", k, i, i*k)
		}
		fmt.Printf("\n")
	}
}
func multiplicationTableReverse() {
	for i := 1; i < 10; i++ {
		for k := 9; k >= i; k-- {
			fmt.Printf("%d*%d=%2d ", i, k, i*k)
		}
		fmt.Printf("\n")
	}
}

func sumArray(x []int) {
	var v int
	sum := 0
	for _, v = range x {
		sum += v
	}
	fmt.Printf("%d", sum)
}

func findAddSum(x [5]int) (g [2]int) {
	resultMap := make(map[string]int, 8)
	var resultArry [2]int
	var v int
	var i int
	target := 8
	for i, v = range x {
		_, b := resultMap[string(v)]
		if b {
			resultArry[0] = resultMap[string(v)]
			resultArry[1] = i
			return resultArry
		}
		resultMap[string(target-v)] = i
	}
	return
}

func main() {
	// multiplicationTable()
	// multiplicationTableReverse()
	var x = [...]int{1, 3, 5, 7, 8}
	fmt.Println(findAddSum(x))
}
