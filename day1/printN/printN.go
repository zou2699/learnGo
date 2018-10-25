package printN

import "fmt"

// 循环
func PrintNa(n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%d\n",i)
	}
}

// 递归 空间占用
func PrintNb(n int) {
	if n>0 {
		PrintNb(n-1)
		fmt.Printf("%d\n",n)
	}
	return
}