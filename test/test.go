/*
 * @Descripttion:
 * @version:
 * @Author: Winston
 * @Date: 2020-04-16 17:43:08
 * @LastEditors: Winston
 * @LastEditTime: 2020-04-16 21:16:34
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
	for _, v = range x {

	}
}

func main() {
	multiplicationTable()
	multiplicationTableReverse()
}
