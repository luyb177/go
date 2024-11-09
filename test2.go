// // package main

// // import (
// // 	"fmt"
// // 	"sync"
// // 	"time"
// // )

// // type message struct {
// // 	Topic     string
// // 	Partition int32
// // 	Offset    int64
// // }

// // type FeedEventDM struct {
// // 	Type    string
// // 	UserID  int
// // 	Title   string
// // 	Content string
// // }

// // type MSG struct {
// // 	ms        message
// // 	feedEvent FeedEventDM //结构体的嵌套
// // }

// // const ConsumeNum = 5

// // func main() {
// // 	var consumeMSG []MSG
// // 	var lastConsumeTime time.Time // 记录上次消费的时间
// // 	msgs := make(chan MSG)        //创建一个chan类型是MSG 默认len为0
// // 	var mutex sync.Mutex

// // 	//这里源源不断的生产信息
// // 	go func() {
// // 		for i := 0; ; i++ {
// // 			msgs <- MSG{ //不断地向chan输入
// // 				ms: message{
// // 					Topic:     "消费主题",
// // 					Partition: 0,
// // 					Offset:    0,
// // 				},
// // 				feedEvent: FeedEventDM{
// // 					Type:    "grade",
// // 					UserID:  i,
// // 					Title:   "成绩提醒",
// // 					Content: "您的成绩是xxx",
// // 				},
// // 			}
// // 			//每次发送信息会停止0.01秒以模拟真实的场景
// // 			time.Sleep(100 * time.Millisecond)
// // 		}
// // 	}()

// // 	//不断接受消息进行消费
// // 	for msg := range msgs { //读取通道
// // 		mutex.Lock() //拿到共享资源之前锁起来

// // 		// 添加新的值到events中
// // 		consumeMSG = append(consumeMSG, msg) //自动扩容
// // 		// 如果数量达到额定值就批量消费
// // 		if len(consumeMSG) >= ConsumeNum {
// // 			//进行异步消费
// // 			go func() {
// // 				m := make([]MSG, ConsumeNum) //拷贝
// // 				copy(m, consumeMSG)
// // 				fn(m)
// // 			}()
// // 			// 更新上次消费时间
// // 			lastConsumeTime = time.Now()
// // 			// 清除插入的数据
// // 			consumeMSG = consumeMSG[ConsumeNum:]
// // 		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
// // 			// 如果距离上次消费已经超过5分钟且有未处理的消息
// // 			if len(consumeMSG) > 0 {
// // 				//进行异步消费
// // 				go func() {
// // 					m := make([]MSG, len(consumeMSG))
// // 					copy(m, consumeMSG)
// // 					fn(m)
// // 				}()
// // 				// 更新上次消费时间
// // 				lastConsumeTime = time.Now()
// // 				// 清空插入的数据
// // 				consumeMSG = consumeMSG[ConsumeNum:]
// // 			}
// // 		}
// // 		mutex.Unlock()
// // 	}
// // }

// // func fn(m []MSG) {
// // 	fmt.Printf("本次消费了%d条消息\n", len(m))
// // }
// // // package main

// // // import (
// // // 	"fmt"
// // // 	"sync"
// // // 	"time"
// // // )

// // // type message struct {
// // // 	Topic     string
// // // 	Partition int32
// // // 	Offset    int64
// // // }

// // // type FeedEventDM struct {
// // // 	Type    string
// // // 	UserID  int
// // // 	Title   string
// // // 	Content string
// // // }

// // // type MSG struct {
// // // 	ms        message
// // // 	feedEvent FeedEventDM //结构体的嵌套
// // // }

// // // const ConsumeNum = 5

// // // func main() {
// // // 	var consumeMSG []MSG
// // // 	var lastConsumeTime time.Time // 记录上次消费的时间
// // // 	msgs := make(chan MSG)        //创建一个chan类型是MSG 默认len为0
// // // 	var mutex sync.Mutex

// // // 	//这里源源不断的生产信息
// // // 	go func() {
// // // 		for i := 0; ; i++ {
// // // 			msgs <- MSG{ //不断地向chan输入
// // // 				ms: message{
// // // 					Topic:     "消费主题",
// // // 					Partition: 0,
// // // 					Offset:    0,
// // // 				},
// // // 				feedEvent: FeedEventDM{
// // // 					Type:    "grade",
// // // 					UserID:  i,
// // // 					Title:   "成绩提醒",
// // // 					Content: "您的成绩是xxx",
// // // 				},
// // // 			}
// // // 			//每次发送信息会停止0.01秒以模拟真实的场景
// // // 			time.Sleep(100 * time.Millisecond)
// // // 		}
// // // 	}()

