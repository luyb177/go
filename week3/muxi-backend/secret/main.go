package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"muxi-backend/tool/getDecryptedPaper"
    "muxi-backend/tool/savePaper"
)

func main() {
	// 目标根URL
	url := "http://121.43.151.190:8000/"
	// 发送 GET 请求,返回的结果还需要进行处理才能得到你需要的结果
	response, err := http.Get(url + "paper")
	//func Get(url string) (resp *Response, err error)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()
	body1, err := ioutil.ReadAll(response.Body) //读取http响应的主体
	if err != nil {
		fmt.Println("ERROR OCCUED:",err.Error())
		return
	}
	p := string(body1) //获取加密论文
	//获取秘钥
	resp,err := http.Get(url + "secret")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
	body2, err := ioutil.ReadAll(resp.Body) //读取http响应的主体
	if err != nil {
		fmt.Println("ERROR OCCUED:",err.Error())
		return
	}
	k := string(body2) 
	a := getDecryptedPaper.GetDecryptedPaper(p,k)
	savePaper.SavePaper("D:\\Go Code\\src\\go - github\\go\\muxi-backend\\paper\\Academician Sun's papers.txt",a)
	//使用绝对路径时\要转义	
}
