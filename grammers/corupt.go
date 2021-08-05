/*
* 一个目录下的同级文件归属一个包。
* 包名可以与其目录不同名。
* 包名为 main 的包为应用程序的入口包，编译源码没有 main 包时，将无法编译输出可执行的文件
 */

//go 关键字 多线程 并发
//defer 用于资源的释放 函数返回之前调用
package main

import (
	"errors"
	"fmt"
)

func hello(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("math noll")
	} else {
		return i, nil
	}
}

func printYIN(a string) string {
	fmt.Println("111:" + a)
	return a
}

func divide(j int) (int, string) {
	if j == -1 {
		dData := new(NokiaPhone)
		errormsg := dData.Error()
		fmt.Println(errormsg)
		return j, errormsg

	}
	return j, ""
}
func getSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