// // // 	//不断接受消息进行消费
// // // 	for msg := range msgs { //读取通道
// // // 		mutex.Lock() //拿到共享资源之前锁起来

// // // 		// 添加新的值到events中
// // // 		consumeMSG = append(consumeMSG, msg) //自动扩容
// // // 		// 如果数量达到额定值就批量消费
// // // 		if len(consumeMSG) >= ConsumeNum {
// // // 			//进行异步消费
// // // 			go func() {
// // // 				m := consumeMSG[:ConsumeNum]
// // // 				fn(m)
// // // 			}()
// // // 			// 更新上次消费时间
// // // 			lastConsumeTime = time.Now()
// // // 			// 清除插入的数据
// // // 			consumeMSG = consumeMSG[ConsumeNum:]
// // // 		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
// // // 			// 如果距离上次消费已经超过5分钟且有未处理的消息
// // // 			if len(consumeMSG) > 0 {
// // // 				//进行异步消费
// // // 				go func() {
// // // 					m := consumeMSG[:len(consumeMSG)]

// // // 					fn(m)
// // // 				}()
// // // 				// 更新上次消费时间
// // // 				lastConsumeTime = time.Now()
// // // 				// 清空插入的数据
// // // 				consumeMSG = consumeMSG[ConsumeNum:]
// // // 			}
// // // 		}
// // // 		mutex.Unlock()
// // // 	}
// // // }

// // // func fn(m []MSG) {
// // // 	fmt.Printf("本次消费了%d条消息\n", len(m))
// // // }

// // package main

// // import (
// // 	"fmt"
// // )

// // func main() {
// // 	ch := make(chan int, 2)
// // 	ch1 := make(chan int, 2)
// // 	go func() {
// // 		for {
// // 			for i := 65; i <=90; i++ {
// // 				ch <- i
// // 				date := <-ch1
// // 				fmt.Print(date)
// // 			}

// // 		}
// // 	}()
// // 	for {
// // 		for i := 0; i < 10; i++ {
// // 			ch1 <- i
// // 			date := <-ch
// // 			fmt.Printf("%c", date)
// // 		}
// // 	}

// // }
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
	var wg sync.WaitGroup
	ch := make(chan gro, 20)
	Num := make([]gro,0)
	//ch1 := make(chan gro, 20)
	// ch2 := make(chan gro,20)
	rand.Seed(time.Now().UnixNano()) //在主程序中种下一颗种子
	wg.Add(20)
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
	for i  := range Num {
		fmt.Println(Num[i])
	}
	//更改 按照number的大小来排序
	// time.Sleep(time.Second * 3)
	// close(ch1)
	// go func() {
	// 	for i := range ch1 {
	// 		for j := range ch1 {
	// 			if i.number > j.number {
	// 				i, j = j, i
	// 			}
	// 		}
	// 	}
	// }()
	// time.Sleep(time.Second * 3)
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	for i := 0; i < 20; i++ {
// 		c := rand.Intn(100)
// 		time.Sleep(time.Second)
// 		fmt.Println(c)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// )

// func generateRandomNumber(wg *sync.WaitGroup, ch chan struct {
// 	num      int
// 	routineNo int
// }) {
// 	defer wg.Done()
// 	randomNumber := rand.Intn(100)
// 	routineNumber := rand.Intn(20)
// 	ch <- struct {
// 		num      int
// 		routineNo int
// 	}{randomNumber, routineNumber}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan struct {
// 		num      int
// 		routineNo int
// 	}, 20)

// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)
// 		go generateRandomNumber(&wg, ch)
// 	}

// 	wg.Wait()
// 	close(ch)

// 	// 排序前输出
// 	fmt.Println("排序前：")
// 	for item := range ch {
// 		fmt.Printf("随机数: %d, goroutine 编号: %d\n", item.num, item.routineNo)
// 	}

// 	// 排序逻辑（选择排序）
// 	for i := range ch {
// 		for j := range ch {
// 			if i.routineNo < j.routineNo {
// 				i, j = j, i
// 			}
// 		}
// 	}

// 	// 排序后输出
// 	fmt.Println("排序后：")
// 	for item := range ch {
// 		fmt.Printf("随机数: %d, goroutine 编号: %d\n", item.num, item.routineNo)
// 	}
// }
