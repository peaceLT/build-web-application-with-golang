package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// 生成一个唯一的token，用于防止CSRF攻击

		timestamp := strconv.Itoa(time.Now().Nanosecond()) // 获取当前时间的纳秒数
		hashWr := md5.New()                                // 创建一个新的MD5哈希对象
		hashWr.Write([]byte(timestamp))                    // 将时间戳写入哈希对象
		token := fmt.Sprintf("%x", hashWr.Sum(nil))        // 生成MD5哈希值作为token

		t, _ := template.ParseFiles("./htmlcode/login.html") // 解析login.html模板文件
		t.Execute(w, token)
	} else {
		r.ParseForm()

		token := r.Form.Get("token")
		if token != "" {
			// 验证token的合法性
			// 这里可以添加具体的token验证逻辑
		} else {
			// 如果token不存在，处理错误
			// 这里可以添加错误处理逻辑
		}

		// 输出用户名和密码信息到服务器端控制台
		fmt.Println("username length: ", len(r.Form["username"][0]))                // 输出用户名的长度
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) // 输出经过HTML转义的用户名
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password"))) // 输出经过HTML转义的密码

		// 将用户名经过HTML转义后输出到客户端
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/login", login)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
