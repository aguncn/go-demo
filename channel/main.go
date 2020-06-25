package main

import "fmt"
import "time"

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func getUrl(url string) {
	fmt.Println("start to getUrl", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}

func main() {
	for i := 0; i < 3; i++ {
		go getUrl("https://www.baidu.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
