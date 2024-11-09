# 第一题
思考以下程序在**并发中**出现panic的可能是什么？并给出解决方案
```go
package main

type message struct {
    Topic     string
    Partition int32
    Offset    int64
}

type FeedEventDM struct {
    Type    string
    UserID  int
    Title   string
    Content string
}

type MSG struct {
    ms        message
    feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
    var consumeMSG []MSG
    var lastConsumeTime time.Time // 记录上次消费的时间
    msgs := make(chan MSG)

    //这里源源不断的生产信息
    go func() {
       for i := 0; ; i++ {
          msgs <- MSG{
             ms: message{
                Topic:     "消费主题",
                Partition: 0,
                Offset:    0,
             },
             feedEvent: FeedEventDM{
                Type:    "grade",
                UserID:  i,
                Title:   "成绩提醒",
                Content: "您的成绩是xxx",
             },
          }
          //每次发送信息会停止0.01秒以模拟真实的场景
          time.Sleep(100 * time.Millisecond)
       }
    }()

    //不断接受消息进行消费
    for msg := range msgs {
       // 添加新的值到events中
       consumeMSG = append(consumeMSG, msg)
       // 如果数量达到额定值就批量消费
       if len(consumeMSG) >= ConsumeNum {
          //进行异步消费
          go func() {
             m := consumeMSG[:ConsumeNum]
             fn(m)
          }()
          // 更新上次消费时间
          lastConsumeTime = time.Now()
          // 清除插入的数据
          consumeMSG = consumeMSG[ConsumeNum:]
       } else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
          // 如果距离上次消费已经超过5分钟且有未处理的消息
          if len(consumeMSG) > 0 {
             //进行异步消费 
             go func() {
                m := consumeMSG[:ConsumeNum] //可能会越界
                fn(m)
             }()
             // 更新上次消费时间
             lastConsumeTime = time.Now()
             // 清空插入的数据
             consumeMSG = consumeMSG[ConsumeNum:]
          }
       }
    }
}

func fn(m []MSG) {
    fmt.Printf("本次消费了%d条消息\n", len(m))
}
```

## 解题
1. 没有加包`import("fmt" "time")`
2. 没有报错小小的运行一下`panic: runtime error: slice bounds out of range [:5] with capacity 3` 显示访问越界 ->长度限制
3. 应该是临界资源的安全问题 ->锁
4. 仔细阅读代码，发现可能越界的切片
```go
package main

  

import (

    "fmt"

    "sync"

    "time"

)

  

type message struct {

    Topic     string

    Partition int32

    Offset    int64

}

  

type FeedEventDM struct {

    Type    string

    UserID  int

    Title   string

    Content string

}

  

type MSG struct {

    ms        message

    feedEvent FeedEventDM //结构体的嵌套

}

  

const ConsumeNum = 5

  

func main() {

   var consumeMSG []MSG
    var lastConsumeTime time.Time // 记录上次消费的时间
    msgs := make(chan MSG)        //创建一个chan类型是MSG 默认len为0
    var mutex sync.Mutex
    //这里源源不断的生产信息
    go func() {
        for i := 0; ; i++ {
            msgs <- MSG{ //不断地向chan输入
                ms: message{
                    Topic:     "消费主题",
                    Partition: 0,
                    Offset:    0,
                },
                feedEvent: FeedEventDM{
                    Type:    "grade",
                    UserID:  i,
                    Title:   "成绩提醒",
                    Content: "您的成绩是xxx",
                },
            }
            //每次发送信息会停止0.01秒以模拟真实的场景
            time.Sleep(100 * time.Millisecond)
        }
    }()
    //不断接受消息进行消费
    for msg := range msgs { //读取通道
        mutex.Lock() //拿到共享资源之前锁起来
        // 添加新的值到events中
        consumeMSG = append(consumeMSG, msg) //自动扩容
        // 如果数量达到额定值就批量消费
        if len(consumeMSG) >= ConsumeNum {
            //进行异步消费
            go func() {
                m := make([]MSG, ConsumeNum) //拷贝
                copy(m, consumeMSG)
                fn(m)
            }()
            // 更新上次消费时间
            lastConsumeTime = time.Now()
            // 清除插入的数据
            consumeMSG = consumeMSG[ConsumeNum:]
        } else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
            // 如果距离上次消费已经超过5分钟且有未处理的消息
            if len(consumeMSG) > 0 {
                //进行异步消费
                go func() {
                    m := make([]MSG, len(consumeMSG))
                    copy(m, consumeMSG)
                    fn(m)
                }()
                // 更新上次消费时间
                lastConsumeTime = time.Now()
                // 清空插入的数据
                consumeMSG = consumeMSG[ConsumeNum:]
            }
        }
        mutex.Unlock()
    }
}
func fn(m []MSG) {
    fmt.Printf("本次消费了%d条消息\n", len(m))
}
```
- 总结
	- 使用锁解决并发问题
	- 限制长度
	- 拷贝的目的是为了防止函数执行延迟后发生的数据不一致的情况
