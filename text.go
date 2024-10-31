// // hello world
// // package main
// // import fm "fmt" // 重命名
// // func main(){
// // 	fm.Println("hello world")
// // }

// //函数
// //斐波那契
// //feibo 是用来递归求斐波那契数列
// package main
// import "fmt"
// func feibo(n int) int {
// 	if n < 2 {
// 		return 1
// 	}else{
// 		return feibo(n - 1)+feibo(n - 2) //递归
// 	}
// }
// func main(){
// 	// //const pi  = 3.1415926 //省略float64
// 	// const pi float64 = 3.1415926
// 	// const(
// 	// 	a = iota //枚举
// 	// 	b
// 	// 	c
// 	// )
// 	// fmt.Println(feibo(a))
// 	// fmt.Println(feibo(b))
// 	// fmt.Println(feibo(c))
// 	// i := 0
// 	// p := &i
// 	// fmt.Println(i)
// 	// fmt.Println(p)
// 	i := 1
// 	p := 2
// 	i, p = p, i //交换两个数
// 	fmt.Println(i,p)

// }
//练习
// package main

// var a = "G"

// func main() {
//    n()	// G
//    m()	// O
//    n()	// G
// }

// func n() { print(a) }

// func m() {
//    a := "O"
//    print(a) //局部变量的作用域只有函数内部
// }

// package main

// var a string

// func main() {
//    a = "G"
//    print(a) // G
//    f1()		//O G
// }

// func f1() {
//    a := "O"
//    print(a)
//    f2() //
// }

//	func f2() {
//	   print(a) // 未在局部定义a 则使用main中的a
//	}
package main

import "fmt"

func main() {
	fmt.Printf("%09d", 5) //输出9位 空余的补0
}
