package main

import (
	"fmt"
)

/*声明全局变量*/
var g int

// type Books struct {
// 	bookName string
// 	author   string
// 	subject  string
// }

func Add(x, y int) int {
	return x + y
}

//交换
func swap(x, y string) (string, string) {
	return y, x
}

//打印切片
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
func Fi(k int) int {
	if k < 2 {
		return k
	}
	return Fi(k-1) + Fi(k-2)

}

const (
	a = iota
	b
	c
)

type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (nk *NokiaPhone) Error() string {
	return "test"
}

func (nk NokiaPhone) call() {
	fmt.Println("nokia phone is calling...")
}

func main() {
	var phone Phone = new(NokiaPhone)
	phone.call()

	channel := make(chan int)
	s := []int{2, 3, 4, 5}
	go getSum(s, channel)
	value := <-channel //从通道读取，默认无缓冲区
	fmt.Println(value)
	divide(-1)
	r, err := hello(-1)
	if err != nil {
		fmt.Println(r)
		fmt.Println(err)
	}
	printYIN("yinpan")

	fmt.Println(a, b, c)
	var ret = Add(2, 3)
	fmt.Printf("相加之和: %f\n", float32(ret)) //类型转换

	var m, n int
	var j int
	var ip *int
	if ip == nil {
		fmt.Printf("ip是空指针\n")
	}
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var salary [12]float32

	salary[2] = 10

	for j = 0; j < 5; j++ {
		fmt.Printf("元素%d = %f\n", j, balance[j])
	}
	m = 10
	n = 20
	g = m + n
	ip = &m

	fmt.Printf("m,n,g值分别为:%d + %d = %d\n", m, n, g)
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
	fmt.Printf("m 地址 :%x\n", &m)
	fmt.Printf("m 值 :%d\n", *ip)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum: ", sum)
	var nums = make([]int, 3, 5)
	printSlice(nums)
	nums = append(nums, 11)
	printSlice(nums)

	kvs := map[string]string{"a": "xyz", "b": "qwe"}
	for k := range kvs {
		fmt.Printf("%s ---> %s\n", k, kvs[k])
	}
	print(kvs["a"])
	for i, c := range "yinpan" {
		fmt.Println(i, c)
	}

	var countryCapitalMap = make(map[string]string)
	countryCapitalMap["china"] = "beijing"
	capital, ok := countryCapitalMap["America"]
	fmt.Println(capital)
	if ok {
		fmt.Println("中国首都是", capital)
	} else {
		fmt.Println("不存在")
	}
}