# 第二题
使用for循环生成20个goroutine，并向一个channel传入随机数和goroutine编号，等待这些goroutine都生成完后，想办法给这些goroutine按照编号进行排序(输出排序前和排序后的结果,要求不使用额外的空间存储着20个数据)
```go
package main
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)
type gro struct {
    random int
    number int
}  
func random(wg *sync.WaitGroup, ch1 chan gro) {
    a := rand.Intn(100)
    b := rand.Intn(100)
    ch1 <- gro{ //结构体的使用
        a,
        b,
    }
    wg.Done()
}
func main() {
    var wg sync.WaitGroup //设成全局变量就不需要传参
    ch := make(chan gro, 20)
    Num := make([]gro, 0)
    //ch1 := make(chan gro, 20)
    // ch2 := make(chan gro,20)
    rand.Seed(time.Now().UnixNano()) //在主程序中种下一颗种子
    wg.Add(20) //判断还有几个 20个
    for i := 0; i < 20; i++ { //生成20个
        go random(&wg, ch)
    }
    wg.Wait()
    close(ch) //关闭通道，避免接收方阻塞
    fmt.Println("------更改前--------")
    for i := range ch {
        Num = append(Num, i)
        fmt.Printf("%d %d\n", i.number, i.random)
    }
    for i := 0; i < 19; i++ {
        for j := 0; j < 19-i; j++ {
            if Num[j].number > Num[j+1].number {
                Num[j], Num[j+1] = Num[j+1], Num[j]
            }
        }
    }
    fmt.Println("------更改后--------")
    for i := range Num {
        fmt.Println(Num[i])
    }

    //更改 按照number的大小来排序

    // time.Sleep(time.Second * 3)

    // close(ch1)

    // go func() {

    //  for i := range ch1 {

    //      for j := range ch1 {

    //          if i.number > j.number {

    //              i, j = j, i

    //          }

    //      }

    //  }

    // }()

    // time.Sleep(time.Second * 3)

}
```
注： 因为没有想到如何不额外存储通道里面的数据，就用切片存储了一下
解题方式：
- 生成随机数`rand.Seed(time.Now().UnixNano()) //在主程序中种下一颗种子`
- 因为要传一个随机数和一个编码->chan的类型设成结构体存放两个数据
- 创建20个goroutine用random传送数据；使用同步等待组`sync.Waigroup`等待goroutine全部完成
- 传输完成后使用`close(ch)` 来关闭通道，避免接收方阻塞
	- 资料：关闭通道主要是用来告诉接收方，通道不再发送数据了，也就是说，接收方可以结束接收操作，不再等待数据。这是因为通道的接收方会阻塞，直到通道关闭或有数据可读。
- 接下里把通道里的数据append至Num中，打印排序前和排序后的数据
	- 资料：如果你在关闭通道时 **第一次全部接收**（即读取了通道中的所有数据），那么之后你就不能再继续从该通道读取数据了，因为通道已经被关闭并且所有数据都已被消费。尝试从一个空且已关闭的通道读取数据，会立即返回该通道类型的零值，并且不会阻塞。
- 切片扩容是在原有数据的前提下继续追加数据，因此创建切片的时候不需要提供长度
# 第三题
3. 经典老题：交叉打印下面两个字符串（要求一个打印完，另一个会继续打印） "ABCDEFGHIJKLMNOPQRSTUVWXYZ" "0123..." 得到："AB01CD23EF34..."
```go
package main

  

import (

    "fmt"

)

func main(){

    ch := make(chan int,2)

    ch1 := make(chan int,2)

    go func (){

        for{

            for  i := 65;i <= 90; i++{

            ch <- i

            date := <- ch1

            fmt.Print(date)

        }

    }

    }()

    for {

        for i := 0; i < 10; i++{

            ch1 <- i

            date := <- ch

            fmt.Printf("%c",date)

        }

    }

}
```
输出结果：

```
A01BC23DE45FG67HI89JK01LM23NO45PQ67RS89TU01VW23XY45ZA...
```
总结：
- 两个线程，使用通道来阻塞一会使得交叉打印