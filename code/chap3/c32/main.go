package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                                 // 解析请求中的参数，默认情况下不会自动解析
	fmt.Println("url: ", r.URL)                   // 打印请求的URL
	fmt.Println("form: ", r.Form)                 // 打印解析后的参数到服务器端控制台
	fmt.Println("path: ", r.URL.Path)             // 打印请求的路径
	fmt.Println("scheme: ", r.URL.Scheme)         // 打印请求的URL方案（如http或https）
	fmt.Println("url_long: ", r.Form["url_long"]) // 打印名为"url_long"的参数值

	// 遍历所有解析的参数并打印键值对
	for k, v := range r.Form {
		fmt.Println("key:", k)                   // 打印参数的键
		fmt.Println("val:", strings.Join(v, "")) // 打印参数的值，多个值用空字符串连接
	}
	fmt.Fprintf(w, "Hello web") // 向客户端响应"Hello web"
}

func main() {
	http.HandleFunc("/", sayhelloName)                        //设置访问的路由
	if err := http.ListenAndServe(":9090", nil); err != nil { //设置监听的端口
		log.Fatal("ListenAndServe: ", err)
	}
}
