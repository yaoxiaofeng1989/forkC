//一个数如果恰好等于它的因子之和,这个数就称为“完数”。例如 6=1+2
//+3.编程找出 1000 以内的所有完数。
//例如：
//6=1+2+3
//28=1+2+3+...+6+7
//496=1+2+3+...+30+31
//8128=1+2+3…+126+127
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 1000; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if i == sum {
			fmt.Println(i)
		}
	}
}
