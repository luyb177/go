// // package main

// // import (
// // 	"fmt"
// // 	"time"
// // )

// // func main() {
// // 	t := time.Now()
// // 	fmt.Println(t.Year())
// // 	fmt.Println(t.Month())
// // 	fmt.Println(t.Day())
// // 	fmt.Println(t.Hour())
// // 	fmt.Println(t.Second())
// // 	fmt.Println(t.Nanosecond())
// // 	fmt.Println(t.Date())
// // 	fmt.Println(t)

// // }
// package main

// import "fmt"

// // func main() {
// // 	count := 0
// // 	s := "abc"

// // 	s += "de"
// // 	c := []byte{'a', 'b', 'c'}
// // 	c = append(c, 'd')
// // 	c = append(c, 'e')
// // 	for range c {
// // 		count++
// // 	}
// // 	fmt.Println(s)
// // 	fmt.Println(c)
// // 	fmt.Println(string(c))

// // 	fmt.Println(count)

// // }

// // func main() {

// // 	m := map[string]int{
// // 		"xiaowang": 90,
// // 		"xiaolu":   100,
// // 	}
// // 	for range m {
// // 		fmt.Println(m) //顺序打印
// // 	}
// // 	for i, v := range m {
// // 		fmt.Printf("%s的语文成绩是:%d ", i, v) //乱序遍历
// // 	}
// // }
// // type user struct{
// // 	name string
// // 	password string
// // }
// // func checkpassword(u *user,password string)bool{
// // 	return u.password == password
// // }
// // func main(){
// // 	u := user{"xiaowang","13222"}
// // 	password := "13222"
// // 	fmt.Println(checkpassword(&u,password))
// // }
// func main() {
// 	//var s string
// 	//fmt.Printf("xxx%sxxx", s)

// }
package main

import "fmt"

// func main() {
// 	c := 'A'
// 	fmt.Printf("%T %s", c, c)
// }
// func main(){
// 	fmt.Printf(`\n`)
// 	fmt.Printf(`fakshj
// 	fafa
// 	faasfa
// 	fafa
// 	ffaas
// 	fafaf
// 	faf`)
// 	c := `fasfa
// 	fafaf
// 	fafaf
// 	fafaf
// 	fafaf`
// 	fmt.Printf("%s",c)
// }
// func main(){
// 	var s int = 94
// 	switch s {
// 	case 90,99,94:
// 		fmt.Println("A")
// 	}
// }
//打印菱形
// func main() {
// 	n := 0
// 	fmt.Scan(&n)
// 	//正三角形
// 	for i := 0; i <= (n-1)/2; i++ { //分为上半部分
// 		//n = 11 i = 0 j = 5 打印5个空格
// 		//		 i = 1 j = 4 打印4个空格
// 		//实现空格的自减
// 		for j := (n - 1) / 2 - i; j > 0; j-- {
// 			fmt.Printf("  ")
// 		}
// 		//n = 11 i = 0 j = 0 打印1个空格
// 		//       i = 1 j = 0 打印3个空格
// 		for j := 0; j < 2 * i + 1; j++ {
// 			fmt.Printf("* ")
// 		}
// 		fmt.Println()
// 	}
// 	//倒三角形
// 	//倒序
// 	for i := (n - 1) / 2; i > 0 ; i--{
// 		//n = 11 i = 5 j = 0 打印1个空格
// 		//       i = 4 j = 0 打印2个空格
// 		for j := 0; j <=(n - 1) /2 - i; j++ {
// 			fmt.Printf("  ")
// 		}
// 		//
// 		for j := 2 * i -1 ; j > 0 ; j --{
// 			fmt.Printf("* ")
// 		}
// 		fmt.Println()
// 	}
// }
//阶乘
// func f1(x int) int {

// 	if x == 1 || x == 0 {
// 		return 1
// 	} else {
// 		return f1(x-1) * x
// 	}

// }
// func main() {
// 	fmt.Println(f1(6))
// }
//回调函数->一个被当做参数的函数
// func add(x, y int, f1 func(int) int) int {
// 	return x + y + f1(x)
// }
// func f1(x int) int {
// 	return x+11
// }
// func main() {
// 	c := add(1, 2, f1)
// 	fmt.Println(c)
// }

//闭包 -> 外层函数，和内层函数，内层函数使用外层函数的变量，然后这个变量的作用域升级了
// func main(){
//     r1 := increament()
//     fmt.Println(r1)
//     fmt.Println(r1())
//     fmt.Println(r1())
//     fmt.Println(r1())
//     fmt.Println(r1())
//     fmt.Println(r1())
//     fmt.Println(r1())
//     r2 := increament()
//     fmt.Println(r2)
//     fmt.Println(r2())
//     fmt.Println(r1())
//     fmt.Println(r2())
// }
// func increament()func()int{
//     i := 0
//     fun := func ()int{
//          i++
//          return i
//     }
//     return fun
// }

func main() {
	s := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := s[:0]
	fmt.Println(s1)
	s2 := s[:5]
	fmt.Println("s2", "len", len(s2), "cap", cap(s2))
	s3 := s[1:5]
	s4 := s[2:5]
	s5 := s[3:5]
	fmt.Println("s3", "len", len(s3), "cap", cap(s3))
	fmt.Println("s4", "len", len(s4), "cap", cap(s4))
	fmt.Println("s5", "len", len(s5), "cap", cap(s5))
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s[1])//开辟的新数组的地址与s[i]相同

	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", s4)
	fmt.Printf("%p\n", s5)

}
