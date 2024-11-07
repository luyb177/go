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
// package main

// func main() {

//		//fmt.Printf("%09d", 5) //输出9位 空余的补0
//		// s := ""
//		// for ; s != "aaaaa"; {
//		// fmt.Println("Value of s:", s)
//		// s = s + "a"
//		// }
//		// s := "G"
//		// for ; s != "GGGGGGGG";{
//		// 	fmt.Println(s)
//		// 	s += "G"
//		// }
//	}
//
// 利用闭包完成计数器
// package main

// import "fmt"

// func main() {
// 	count := 0
// 	increase := func() (x int) {
// 		count++
// 		return count
// 	}
// 	decrease := func() (x int) {
// 		count--
// 		return count
// 	}
// 	fmt.Println(increase())
// 	fmt.Println(increase())
// 	fmt.Println(decrease())
// 	fmt.Println(decrease())

// }
// package main

// import "fmt"

// func main(){
// 	a := [...]string{"a", "b", "c", "d"}
// 	for i := range a {
// 	fmt.Println("Array item", i, "is", a[i])
// }
// }
// 数组实现斐波那契
// func main() {
// 	arr := [12]int{0}
// 	arr[1] = 1
// 	arr[2] = 2
// 	for i := 3; i < len(arr); i++ {
// 		arr[i] = arr[i-1] + arr[i-2]
// 		fmt.Printf("%d ", arr[i])
// 	}

// }
// func feibo(n int) int {
// 	if n < 2 {
// 		return 1
// 	}else{
// 		return feibo(n - 1)+feibo(n - 2) //递归
// 	}
// }
// func shu(x int)([]int){
// 	if x < 2{
// 		return []int{1}
// 	}else{
// 		return []int{feibo(x)}
// 	}

// }
//
//	func main() {
//		fmt.Println(shu(10))
//	}
// package main

// import "fmt"

// func main() {
// 	count := 0
// 	s := "hello, 世界"
// 	len := len(s)
// 	fmt.Println(len)
// 	for i := range s {
// 		fmt.Printf("%d ",i)
// 		count++
// 	}
// 	fmt.Println()
// 	fmt.Println(count)
// }
// package main
// import "fmt"
// type CustomStruct struct {
//     Name    string
//     Values  []int // 切片是不可比较的
// }

// func (c CustomStruct) Equals(other CustomStruct) bool {
//     if c.Name != other.Name {
//         return false
//     }
//     if len(c.Values) != len(other.Values) {
//         return false
//     }
//     for i, v := range c.Values {
//         if v != other.Values[i] {
//             return false
//         }
//     }
//     return true
// }

// func main() {
//     cs1 := CustomStruct{Name: "Test", Values: []int{1, 2, 3}}
//     cs2 := CustomStruct{Name: "Test", Values: []int{1, 2, 3}}
//     cs3 := CustomStruct{Name: "Test", Values: []int{4, 5, 6}}

//	    fmt.Println(cs1.Equals(cs2)) // true
//	    fmt.Println(cs1.Equals(cs3)) // false
//	}
// package main

// import "fmt"

// type student struct {
// 	name  string
// 	grade int
// 	class int
// }

//	func main() {
//		p1 := student{name: "xiaowang", grade: 2, class: 1}
//		p2 := student{name: "xiaowang", grade: 2, class: 1}
//		p3 := student{name: "xiaowang", grade: 2, class: 2}
//		if p1 == p2 {
//			fmt.Printf("p1 = p2\n")
//		} else {
//			fmt.Printf("p1 != p2\n")
//		}
//		if p1 == p3 {
//			fmt.Printf("p1 == p3\n")
//		} else {
//			fmt.Printf("p1 != p3\n")
//		}
//	}
// package main

// import "fmt"

// type student struct {
// 	name  string
// 	score []int
// }

// func (c student) Compare(o student) bool {
// 	if c.name != o.name { //string直接比较
// 		return false
// 	}
// 	if len(c.score) != len(o.score) { //比较切片的长度
// 		return false
// 	}
// 	for i, v := range c.score { //遍历 比较切片的值
// 		if v != o.score[i] {
// 			return false
// 		}
// 	}
// 	return true //都正确则返回true
// }
// func main() {
// 	p1 := student{name: "xiaowang", score: []int{90, 89, 90}}
// 	p2 := student{name: "xiaowang", score: []int{90, 89, 90}}
// 	p3 := student{name: "xiaowang", score: []int{99, 100, 99}}
// 	fmt.Printf("p1 = p2是%v\n", p1.Compare(p2))
// 	fmt.Printf("p1 = p3是%v\n", p1.Compare(p3))

// }
package main

// import "fmt"

// func main() {
// 	x := "hello!"
// 	for i := 0; i < len(x); i++ {
// 		x := x[i]
// 		if x != '!' {
// 			x := x + 'A' - 'a'
// 			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
// 		}
// 	}
// }
