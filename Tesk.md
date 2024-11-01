# 1. 利用闭包实现一个计数器

```go
package main
import "fmt"
func main() {
    count := 0
    increase := func() (x int) {
        count++
        return count
    }
    decrease := func() (x int) {
        count--
        return count
    }
    fmt.Println(increase())
    fmt.Println(increase())
    fmt.Println(decrease())
    fmt.Println(decrease())
}
```
这是一种较为简单的计数器

---
# 2. map和slice内存扩容的区别

- slice 是一个长度可变的数组
	- 切片有最大长度 - 整个相关数组
	- 切片的长度永远不会超过它的容量，即 0 <= len(s) <= cap(s)
	- 而如果想增加切片的容量，必须创建一个新的更大的切片并把原分片的内容都拷贝过来
- map 是可以动态增长的
	- 不存在固定长度或者最大限制
	- 当map 增长到容量上限，继续增加key-value对，map的大小会自动加1

# 3.比较字符串"hello，世界"的**长度** 和for range 该字符串的循环次数

```go
package main
import "fmt"
func main() {
count := 0
    s := "hello, 世界"
    len := len(s)
    fmt.Println(len) // 13
    for i := range s {
        fmt.Printf("%d ",i) // 0 1 2 3 4 5 6 7 10
        count++
    }
    fmt.Println()
    fmt.Println(count) // 9
}
```
- s 的长度是13，即在ASCII表上的字符占一个字节，而非ASCII编码的字符占2~4个字符
- for range 循环次数为9 表明s有9个字符，而且索引也只会出现第一个，具体形式如下：
- ![[Pasted image 20241101211903.png]]
# 4.如何进行结构体之间的比较？
- 对于简单的结构体（不包含切片、映射、函数等）可以直接使用`==` 比较
```go
package main
import "fmt"
type student struct {
    name  string
    grade int
    class int
}
func main() {
    p1 := student{name: "xiaowang", grade: 2, class: 1}
    p2 := student{name: "xiaowang", grade: 2, class: 1}
    p3 := student{name: "xiaowang", grade: 2, class: 2}
    if p1 == p2 {
        fmt.Printf("p1 = p2\n")
    } else {
        fmt.Printf("p1 != p2\n")
    }
    if p1 == p3 {
        fmt.Printf("p1 == p3\n")
    } else {
        fmt.Printf("p1 != p3\n")
    }
}
```
- 输出结果就是`p1 = p2` 和`p1 != p3` ，因此简单的结构体可以使用`==`直接判断
---
- 对于含有切片、映射、函数等的结构体，可以自定义比较方法
```go
package main
import "fmt"
type student struct {
    name  string
    score []int
}
func (c student) Compare(o student) bool { //定义方法
    if c.name != o.name { //string直接比较
        return false
    }
    if len(c.score) != len(o.score) { //比较切片的长度
        return false
    }
    for i, v := range c.score { //遍历 比较切片的值
        if v != o.score[i] {
            return false
        }
    }
    return true //都正确则返回true
}
func main() {
    p1 := student{name: "xiaowang", score: []int{90, 89, 90}}
    p2 := student{name: "xiaowang", score: []int{90, 89, 90}}
    p3 := student{name: "xiaowang", score: []int{99, 100, 99}}
    fmt.Printf("p1 = p2是%v\n", p1.Compare(p2))
    fmt.Printf("p1 = p3是%v\n", p1.Compare(p3))  
}
```

# 5.以下哪里x进行重新声明，为什么

```go

  func main() {
      x := "hello!"
      for i := 0; i < len(x); i++ {
          x := x[i]
          if x != '!' {
              x := x + 'A' - 'a'
              fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
         }
     }
  }

```
- `x := x[i]` 实现了对x的重新声明，将x 原来的`var x string` 变为`var x byte` 
- 因为`x[i]` 是一个字符，此时相当于`x := (byte)` 因此x被重新声明和初始化